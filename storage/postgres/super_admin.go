package postgres

import (
	"context"
	"shifolink/api/models"
	"shifolink/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type superAdminRepo struct {
	pool *pgxpool.Pool
}

func NewSuperAdminRepo(pool *pgxpool.Pool) storage.ISuperAdminRepo {
	return &superAdminRepo{
		pool: pool,
	}
}

func (s *superAdminRepo) Create(ctx context.Context, request models.CreateSuperAdmin) (string, error) {

	return "", nil
}

func (s *superAdminRepo) Get(ctx context.Context, request models.PrimaryKey) (models.SuperAdmin, error) {

	return models.SuperAdmin{}, nil
}

func (s *superAdminRepo) GetList(ctx context.Context, request models.GetListRequest) (models.SuperAdminsResponse, error) {

	return models.SuperAdminsResponse{}, nil
}

func (s *superAdminRepo) Update(ctx context.Context, request models.UpdateSuperAdmin) (string, error) {

	return "", nil
}

func (s *superAdminRepo) Delete(ctx context.Context, id string) error {

	return nil
}

func (s *superAdminRepo) UpdatePassword(ctx context.Context, id string) error {

	return nil
}
