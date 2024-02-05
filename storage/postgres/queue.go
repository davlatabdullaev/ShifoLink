package postgres

import (
	"context"
	"shifolink/api/models"
	"shifolink/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type queueRepo struct {
	pool *pgxpool.Pool
}

func NewQueueRepo(pool *pgxpool.Pool) storage.IQueueRepo {
	return &queueRepo{
		pool: pool,
	}
}

func (q *queueRepo) Create(ctx context.Context, request models.CreateQueue) (string, error) {

	return "", nil
}

func (q *queueRepo) Get(ctx context.Context, request models.PrimaryKey) (models.Queue, error) {

	return models.Queue{}, nil
}

func (q *queueRepo) GetList(ctx context.Context, request models.GetListRequest) (models.QueuesResponse, error) {

	return models.QueuesResponse{}, nil
}

func (q *queueRepo) Update(ctx context.Context, request models.UpdateQueue) (string, error) {

	return "", nil
}

func (q *queueRepo) Delete(ctx context.Context, id string) error {

	return nil
}
