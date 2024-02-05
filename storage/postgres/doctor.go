package postgres

import (
	"context"
	"shifolink/api/models"
	"shifolink/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type doctorRepo struct {
	pool *pgxpool.Pool
}

func NewDoctorRepo(pool *pgxpool.Pool) storage.IDoctorRepo {
	return &doctorRepo{
		pool: pool,
	}
}

func (d *doctorRepo) Create(ctx context.Context, request models.CreateDoctor) (string, error) {

	return "", nil
}

func (d *doctorRepo) Get(ctx context.Context, request models.PrimaryKey) (models.Doctor, error) {

	return models.Doctor{}, nil
}

func (d *doctorRepo) GetList(ctx context.Context, request models.GetListRequest) (models.DoctorsResponse, error) {

	return models.DoctorsResponse{}, nil
}

func (d *doctorRepo) Update(ctx context.Context, request models.UpdateDoctor) (string, error) {

	return "", nil
}

func (d *doctorRepo) Delete(ctx context.Context, id string) error {

	return nil
}

func (d *doctorRepo) UpdatePassword(ctx context.Context, id string) error {

	return nil
}
