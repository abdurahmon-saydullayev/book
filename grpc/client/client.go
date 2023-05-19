package client

import (
	"GoProject/book/config"
	"GoProject/book/genproto/author_service"
	"GoProject/book/genproto/book_service"

	"google.golang.org/grpc"
)

type ServiceManagerI interface {
	BookService() book_service.BookServiceClient
	AuthorService() author_service.AuthorServiceClient
}

type grpcClients struct {
	bookService book_service.BookServiceClient
	authorService author_service.AuthorServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connBookService, err := grpc.Dial(
		cfg.ServiceHost+cfg.ServicePort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	connAuthorService, err := grpc.Dial(
		cfg.ServiceHost+cfg.ServicePort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		bookService: book_service.NewBookServiceClient(connBookService),
		authorService: author_service.NewAuthorServiceClient(connAuthorService),
	}, nil
}

func (g *grpcClients) BookService() book_service.BookServiceClient {
	return g.bookService
}

func (g *grpcClients) AuthorService()author_service.AuthorServiceClient{
	return g.authorService
}
