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

type storeService struct { // ? branchService ochilgan service ning book ga tegishli methodlarini birlashtirish uchun
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	// services client.ServiceManagerI
	organization_service.UnimplementedStoreServiceServer
}

// svcs client.ServiceManagerI
func NewStoreService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) organization_service.StoreServiceServer {
	return &storeService{
		cfg:  cfg,
		log:  log,
		strg: strg,
		// services: svcs,
	}
}

func (b *storeService) Create(ctx context.Context, req *organization_service.CreateStoreRequest) (resp *organization_service.Store, err error) {
	b.log.Info("---CreateStore--->", logger.Any("req", req))

	pKey, err := b.strg.Store().Create(ctx, req)

	if err != nil {
		b.log.Error("!!!CreateStore--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return b.strg.Store().Get(ctx, pKey)
}

func (b *storeService) Get(ctx context.Context, req *organization_service.PrimaryKey) (resp *organization_service.Store, err error) {
	b.log.Info("---GetStore--->", logger.Any("req", req))

	resp, err = b.strg.Store().Get(ctx, req)
	if err != nil {
		b.log.Error("!!!GetStore--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, err
}

func (b *storeService) GetList(ctx context.Context, req *organization_service.GetListStoreRequest) (resp *organization_service.GetStoresListResponse, err error) {
	b.log.Info("---GetStoresList--->", logger.Any("req", req))

	resp, err = b.strg.Store().GetList(ctx, req)

	if err != nil {
		b.log.Error("!!!GetStoresList--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, err
}

func (b *storeService) Update(ctx context.Context, req *organization_service.UpdateStoreRequest) (resp *organization_service.Store, err error) {
	b.log.Info("---UpdateStore--->", logger.Any("req", req))

	rowsAffected, err := b.strg.Store().Update(ctx, req)

	if err != nil {
		b.log.Error("!!!UpdateStore--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = b.strg.Store().Get(ctx, &organization_service.PrimaryKey{Id: req.Store.Id})
	if err != nil {
		b.log.Error("!!!UpdateStore--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (b *storeService) PatchUpdate(ctx context.Context, req *organization_service.PatchUpdateRequest) (resp *organization_service.Store, err error) {
	b.log.Info("---PatchUpdate--->", logger.Any("req", req))

	rowsAffected, err := b.strg.Store().PatchUpdate(ctx, req)

	if err != nil {
		b.log.Error("!!!PatchUpdate--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = b.strg.Store().Get(ctx, &organization_service.PrimaryKey{Id: req.Id})
	if err != nil {
		b.log.Error("!!!Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (b *storeService) Delete(ctx context.Context, req *organization_service.PrimaryKey) (resp *empty.Empty, err error) {
	b.log.Info("---DeleteStore--->", logger.Any("req", req))

	resp = &empty.Empty{}

	rowsAffected, err := b.strg.Store().Delete(ctx, req)

	if err != nil {
		b.log.Error("!!!DeleteStore--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return resp, err
}
