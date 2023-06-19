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

type categoryService struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	// services client.ServiceManagerI
	organization_service.UnimplementedCategoryServiceServer
}

// svcs client.ServiceManagerI
func NewCategoryService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) organization_service.CategoryServiceServer {
	return &categoryService{
		cfg:  cfg,
		log:  log,
		strg: strg,
		// services: svcs,
	}
}

func (b *categoryService) Create(ctx context.Context, req *organization_service.CreateCategoryRequest) (resp *organization_service.Category, err error) {
	b.log.Info("---CreateCategory--->", logger.Any("req", req))

	pKey, err := b.strg.Category().Create(ctx, req)

	if err != nil {
		b.log.Error("!!!CreateCategory--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return b.strg.Category().Get(ctx, pKey)
}

func (b *categoryService) Get(ctx context.Context, req *organization_service.PrimaryKey) (resp *organization_service.Category, err error) {
	b.log.Info("---GetCategory--->", logger.Any("req", req))

	resp, err = b.strg.Category().Get(ctx, req)
	if err != nil {
		b.log.Error("!!!GetCategory--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, err
}

// func (b *categoryService) GetBooksList(ctx context.Context, req *book_service.GetBooksListRequest) (resp *book_service.GetBooksListResponse, err error) {
// 	b.log.Info("---GetBooksList--->", logger.Any("req", req))

// 	resp, err = b.strg.Book().GetList(ctx, req)

// 	if err != nil {
// 		b.log.Error("!!!GetBooksList--->", logger.Error(err))
// 		return nil, status.Error(codes.InvalidArgument, err.Error())
// 	}

// 	return resp, err
// }

func (b *categoryService) Update(ctx context.Context, req *organization_service.UpdateCategoryRequest) (resp *organization_service.Category, err error) {
	b.log.Info("---UpdateCategory--->", logger.Any("req", req))

	rowsAffected, err := b.strg.Category().Update(ctx, req)

	if err != nil {
		b.log.Error("!!!UpdateCategory--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = b.strg.Category().Get(ctx, &organization_service.PrimaryKey{Id: req.Category.Id})
	if err != nil {
		b.log.Error("!!!UpdateCategory--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (b *categoryService) PatchUpdate(ctx context.Context, req *organization_service.PatchUpdateRequest) (resp *organization_service.Category, err error) {
	b.log.Info("---PatchUpdate--->", logger.Any("req", req))

	rowsAffected, err := b.strg.Category().PatchUpdate(ctx, req)

	if err != nil {
		b.log.Error("!!!PatchUpdate--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = b.strg.Category().Get(ctx, &organization_service.PrimaryKey{Id: req.Id})
	if err != nil {
		b.log.Error("!!!Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (b *categoryService) Delete(ctx context.Context, req *organization_service.PrimaryKey) (resp *empty.Empty, err error) {
	b.log.Info("---DeleteCategory--->", logger.Any("req", req))

	resp = &empty.Empty{}

	rowsAffected, err := b.strg.Category().Delete(ctx, req)

	if err != nil {
		b.log.Error("!!!DeleteCategory--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return resp, err
}
