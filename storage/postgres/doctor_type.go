package postgres

import (
	"context"
	"shifolink/api/models"
	"shifolink/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type doctorTypeRepo struct {
	pool *pgxpool.Pool
}

func NewDoctorTypeRepo(pool *pgxpool.Pool) storage.IDoctorTypeRepo {
	return &doctorTypeRepo{
		pool: pool,
	}
}

func (d *doctorTypeRepo) Create(ctx context.Context, request models.CreateDoctorType) (string, error) {

	return "", nil
}

func (d *doctorTypeRepo) Get(ctx context.Context, request models.PrimaryKey) (models.DoctorType, error) {

	return models.DoctorType{}, nil
}

func (d *doctorTypeRepo) GetList(ctx context.Context, request models.GetListRequest) (models.DoctorTypesResponse, error) {

	return models.DoctorTypesResponse{}, nil
}

func (d *doctorTypeRepo) Update(ctx context.Context, request models.UpdateDoctorType) (string, error) {

	return "", nil
}

func (d *doctorTypeRepo) Delete(ctx context.Context, id string) error {

	return nil
}
