package grpc

import (
	"GoProject/book/config"
	"GoProject/book/genproto/book_service"
	"GoProject/book/grpc/client"
	"GoProject/book/grpc/service"
	"GoProject/book/pkg/logger"
	"GoProject/book/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	book_service.RegisterBookServiceServer(grpcServer, service.NewBookService(cfg, log, strg, svcs))

	reflection.Register(grpcServer)
	return
}
