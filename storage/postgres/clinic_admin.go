package postgres

import (
	"context"
	"shifolink/api/models"
	"shifolink/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type clinicAdminRepo struct {
	pool *pgxpool.Pool
}

func NewClinicAdminRepo(pool *pgxpool.Pool) storage.IClinicAdminRepo {
	return &clinicAdminRepo{
		pool: pool,
	}
}

func (c *clinicAdminRepo) Create(ctx context.Context, request models.CreateClinicAdmin) (string, error) {

	return "", nil
}

func (c *clinicAdminRepo) Get(ctx context.Context, request models.PrimaryKey) (models.ClinicAdmin, error) {

	return models.ClinicAdmin{}, nil
}

func (c *clinicAdminRepo) GetList(ctx context.Context, request models.GetListRequest) (models.ClinicAdminsResponse, error) {

	return models.ClinicAdminsResponse{}, nil
}

func (c *clinicAdminRepo) Update(ctx context.Context, request models.UpdateClinicAdmin) (string, error) {

	return "", nil
}

func (c *clinicAdminRepo) Delete(ctx context.Context, id string) error {

	return nil
}

func (c *clinicAdminRepo) UpdatePassword(ctx context.Context, id string) error {

	return nil
}
