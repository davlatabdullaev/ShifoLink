package postgres

import (
	"context"
	"shifolink/api/models"
	"shifolink/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type drugStoreRepo struct {
	pool *pgxpool.Pool
}

func NewDrugStoreRepo(pool *pgxpool.Pool) storage.IDrugStoreRepo {
	return &drugStoreRepo{
		pool: pool,
	}
}

func (d *drugStoreRepo) Create(ctx context.Context, request models.CreateDrugStore) (string, error) {

	return "", nil
}

func (d *drugStoreRepo) Get(ctx context.Context, request models.PrimaryKey) (models.DrugStore, error) {

	return models.DrugStore{}, nil
}

func (d *drugStoreRepo) GetList(ctx context.Context, request models.GetListRequest) (models.DrugStoresResponse, error) {

	return models.DrugStoresResponse{}, nil
}

func (d *drugStoreRepo) Update(ctx context.Context, request models.UpdateDrugStore) (string, error) {

	return "", nil
}

func (d *drugStoreRepo) Delete(ctx context.Context, id string) error {

	return nil
}
