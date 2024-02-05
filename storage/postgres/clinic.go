package postgres

import (
	"context"
	"shifolink/api/models"
	"shifolink/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type clinicRepo struct {
	pool *pgxpool.Pool
}

func NewClinicRepo(pool *pgxpool.Pool) storage.IClinicRepo {
	return &clinicRepo{
		pool: pool,
	}
}

func (c *clinicRepo) Create(ctx context.Context, request models.CreateClinic) (string, error) {

	return "", nil
}

func (c *clinicRepo) Get(ctx context.Context, request models.PrimaryKey) (models.Clinic, error) {

	return models.Clinic{}, nil
}

func (c *clinicRepo) GetList(ctx context.Context, request models.GetListRequest) (models.ClinicsResponse, error) {

	return models.ClinicsResponse{}, nil
}

func (c *clinicRepo) Update(ctx context.Context, request models.UpdateClinic) (string, error) {

	return "", nil
}

func (c *clinicRepo) Delete(ctx context.Context, id string) error {

	return nil
}
