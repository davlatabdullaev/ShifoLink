package postgres

import (
	"context"
	"shifolink/api/models"
	"shifolink/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type authorRepo struct {
	pool *pgxpool.Pool
}

func NewAuthorRepo(pool *pgxpool.Pool) storage.IAuthorRepo {
	return &authorRepo{
		pool: pool,
	}
}

func (a *authorRepo) Create(ctx context.Context, request models.CreateAuthor) (string, error) {

	return "", nil
}

func (a *authorRepo) Get(ctx context.Context, request models.PrimaryKey) (models.Author, error) {

	return models.Author{}, nil
}

func (a *authorRepo) GetList(ctx context.Context, request models.GetListRequest) (models.AuthorsResponse, error) {

	return models.AuthorsResponse{}, nil
}

func (a *authorRepo) Update(ctx context.Context, request models.UpdateAuthor) (string, error) {

	return "", nil
}

func (a *authorRepo) Delete(ctx context.Context, id string) error {

	return nil
}

func (a *authorRepo) UpdatePassword(ctx context.Context, id string) error {

	return nil
}
