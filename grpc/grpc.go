package grpc

import (
	"backend-templates/golang-psql/config"
	"backend-templates/golang-psql/genproto/book_service"
	"backend-templates/golang-psql/grpc/client"
	"backend-templates/golang-psql/grpc/service"
	"backend-templates/golang-psql/pkg/logger"
	"backend-templates/golang-psql/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer() // ? grpc yangi server ochmoqda

	book_service.RegisterBookServiceServer(grpcServer, service.NewBookService(cfg, log, strg, svcs)) // ? bu proto ni qaysi serverdan nimani uzatishni yozib olamiz

	reflection.Register(grpcServer) 
	return
}
