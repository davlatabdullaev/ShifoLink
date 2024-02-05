package postgres

import (
	"context"
	"shifolink/api/models"
	"shifolink/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type customerRepo struct {
	pool *pgxpool.Pool
}

func NewCustomerRepo(pool *pgxpool.Pool) storage.ICustomerRepo {
	return &customerRepo{
		pool: pool,
	}
}

func (c *customerRepo) Create(ctx context.Context, request models.CreateCustomer) (string, error) {

	return "", nil
}

func (c *customerRepo) Get(ctx context.Context, request models.PrimaryKey) (models.Customer, error) {

	return models.Customer{}, nil
}

func (c *customerRepo) GetList(ctx context.Context, request models.GetListRequest) (models.CustomersResponse, error) {

	return models.CustomersResponse{}, nil
}

func (c *customerRepo) Update(ctx context.Context, request models.UpdateCustomer) (string, error) {

	return "", nil
}

func (c *customerRepo) Delete(ctx context.Context, id string) error {

	return nil
}

func (c *customerRepo) UpdatePassword(ctx context.Context, id string) error {

	return nil
}
