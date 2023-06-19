package service

import (
	"context"
	"organization_service/config"
	"organization_service/genproto/organization_service"
	"organization_service/pkg/logger"
	"organization_service/storage"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type employeesService struct { // ? branchService ochilgan service ning book ga tegishli methodlarini birlashtirish uchun
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	// services client.ServiceManagerI
	organization_service.UnimplementedEmployeesServiceServer
}

// svcs client.ServiceManagerI
func NewEmployeesService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) organization_service.EmployeesServiceServer {
	return &employeesService{
		cfg:  cfg,
		log:  log,
		strg: strg,
		// services: svcs,
	}
}

func (b *employeesService) Create(ctx context.Context, req *organization_service.CreateEmployeesRequest) (resp *organization_service.Employees, err error) {
	b.log.Info("---CreateEmployees--->", logger.Any("req", req))

	pKey, err := b.strg.Employees().Create(ctx, req)

	if err != nil {
		b.log.Error("!!!CreateEmployees--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return b.strg.Employees().Get(ctx, pKey)
}

func (b *employeesService) Get(ctx context.Context, req *organization_service.PrimaryKey) (resp *organization_service.Employees, err error) {
	b.log.Info("---GetEmployees--->", logger.Any("req", req))

	resp, err = b.strg.Employees().Get(ctx, req)
	if err != nil {
		b.log.Error("!!!GetEmployees--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, err
}

func (b *employeesService) GetList(ctx context.Context, req *organization_service.GetEmployeesListRequest) (resp *organization_service.GetEmployeessListResponse, err error) {
	b.log.Info("---GetEmployeessList--->", logger.Any("req", req))

	resp, err = b.strg.Employees().GetList(ctx, req)

	if err != nil {
		b.log.Error("!!!GetEmployeessList--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, err
}

func (b *employeesService) Update(ctx context.Context, req *organization_service.UpdateEmployeesRequest) (resp *organization_service.Employees, err error) {
	b.log.Info("---UpdateEmployees--->", logger.Any("req", req))

	rowsAffected, err := b.strg.Employees().Update(ctx, req)

	if err != nil {
		b.log.Error("!!!UpdateBranch--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = b.strg.Employees().Get(ctx, &organization_service.PrimaryKey{Id: req.Employees.Id})
	if err != nil {
		b.log.Error("!!!UpdateEmployees--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (b *employeesService) PatchUpdate(ctx context.Context, req *organization_service.PatchUpdateRequest) (resp *organization_service.Employees, err error) {
	b.log.Info("---PatchUpdate--->", logger.Any("req", req))

	rowsAffected, err := b.strg.Employees().PatchUpdate(ctx, req)
	if err != nil {
		b.log.Error("!!!PatchUpdate--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = b.strg.Employees().Get(ctx, &organization_service.PrimaryKey{Id: req.Id})
	if err != nil {
		b.log.Error("!!!Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (b *employeesService) Delete(ctx context.Context, req *organization_service.PrimaryKey) (resp *empty.Empty, err error) {
	b.log.Info("---DeleteEmployees--->", logger.Any("req", req))

	resp = &empty.Empty{}

	rowsAffected, err := b.strg.Employees().Delete(ctx, req)

	if err != nil {
		b.log.Error("!!!DeleteEmployees--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return resp, err
}
