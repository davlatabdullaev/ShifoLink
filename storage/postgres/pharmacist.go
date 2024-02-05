package postgres

import (
	"context"
	"shifolink/api/models"
	"shifolink/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type pharmacistRepo struct {
	pool *pgxpool.Pool
}

func NewPharmacistRepo(pool *pgxpool.Pool) storage.IPharmacistRepo {
	return &pharmacistRepo{
		pool: pool,
	}
}

func (p *pharmacistRepo) Create(ctx context.Context, request models.CreatePharmacist) (string, error) {

	return "", nil
}

func (p *pharmacistRepo) Get(ctx context.Context, request models.PrimaryKey) (models.Pharmacist, error) {

	return models.Pharmacist{}, nil
}

func (p *pharmacistRepo) GetList(ctx context.Context, request models.GetListRequest) (models.PharmacistsResponse, error) {

	return models.PharmacistsResponse{}, nil
}

func (p *pharmacistRepo) Update(ctx context.Context, request models.UpdatePharmacist) (string, error) {

	return "", nil
}

func (p *pharmacistRepo) Delete(ctx context.Context, id string) error {

	return nil
}

func (p *pharmacistRepo) UpdatePassword(ctx context.Context, id string) error {

	return nil
}
