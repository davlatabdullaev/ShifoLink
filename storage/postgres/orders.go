package postgres

import (
	"context"
	"shifolink/api/models"
	"shifolink/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ordersRepo struct {
	pool *pgxpool.Pool
}

func NewOrdersRepo(pool *pgxpool.Pool) storage.IOrdersRepo {
	return &ordersRepo{
		pool: pool,
	}
}

func (o *ordersRepo) Create(ctx context.Context, request models.CreateOrders) (string, error) {

	return "", nil
}

func (o *ordersRepo) Get(ctx context.Context, request models.PrimaryKey) (models.Orders, error) {

	return models.Orders{}, nil
}

func (o *ordersRepo) GetList(ctx context.Context, request models.GetListRequest) (models.OrdersResponse, error) {

	return models.OrdersResponse{}, nil
}

func (o *ordersRepo) Update(ctx context.Context, request models.UpdateOrders) (string, error) {

	return "", nil
}

func (o *ordersRepo) Delete(ctx context.Context, id string) error {

	return nil
}
