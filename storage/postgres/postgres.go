package postgres

import (
	"context"
	"fmt"
	"log"
	"organization_service/config"
	"organization_service/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	db        *pgxpool.Pool
	branch    storage.BranchRepoI
	store     storage.StoreRepoI
	employees storage.EmployeesRepoI
	supplier  storage.SupplierRepoI
	category  storage.CategoryRepoI
	product storage.ProductRepoI
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pool,
	}, err
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (l *Store) Log(ctx context.Context, msg string, data map[string]interface{}) {
	args := make([]interface{}, 0, len(data)+2) // making space for arguments + msg
	args = append(args, msg)
	for k, v := range data {
		args = append(args, fmt.Sprintf("%s=%v", k, v))
	}
	log.Println(args...)
}

func (s *Store) Branch() storage.BranchRepoI {
	if s.branch == nil {
		s.branch = NewBranchRepo(s.db)
	}

	return s.branch
}

func (s *Store) Store() storage.StoreRepoI {
	if s.store == nil {
		s.store = NewStoreRepo(s.db)
	}

	return s.store
}

func (s *Store) Employees() storage.EmployeesRepoI {
	if s.employees == nil {
		s.employees = NewEmployeesRepo(s.db)
	}

	return s.employees
}

func (s *Store) Supplier() storage.SupplierRepoI {
	if s.supplier == nil {
		s.supplier = NewSupplierRepo(s.db)
	}

	return s.supplier
}

func (s *Store) Category() storage.CategoryRepoI {
	if s.category == nil {
		s.category = NewCategoryRepo(s.db)
	}

	return s.category
}

func (s *Store) Product() storage.ProductRepoI {
	if s.product == nil {
		s.product = NewProductRepo(s.db)
	}

	return s.product
}