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

type branchService struct { // ? branchService ochilgan service ning book ga tegishli methodlarini birlashtirish uchun
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	// services client.ServiceManagerI
	organization_service.UnimplementedBranchServiceServer
}

// svcs client.ServiceManagerI
func NewBranchService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) organization_service.BranchServiceServer {
	return &branchService{
		cfg:  cfg,
		log:  log,
		strg: strg,
		// services: svcs,
	}
}

func (b *branchService) Create(ctx context.Context, req *organization_service.CreateBranchRequest) (resp *organization_service.Branch, err error) {
	b.log.Info("---CreateBranch--->", logger.Any("req", req))

	pKey, err := b.strg.Branch().Create(ctx, req)

	if err != nil {
		b.log.Error("!!!CreateBranch--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return b.strg.Branch().Get(ctx, pKey)
}

func (b *branchService) Get(ctx context.Context, req *organization_service.PrimaryKey) (resp *organization_service.Branch, err error) {
	b.log.Info("---GetBranch--->", logger.Any("req", req))

	resp, err = b.strg.Branch().Get(ctx, req)
	if err != nil {
		b.log.Error("!!!GetBranch--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, err
}

func (b *branchService) GetList(ctx context.Context, req *organization_service.GetListBranchRequest) (resp *organization_service.GetBranchsListResponse, err error) {
	b.log.Info("---GetBranchsList--->", logger.Any("req", req))

	resp, err = b.strg.Branch().GetList(ctx, req)

	if err != nil {
		b.log.Error("!!!GetBranchsList--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, err
}

func (b *branchService) Update(ctx context.Context, req *organization_service.UpdateBranchRequest) (resp *organization_service.Branch, err error) {
	b.log.Info("---UpdateBranch--->", logger.Any("req", req))

	rowsAffected, err := b.strg.Branch().Update(ctx, req)

	if err != nil {
		b.log.Error("!!!UpdateBranch--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = b.strg.Branch().Get(ctx, &organization_service.PrimaryKey{Id: req.Branch.Id})
	if err != nil {
		b.log.Error("!!!UpdateBranch--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (b *branchService) PatchUpdate(ctx context.Context, req *organization_service.PatchUpdateRequest) (resp *organization_service.Branch, err error) {
	b.log.Info("---PatchUpdate--->", logger.Any("req", req))

	rowsAffected, err := b.strg.Branch().PatchUpdate(ctx, req)

	if err != nil {
		b.log.Error("!!!PatchUpdate--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = b.strg.Branch().Get(ctx, &organization_service.PrimaryKey{Id: req.Id})
	if err != nil {
		b.log.Error("!!!Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (b *branchService) Delete(ctx context.Context, req *organization_service.PrimaryKey) (resp *empty.Empty, err error) {
	b.log.Info("---DeleteBranch--->", logger.Any("req", req))

	resp = &empty.Empty{}

	rowsAffected, err := b.strg.Branch().Delete(ctx, req)

	if err != nil {
		b.log.Error("!!!DeleteBranch--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return resp, err
}
