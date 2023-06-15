// Copyright 2019 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logging

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type registerConfig struct {
	shortName      string
	isTraceEnabled *bool
	onUpdate       func(newLogger *zap.Logger)
}

// RegisterOption are option parameters that you can set when registering a new logger
// in the system using `Register` function.
type RegisterOption interface {
	apply(config *registerConfig)
}

type registerOptionFunc func(config *registerConfig)

func (f registerOptionFunc) apply(config *registerConfig) {
	f(config)
}

// RegisterOnUpdate enable you to have a hook function that will receive the new logger
// that is going to be assigned to your logger instance. This is useful in some situation
// where you need to update other instances or re-configuring a bit the logger when
// a new one is attached.
//
// This is called **after** the instance has been re-assigned.
func RegisterOnUpdate(onUpdate func(newLogger *zap.Logger)) RegisterOption {
	return registerOptionFunc(func(config *registerConfig) {
		config.onUpdate = onUpdate
	})
}

func registerShortName(shortName string) RegisterOption {
	return registerOptionFunc(func(config *registerConfig) {
		config.shortName = shortName
	})
}

func registerWithTracer(isEnabled *bool) RegisterOption {
	return registerOptionFunc(func(config *registerConfig) {
		config.isTraceEnabled = isEnabled
	})
}

type LoggerExtender func(*zap.Logger) *zap.Logger

type registryEntry struct {
	packageID    string
	shortName    string
	atomicLevel  zap.AtomicLevel
	traceEnabled *bool
	logPtr       *zap.Logger
	onUpdate     func(newLogger *zap.Logger)
}

func (e *registryEntry) String() string {
	loggerPtr := "<nil>"
	if e.logPtr != nil {
		loggerPtr = fmt.Sprintf("%p", e.logPtr)
	}

	traceEnabled := false
	if e.traceEnabled != nil {
		traceEnabled = *e.traceEnabled
	}

	return fmt.Sprintf("%s @ %s (level: %s, trace?: %t, ptr: %s)", e.shortName, e.packageID, e.atomicLevel.Level(), traceEnabled, loggerPtr)
}

var globalRegistry = newRegistry("global")
var defaultLogger = zap.NewNop()

// Deprecated: will be replaced by RegisterLogger
func Register(packageID string, zlogPtr **zap.Logger, options ...RegisterOption) {
	if *zlogPtr == nil {
		*zlogPtr = zap.NewNop()
	}
	register(globalRegistry, packageID, *zlogPtr, options...)
}

func RegisterLogger(packageID string, zlogPtr *zap.Logger, options ...RegisterOption) {
	register(globalRegistry, packageID, zlogPtr, options...)
}

func register2(registry *registry, shortName string, packageID string, options ...RegisterOption) (*zap.Logger, Tracer) {
	logger := zap.NewNop()
	tracer := boolTracer{new(bool)}

	allOptions := append([]RegisterOption{
		registerShortName(shortName),
		registerWithTracer(tracer.value),
	}, options...)

	register(registry, packageID, logger, allOptions...)

	return logger, tracer
}

func registerDebug(registry *registry, shortName string, packageID string, logger *zap.Logger, options ...RegisterOption) (*zap.Logger, Tracer) {
	tracer := boolTracer{new(bool)}

	allOptions := append([]RegisterOption{
		registerShortName(shortName),
		registerWithTracer(tracer.value),
	}, options...)

	register(registry, packageID, logger, allOptions...)

	return logger, tracer
}

func register(registry *registry, packageID string, zlogPtr *zap.Logger, options ...RegisterOption) {
	if zlogPtr == nil {
		panic("the zlog pointer (of type **zap.Logger) must be set")
	}

	config := registerConfig{}
	for _, opt := range options {
		opt.apply(&config)
	}

	entry := &registryEntry{
		packageID:    packageID,
		shortName:    config.shortName,
		traceEnabled: config.isTraceEnabled,
		atomicLevel:  zap.NewAtomicLevelAt(zapcore.ErrorLevel),
		logPtr:       zlogPtr,
		onUpdate:     config.onUpdate,
	}

	registry.registerEntry(entry)

	logger := defaultLogger
	if zlogPtr != nil {
		logger = zlogPtr
	}

	// The tracing has already been set, so we can go unspecified here to not change anything
	setLogger(entry, logger, unspecifiedTracing)
}

func Set(logger *zap.Logger, regexps ...string) {
	for name, entry := range globalRegistry.entriesByPackageID {
		if len(regexps) == 0 {
			setLogger(entry, logger, unspecifiedTracing)
		} else {
			for _, re := range regexps {
				regex, err := regexp.Compile(re)
				if (err == nil && regex.MatchString(name)) || (err != nil && name == re) {
					setLogger(entry, logger, unspecifiedTracing)
				}
			}
		}
	}
}

// Extend is different than `Set` by being able to re-configure the existing logger set for
// all registered logger in the registry. This is useful for example to add a field to the
// currently set logger:
//
// ```
// logger.Extend(func (current *zap.Logger) { return current.With("name", "value") }, "github.com/dfuse-io/app.*")
// ```
func Extend(extender LoggerExtender, regexps ...string) {
	extend(extender, unspecifiedTracing, regexps...)
}

func extend(extender LoggerExtender, tracing tracingType, regexps ...string) {
	for name, entry := range globalRegistry.entriesByPackageID {
		if entry.logPtr == nil {
			continue
		}

		if len(regexps) == 0 {
			setLogger(entry, extender(entry.logPtr), tracing)
		} else {
			for _, re := range regexps {
				if regexp.MustCompile(re).MatchString(name) {
					setLogger(entry, extender(entry.logPtr), tracing)
				}
			}
		}
	}
}

// Override sets the given logger on previously registered and next
// registrations.  Useful in tests.
func Override(logger *zap.Logger) {
	defaultLogger = logger
	Set(logger)
}

// TestingOverride calls `Override` (or `Set`, see below) with a development
// logger setup correctly with the right level based on some environment variables.
//
// By default, override using a `zap.NewDevelopment` logger (`info`), if
// environment variable `DEBUG` is set to anything or environment variable `TRACE`
// is set to `true`, logger is set in `debug` level.
//
// If `DEBUG` is set to something else than `true` and/or if `TRACE` is set
// to something else than
func TestingOverride() {
	debug := os.Getenv("DEBUG")
	trace := os.Getenv("TRACE")
	if debug == "" && trace == "" {
		return
	}

	logger, _ := zap.NewDevelopment()

	regex := ""
	if debug != "true" {
		regex = debug
	}

	if regex == "" && trace != "true" {
		regex = trace
	}

	if regex == "" {
		Override(logger)
	} else {
		for _, regexPart := range strings.Split(regex, ",") {
			regexPart = strings.TrimSpace(regexPart)
			if regexPart != "" {
				Set(logger, regexPart)
			}
		}
	}
}

type tracingType uint8

const (
	unspecifiedTracing tracingType = iota
	enableTracing
	disableTracing
)

func setLogger(entry *registryEntry, logger *zap.Logger, tracing tracingType) {
	if entry == nil || logger == nil {
		return
	}

	ve := reflect.ValueOf(entry.logPtr).Elem()
	ve.Set(reflect.ValueOf(logger).Elem())

	if entry.traceEnabled != nil && tracing != unspecifiedTracing {
		switch tracing {
		case enableTracing:
			*entry.traceEnabled = true
		case disableTracing:
			*entry.traceEnabled = false
		}
	}

	if entry.onUpdate != nil {
		entry.onUpdate(logger)
	}
}

type registry struct {
	sync.RWMutex

	name               string
	factory            loggerFactory
	entriesByPackageID map[string]*registryEntry
	entriesByShortName map[string][]*registryEntry

	dbgLogger *zap.Logger
}

func newRegistry(name string) *registry {
	return &registry{
		name:               name,
		entriesByPackageID: make(map[string]*registryEntry),
		entriesByShortName: make(map[string][]*registryEntry),
		factory: func(name string, level zap.AtomicLevel) *zap.Logger {
			loggerOptions := newLoggerOptions("", WithAtomicLevel(level))
			if name != "" {
				loggerOptions.loggerName = name
			}

			return newLogger(&loggerOptions)
		},

		dbgLogger: dbgZlog.With(zap.String("registry", name)),
	}
}

func (r *registry) registerEntry(entry *registryEntry) {
	if entry == nil {
		panic("refusing to add a nil registry entry")
	}

	id := validateEntryIdentifier("package ID", entry.packageID, false)
	shortName := validateEntryIdentifier("short name", entry.shortName, true)

	if actual := r.entriesByPackageID[id]; actual != nil {
		panic(fmt.Sprintf("packageID %q is already registered", id))
	}

	entry.packageID = id
	entry.shortName = shortName

	r.entriesByPackageID[id] = entry
	if shortName != "" {
		r.entriesByShortName[shortName] = append(r.entriesByShortName[shortName], entry)
	}

	r.dbgLogger.Info("registered entry", zap.String("short_name", shortName), zap.String("id", id))
}

func (r *registry) forAllEntries(callback func(entry *registryEntry)) {
	for _, entry := range r.entriesByPackageID {
		callback(entry)
	}
}

func (r *registry) forAllEntriesMatchingSpec(spec *logLevelSpec, callback func(entry *registryEntry, level zapcore.Level, trace bool)) {
	for _, specForKey := range spec.sortedSpecs() {
		if specForKey.key == "true" || specForKey.key == "*" {
			for _, entry := range r.entriesByPackageID {
				callback(entry, specForKey.level, specForKey.trace)
			}

			continue
		}

		r.forEntriesMatchingSpec(specForKey, callback)
	}
}

func (r *registry) forEntriesMatchingSpec(spec *levelSpec, callback func(entry *registryEntry, level zapcore.Level, trace bool)) {
	r.dbgLogger.Debug("looking in short names to find spec key", zap.String("key", spec.key))
	entries, found := r.entriesByShortName[spec.key]
	if found {
		r.dbgLogger.Debug("found logger in short names", zap.Int("count", len(entries)))
		for _, entry := range entries {
			callback(entry, spec.level, spec.trace)
		}
		return
	}

	r.dbgLogger.Debug("looking in package IDs to find spec key", zap.String("key", spec.key))
	entry, found := r.entriesByPackageID[spec.key]
	if found {
		r.dbgLogger.Debug("found logger in package ID", zap.Stringer("entry", entry))
		callback(entry, spec.level, spec.trace)
		return
	}

	r.dbgLogger.Debug("looking in package IDs by regex", zap.String("key", spec.key))
	regex, err := regexp.Compile(spec.key)
	if err != nil {
		r.dbgLogger.Debug("spec key is not a regex, we already matched exact package ID, nothing to do more", zap.Error(err))
		return
	}

	for packageID, entry := range r.entriesByPackageID {
		if regex.MatchString(packageID) {
			callback(entry, spec.level, spec.trace)
		}
	}
}

func (r *registry) setLoggerForEntry(entry *registryEntry, level zapcore.Level, trace bool) {
	if entry == nil {
		return
	}

	entry.atomicLevel.SetLevel(level)
	logger := r.factory(entry.shortName, entry.atomicLevel)

	ve := reflect.ValueOf(entry.logPtr).Elem()
	ve.Set(reflect.ValueOf(logger).Elem())

	// It's possible for an entry to have no tracer registered, for example if the legacy
	// register method is used. We must protect from this and not set anything.
	if entry.traceEnabled != nil {
		*entry.traceEnabled = trace
	}

	if entry.onUpdate != nil {
		entry.onUpdate(logger)
	}

	r.dbgLogger.Info("set logger on entry", zap.Stringer("entry", entry), zap.Stringer("to_level", level), zap.Bool("trace_enabled", trace))
}

func (r *registry) setLevelForEntry(entry *registryEntry, level zapcore.Level, trace bool) {
	if entry == nil {
		return
	}

	entry.atomicLevel.SetLevel(level)

	// It's possible for an entry to have no tracer registered, for example if the legacy
	// register method is used. We must protect from this and not set anything.
	if entry.traceEnabled != nil {
		*entry.traceEnabled = trace
	}

	r.dbgLogger.Info("set level on entry", zap.Stringer("entry", entry))
}

func validateEntryIdentifier(tag string, rawInput string, allowEmpty bool) string {
	input := strings.TrimSpace(rawInput)
	if input == "" && !allowEmpty {
		panic(fmt.Errorf("the %s %q is invalid, must not be empty", tag, input))
	}

	if input == "true" {
		panic(fmt.Errorf("the %s %q is invalid, the identifier 'true' is reserved", tag, input))
	}

	if input == "*" {
		panic(fmt.Errorf("the %s %q is invalid, the identifier '*' is reserved", tag, input))
	}

	if strings.HasPrefix(input, "-") {
		panic(fmt.Errorf("the %s %q is invalid, must not starts with the '-' character", tag, input))
	}

	if strings.Contains(input, ",") {
		panic(fmt.Errorf("the %s %q is invalid, must not contain the ',' character", tag, input))
	}

	if strings.Contains(input, "=") {
		panic(fmt.Errorf("the %s %q is invalid, must not contain the '=' character", tag, input))
	}

	return input
}
