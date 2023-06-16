package main

import (
	"context"
	"net"
	"organization_service/config"
	"organization_service/grpc"
	"organization_service/pkg/logger"
	"organization_service/storage/postgres"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	loggerLevel := logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer logger.Cleanup(log)

	pgStore, err := postgres.NewPostgres(context.Background(), cfg)
	if err != nil {
		log.Panic("postgres.NewPostgres", logger.Error(err))
	}
	defer pgStore.CloseDB()

	// fmt.Println(pgStore.Branch().Create(context.Background(),&organization_service.CreateBranchRequest{BranchCode: "maxmurjon",Name: "sndivnsfd",Address: "asfnoisdngiovrsn",PhoneNumber: "+998971153311"}))

	// fmt.Println(pgStore.Branch().Get(context.Background(),&organization_service.PrimaryKey{Id: "7df95af6-17aa-4dd6-9f90-a01d29ebd5ee"}))

	// fmt.Println(pgStore.Branch().Update(context.Background(),&organization_service.UpdateBranchRequest{Branch: &organization_service.Branch{Id: "7df95af6-17aa-4dd6-9f90-a01d29ebd5ee",BranchCode: "MP-0000001",Name: "Mega Planet",Address: "Yunusobot",PhoneNumber: "+998971153311"}}))

	// fmt.Println(pgStore.Branch().Delete(context.Background(),&organization_service.PrimaryKey{Id: "7df95af6-17aa-4dd6-9f90-a01d29ebd5ee"}))

	// fmt.Println(pgStore.Branch().PatchUpdate(context.Background(),&organization_service.PatchUpdateRequest{Id: "99bbf686-6b4b-4398-bfdb-1e15fe239742",Params: []*organization_service.Obj{{Key: "branch_code",Value: "MP-0000001"}}}))

	// svcs, err := client.NewGrpcClients(cfg)
	// if err != nil {
	// 	log.Panic("client.NewGrpcClients", logger.Error(err))
	// }

	grpcServer := grpc.SetUpServer(cfg, log, pgStore)

	lis, err := net.Listen("tcp", cfg.ServicePort)
	if err != nil {
		log.Panic("net.Listen", logger.Error(err))
	}

	log.Info("GRPC: Server being started...", logger.String("port", cfg.ServicePort))

	if err := grpcServer.Serve(lis); err != nil {
		log.Panic("grpcServer.Serve", logger.Error(err))
	}
}
