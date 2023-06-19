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

type supplierService struct { 
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	// services client.ServiceManagerI
	organization_service.UnimplementedSupplierServiceServer
}

// svcs client.ServiceManagerI
func NewSupplierService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) organization_service.SupplierServiceServer {
	return &supplierService{
		cfg:  cfg,
		log:  log,
		strg: strg,
		// services: svcs,
	}
}

func (b *supplierService) Create(ctx context.Context, req *organization_service.CreateSupplierRequest) (resp *organization_service.Supplier, err error) {
	b.log.Info("---CreateSupplier--->", logger.Any("req", req))

	pKey, err := b.strg.Supplier().Create(ctx, req)

	if err != nil {
		b.log.Error("!!!CreateSupplier--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return b.strg.Supplier().Get(ctx, pKey)
}

func (b *supplierService) Get(ctx context.Context, req *organization_service.PrimaryKey) (resp *organization_service.Supplier, err error) {
	b.log.Info("---GetSupplier--->", logger.Any("req", req))

	resp, err = b.strg.Supplier().Get(ctx, req)
	if err != nil {
		b.log.Error("!!!GetSupplier--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, err
}

func (b *supplierService) GetList(ctx context.Context, req *organization_service.GetListSupplierRequest) (resp *organization_service.GetSuppliersListResponse, err error) {
	b.log.Info("---GetSuppliersList--->", logger.Any("req", req))

	resp, err = b.strg.Supplier().GetList(ctx, req)

	if err != nil {
		b.log.Error("!!!GetSuppliersList--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, err
}

func (b *supplierService) Update(ctx context.Context, req *organization_service.UpdateSupplierRequest) (resp *organization_service.Supplier, err error) {
	b.log.Info("---UpdateSupplier--->", logger.Any("req", req))

	rowsAffected, err := b.strg.Supplier().Update(ctx, req)

	if err != nil {
		b.log.Error("!!!UpdateSupplier--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = b.strg.Supplier().Get(ctx, &organization_service.PrimaryKey{Id: req.Supplier.Id})
	if err != nil {
		b.log.Error("!!!UpdateSupplier--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (b *supplierService) PatchUpdate(ctx context.Context, req *organization_service.PatchUpdateRequest) (resp *organization_service.Supplier, err error) {
	b.log.Info("---PatchUpdate--->", logger.Any("req", req))

	rowsAffected, err := b.strg.Supplier().PatchUpdate(ctx, req)

	if err != nil {
		b.log.Error("!!!PatchUpdate--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = b.strg.Supplier().Get(ctx, &organization_service.PrimaryKey{Id: req.Id})
	if err != nil {
		b.log.Error("!!!Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (b *supplierService) Delete(ctx context.Context, req *organization_service.PrimaryKey) (resp *empty.Empty, err error) {
	b.log.Info("---DeleteSupplier--->", logger.Any("req", req))

	resp = &empty.Empty{}

	rowsAffected, err := b.strg.Supplier().Delete(ctx, req)

	if err != nil {
		b.log.Error("!!!DeleteSupplier--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return resp, err
}
