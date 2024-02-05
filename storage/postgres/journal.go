package postgres

import (
	"context"
	"shifolink/api/models"
	"shifolink/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type journalRepo struct {
	pool *pgxpool.Pool
}

func NewJournalRepo(pool *pgxpool.Pool) storage.IJournalRepo {
	return &journalRepo{
		pool: pool,
	}
}

func (j *journalRepo) Create(ctx context.Context, request models.CreateJournal) (string, error) {

	return "", nil
}

func (j *journalRepo) Get(ctx context.Context, request models.PrimaryKey) (models.Journal, error) {

	return models.Journal{}, nil
}

func (j *journalRepo) GetList(ctx context.Context, request models.GetListRequest) (models.JournalsResponse, error) {

	return models.JournalsResponse{}, nil
}

func (j *journalRepo) Update(ctx context.Context, request models.UpdateJournal) (string, error) {

	return "", nil
}

func (j *journalRepo) Delete(ctx context.Context, id string) error {

	return nil
}
