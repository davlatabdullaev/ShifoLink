package postgres

import (
	"context"
	"shifolink/api/models"
	"shifolink/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type orderDrugRepo struct {
	pool *pgxpool.Pool
}

func NewOrderDrugRepo(pool *pgxpool.Pool) storage.IOrderDrugRepo {
	return &orderDrugRepo{
		pool: pool,
	}
}

func (o *orderDrugRepo) Create(ctx context.Context, request models.CreateOrderDrug) (string, error) {

	return "", nil
}

func (o *orderDrugRepo) Get(ctx context.Context, request models.PrimaryKey) (models.OrderDrug, error) {

	return models.OrderDrug{}, nil
}

func (o *orderDrugRepo) GetList(ctx context.Context, request models.GetListRequest) (models.OrderDrugsResponse, error) {

	return models.OrderDrugsResponse{}, nil
}

func (o *orderDrugRepo) Update(ctx context.Context, request models.UpdateOrderDrug) (string, error) {

	return "", nil
}

func (o *orderDrugRepo) Delete(ctx context.Context, id string) error {

	return nil
}
