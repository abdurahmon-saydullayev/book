package postgres

import (
	"GoProject/book/genproto/author_service"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type authorRepo struct {
	db *pgxpool.Pool
}

func NewAuthorRepo(db *pgxpool.Pool) *authorRepo {
	return &authorRepo{
		db: db,
	}
}

func (r *authorRepo) Create(ctx context.Context, req *author_service.CreateAuthorRequest) (resp *author_service.AuthorPrimaryKey, err error) {
	query := `INSERT INTO author 
	(
	id,
	name, 
	age, 
	secondname
	 	) values
	(
		$1,
		$2,
		$3,
	)
	`

	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec(ctx, query,
	uuid.String(),
	req.Name,
	req.Age,
	req.Secondname,
	)
	if err != nil {
		return nil, err
	}
	
	resp = &author_service.AuthorPrimaryKey{
		Id :uuid.String(),
	}
	return
}

func (r *authorRepo) GetById(ctx context.Context, req *author_service.AuthorPrimaryKey)(resp *author_service.Author, err error){
	return 
}