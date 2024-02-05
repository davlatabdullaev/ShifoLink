package postgres

import (
	"context"
	"shifolink/api/models"
	"shifolink/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type clinicBranchRepo struct {
	pool *pgxpool.Pool
}

func NewClinicBranchRepo(pool *pgxpool.Pool) storage.IClinicBranchRepo {
	return &clinicBranchRepo{
		pool: pool,
	}
}

func (c *clinicBranchRepo) Create(ctx context.Context, request models.CreateClinicBranch) (string, error) {

	return "", nil
}

func (c *clinicBranchRepo) Get(ctx context.Context, request models.PrimaryKey) (models.ClinicBranch, error) {

	return models.ClinicBranch{}, nil
}

func (c *clinicBranchRepo) GetList(ctx context.Context, request models.GetListRequest) (models.ClinicBranchsResponse, error) {

	return models.ClinicBranchsResponse{}, nil
}

func (c *clinicBranchRepo) Update(ctx context.Context, request models.UpdateClinicBranch) (string, error) {

	return "", nil
}

func (c *clinicBranchRepo) Delete(ctx context.Context, id string) error {

	return nil
}
