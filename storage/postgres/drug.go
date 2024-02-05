package postgres

import (
	"context"
	"shifolink/api/models"
	"shifolink/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type drugRepo struct {
	pool *pgxpool.Pool
}

func NewDrugRepo(pool *pgxpool.Pool) storage.IDrugRepo {
	return &drugRepo{
		pool: pool,
	}
}

func (d *drugRepo) Create(ctx context.Context, request models.CreateDrug) (string, error) {

	return "", nil
}

func (d *drugRepo) Get(ctx context.Context, request models.PrimaryKey) (models.Drug, error) {

	return models.Drug{}, nil
}

func (d *drugRepo) GetList(ctx context.Context, request models.GetListRequest) (models.DrugsResponse, error) {

	return models.DrugsResponse{}, nil
}

func (d *drugRepo) Update(ctx context.Context, request models.UpdateDrug) (string, error) {

	return "", nil
}

func (d *drugRepo) Delete(ctx context.Context, id string) error {

	return nil
}
