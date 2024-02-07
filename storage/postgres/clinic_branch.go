package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"shifolink/api/models"
	"shifolink/storage"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type clinicBranchRepo struct {
	pool *pgxpool.Pool
}

func NewClinicBranchRepo(pool *pgxpool.Pool) storage.IClinicBranchRepo {
	return &clinicBranchRepo{
		pool: pool,
	}
}

func (c *clinicBranchRepo) Create(ctx context.Context, request models.CreateClinicBranch) (string, error) {

	id := uuid.New()

	query := `insert into clinic_branch
	 (id, 
	  clinic_id,
	  address,
	  phone,
	  working_time) 
	  values ($1, $2, $3, $4, $5)`

	rowsAffected, err := c.pool.Exec(ctx, query,
		id,
		request.ClinicID,
		request.Address,
		request.Phone,
		request.WorkingTime,
	)

	if r := rowsAffected.RowsAffected(); r == 0 {
		log.Println("error is while rows affected ", err.Error())
		return "", err
	}

	if err != nil {
		log.Println("error while inserting clinic_admin", err.Error())
		return "", err
	}

	return id.String(), nil
}

func (c *clinicBranchRepo) Get(ctx context.Context, request models.PrimaryKey) (models.ClinicBranch, error) {

	var updatedAt = sql.NullTime{}

	clinicBranch := models.ClinicBranch{}

	query := `select 
	 id,
	 clinic_id,
	 address,
	 phone,
	 working_time,
	 created_at,
	 updated_at
	 from clinic_branch where deleted_at is null and id = $1`

	row := c.pool.QueryRow(ctx, query, request.ID)

	err := row.Scan(
		&clinicBranch.ID,
		&clinicBranch.ClinicID,
		&clinicBranch.Address,
		&clinicBranch.Phone,
		&clinicBranch.WorkingTime,
		&clinicBranch.CreatedAt,
		&updatedAt,
	)

	if err != nil {
		log.Println("error while selecting clinic branch", err.Error())
		return models.ClinicBranch{}, err
	}

	if updatedAt.Valid {
		clinicBranch.UpdatedAt = updatedAt.Time
	}

	return clinicBranch, nil
}

func (c *clinicBranchRepo) GetList(ctx context.Context, request models.GetListRequest) (models.ClinicBranchsResponse, error) {

	var (
		updatedAt         = sql.NullTime{}
		clinicBranchs     = []models.ClinicBranch{}
		count             = 0
		query, countQuery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)

	countQuery = `select count(1) from clinic_branch where deleted_at is null`

	if search != "" {
		countQuery += fmt.Sprintf(` and (address ilike '%%%s%%' or phone ilike '%%%s%%')`, search, search)
	}
	if err := c.pool.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		fmt.Println("error is while selecting count", err.Error())
		return models.ClinicBranchsResponse{}, err
	}

	query = `select 
	 id,
	 clinic_id,
	 address,
	 phone,
	 working_time,
	 created_at, 
	 updated_at from clinic_branch where deleted_at is null`

	if search != "" {
		query += fmt.Sprintf(` and (address ilike '%%%s%%' or phone ilike '%%%s%%')`, search, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := c.pool.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting clinic branch", err.Error())
		return models.ClinicBranchsResponse{}, err
	}

	for rows.Next() {
		clinicBranch := models.ClinicBranch{}
		if err = rows.Scan(
			&clinicBranch.ID,
			&clinicBranch.ClinicID,
			&clinicBranch.Address,
			&clinicBranch.Phone,
			&clinicBranch.WorkingTime,
			&clinicBranch.CreatedAt,
			&updatedAt,
		); err != nil {
			fmt.Println("error is while scanning clinic branch data", err.Error())
			return models.ClinicBranchsResponse{}, err
		}

		if updatedAt.Valid {
			clinicBranch.UpdatedAt = updatedAt.Time
		}

		clinicBranchs = append(clinicBranchs, clinicBranch)

	}

	return models.ClinicBranchsResponse{
		ClinicBranchs: clinicBranchs,
		Count:         count,
	}, nil
}

func (c *clinicBranchRepo) Update(ctx context.Context, request models.UpdateClinicBranch) (string, error) {

	query := `update clinic_branch set
	clinic_id = $1,
	address = $2,
    phone = $3,
	working_time = $4,
    updated_at = $5 
	 where id = $6  
   `

	rowsAffected, err := c.pool.Exec(ctx, query,
		request.ClinicID,
		request.Address,
		request.Phone,
		request.WorkingTime,
		time.Now(),
		request.ID)

	if r := rowsAffected.RowsAffected(); r == 0 {
		log.Println("error is while rows affected ", err.Error())
		return "", err
	}

	if err != nil {
		log.Println("error while updating clinic branch data...", err.Error())
		return "", err
	}

	return request.ID, nil
}

func (c *clinicBranchRepo) Delete(ctx context.Context, id string) error {

	query := `
	update clinic_branch
	 set deleted_at = $1
	  where id = $2
	`

	rowsAffected, err := c.pool.Exec(ctx, query, time.Now(), id)

	if r := rowsAffected.RowsAffected(); r == 0 {
		log.Println("error is while rows affected ", err.Error())
		return err
	}

	if err != nil {
		log.Println("error while deleting clinic branch by id", err.Error())
		return err
	}

	return nil
}
