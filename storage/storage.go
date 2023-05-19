package storage

import (
	"GoProject/book/genproto/author_service"
	"GoProject/book/genproto/book_service"
	"context"
)

type StorageI interface {
	CloseDB()
	Book() BookRepoI
	Author() AuthorRepoI
}

type BookRepoI interface {
	Create(ctx context.Context, req *book_service.CreateBookRequest) (resp *book_service.BookPrimaryKey, err error)
	Get(ctx context.Context, req *book_service.BookPrimaryKey) (resp *book_service.Book, err error)
	GetList(ctx context.Context, req *book_service.GetBooksListRequest) (resp *book_service.GetBooksListResponse, err error)
	Update(ctx context.Context, req *book_service.UpdateBookRequest) (rowsAffected int64, err error)
	PatchUpdate(ctx context.Context, req *book_service.PatchUpdateRequest) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *book_service.BookPrimaryKey) (rowsAffected int64, err error)
}

type AuthorRepoI interface{
	Create (ctx context.Context, req *author_service.CreateAuthorRequest)(resp *author_service.AuthorPrimaryKey, err error)
	GetById(ctx context.Context, req *author_service.GetAuthorListRequest)(resp *author_service.AuthorPrimaryKey, err error)
}
