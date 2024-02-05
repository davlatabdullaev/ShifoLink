package postgres

import (
	"context"
	"shifolink/api/models"
	"shifolink/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type drugStoreBranchRepo struct {
	pool *pgxpool.Pool
}

func NewDrugStoreBranchRepo(pool *pgxpool.Pool) storage.IDrugStoreBranchRepo {
	return &drugStoreBranchRepo{
		pool: pool,
	}
}

func (d *drugStoreBranchRepo) Create(ctx context.Context, request models.CreateDrugStoreBranch) (string, error) {

	return "", nil
}

func (d *drugStoreBranchRepo) Get(ctx context.Context, request models.PrimaryKey) (models.DrugStoreBranch, error) {

	return models.DrugStoreBranch{}, nil
}

func (d *drugStoreBranchRepo) GetList(ctx context.Context, request models.GetListRequest) (models.DrugStoreBranchsResponse, error) {

	return models.DrugStoreBranchsResponse{}, nil
}

func (d *drugStoreBranchRepo) Update(ctx context.Context, request models.UpdateDrugStoreBranch) (string, error) {

	return "", nil
}

func (d *drugStoreBranchRepo) Delete(ctx context.Context, id string) error {

	return nil
}
