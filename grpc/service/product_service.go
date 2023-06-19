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

type productService struct { 
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	// services client.ServiceManagerI
	organization_service.UnimplementedProductServiceServer
}

// svcs client.ServiceManagerI
func NewProductService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) organization_service.ProductServiceServer {
	return &productService{
		cfg:  cfg,
		log:  log,
		strg: strg,
		// services: svcs,
	}
}

func (b *productService) Create(ctx context.Context, req *organization_service.CreateProductRequest) (resp *organization_service.Product, err error) {
	b.log.Info("---CreateProduct--->", logger.Any("req", req))

	pKey, err := b.strg.Product().Create(ctx, req)

	if err != nil {
		b.log.Error("!!!CreateProduct--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return b.strg.Product().Get(ctx, pKey)
}

func (b *productService) Get(ctx context.Context, req *organization_service.PrimaryKey) (resp *organization_service.Product, err error) {
	b.log.Info("---GetProduct--->", logger.Any("req", req))

	resp, err = b.strg.Product().Get(ctx, req)
	if err != nil {
		b.log.Error("!!!GetProduct--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, err
}

func (b *productService) GetList(ctx context.Context, req *organization_service.GetListProductRequest) (resp *organization_service.GetProductsListResponse, err error) {
	b.log.Info("---GetProductList--->", logger.Any("req", req))

	resp, err = b.strg.Product().GetList(ctx, req)

	if err != nil {
		b.log.Error("!!!GetProductList--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, err
}

func (b *productService) Update(ctx context.Context, req *organization_service.UpdateProductRequest) (resp *organization_service.Product, err error) {
	b.log.Info("---UpdateProduct--->", logger.Any("req", req))

	rowsAffected, err := b.strg.Product().Update(ctx, req)

	if err != nil {
		b.log.Error("!!!UpdateProduct--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = b.strg.Product().Get(ctx, &organization_service.PrimaryKey{Id: req.Product.Id})
	if err != nil {
		b.log.Error("!!!UpdateProduct--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (b *productService) PatchUpdate(ctx context.Context, req *organization_service.PatchUpdateRequest) (resp *organization_service.Product, err error) {
	b.log.Info("---PatchUpdate--->", logger.Any("req", req))

	rowsAffected, err := b.strg.Product().PatchUpdate(ctx, req)

	if err != nil {
		b.log.Error("!!!PatchUpdate--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = b.strg.Product().Get(ctx, &organization_service.PrimaryKey{Id: req.Id})
	if err != nil {
		b.log.Error("!!!Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (b *productService) Delete(ctx context.Context, req *organization_service.PrimaryKey) (resp *empty.Empty, err error) {
	b.log.Info("---DeleteProduct--->", logger.Any("req", req))

	resp = &empty.Empty{}

	rowsAffected, err := b.strg.Product().Delete(ctx, req)

	if err != nil {
		b.log.Error("!!!DeleteProduct--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return resp, err
}
