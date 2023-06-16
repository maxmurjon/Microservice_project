package grpc

import (
	"organization_service/config"
	"organization_service/genproto/organization_service"
	"organization_service/grpc/service"
	"organization_service/pkg/logger"
	"organization_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)
// , svcs client.ServiceManagerI
func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	organization_service.RegisterBranchServiceServer(grpcServer, service.NewBranchService(cfg, log, strg))

	organization_service.RegisterStoreServiceServer(grpcServer, service.NewStoreService(cfg, log, strg))

	organization_service.RegisterEmployeesServiceServer(grpcServer, service.NewEmployeesService(cfg, log, strg))

	organization_service.RegisterCategoryServiceServer(grpcServer, service.NewCategoryService(cfg, log, strg))

	organization_service.RegisterProductServiceServer(grpcServer, service.NewProductService(cfg, log, strg))

	organization_service.RegisterSupplierServiceServer(grpcServer, service.NewSupplierService(cfg, log, strg))

	reflection.Register(grpcServer)
	return
}
