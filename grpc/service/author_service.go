package service

import (
	"GoProject/book/config"
	"GoProject/book/genproto/author_service"
	"GoProject/book/grpc/client"
	"GoProject/book/pkg/logger"
	"GoProject/book/storage"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authorService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	author_service.UnimplementedAuthorServiceServer
}

func NewAuthorService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *authorService {
	return &authorService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (b *bookService) CreateAuthor(ctx context.Context, req *author_service.CreateAuthorRequest) (resp *author_service.Author, err error) {
	b.log.Info("---CreateAuthor--->", logger.Any("req", req))

	pKey, err := b.strg.Author().Create(ctx, req)
	if err != nil {
		b.log.Error("!!!CreateAuthor--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &author_service.Author{
		Id:         pKey.Id,
		Name:       "me",
		Secondname: "just",
		Age:        19,
	}, err
}
