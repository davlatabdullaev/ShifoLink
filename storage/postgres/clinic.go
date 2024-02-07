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

type clinicRepo struct {
	pool *pgxpool.Pool
}

func NewClinicRepo(pool *pgxpool.Pool) storage.IClinicRepo {
	return &clinicRepo{
		pool: pool,
	}
}

func (c *clinicRepo) Create(ctx context.Context, request models.CreateClinic) (string, error) {

	id := uuid.New()

	query := `insert into clinic
	 (id, 
	  name,
	  description) 
	  values ($1, $2, $3)`

	_, err := c.pool.Exec(ctx, query,
		id,
		request.Name,
		request.Description,
	)
	if err != nil {
		log.Println("error while inserting clinic", err.Error())
		return "", err
	}

	return id.String(), nil
}

func (c *clinicRepo) Get(ctx context.Context, request models.PrimaryKey) (models.Clinic, error) {

	var updatedAt = sql.NullTime{}

	clinic := models.Clinic{}

	query := `select 
	 id,
	 name,
	 description,
	 created_at,
	 updated_at
	 from clinic where deleted_at is null and id = $1`

	row := c.pool.QueryRow(ctx, query, request.ID)

	err := row.Scan(
		&clinic.ID,
		&clinic.Name,
		&clinic.Description,
		&clinic.CreatedAt,
		&updatedAt,
	)

	if err != nil {
		log.Println("error while selecting clinic ", err.Error())
		return models.Clinic{}, err
	}

	if updatedAt.Valid {
		clinic.UpdatedAt = updatedAt.Time
	}

	return clinic, nil
}

func (c *clinicRepo) GetList(ctx context.Context, request models.GetListRequest) (models.ClinicsResponse, error) {

	var (
		updatedAt         = sql.NullTime{}
		clinics           = []models.Clinic{}
		count             = 0
		query, countQuery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)

	countQuery = `select count(1) from clinic where deleted_at is null`

	if search != "" {
		countQuery += fmt.Sprintf(` and (name ilike '%%%s%%')`, search)
	}
	if err := c.pool.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		fmt.Println("error is while selecting count", err.Error())
		return models.ClinicsResponse{}, err
	}

	query = `select 
	 id,
	 name,
	 description,
	 created_at, 
	 updated_at from clinic where deleted_at is null`

	if search != "" {
		query += fmt.Sprintf(` and (name ilike '%%%s%%')`, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := c.pool.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting clinic ", err.Error())
		return models.ClinicsResponse{}, err
	}

	for rows.Next() {
		clinic := models.Clinic{}
		if err = rows.Scan(
			&clinic.ID,
			&clinic.Name,
			&clinic.Description,
			&clinic.CreatedAt,
			&updatedAt,
		); err != nil {
			fmt.Println("error is while scanning clinic data", err.Error())
			return models.ClinicsResponse{}, err
		}

		if updatedAt.Valid {
			clinic.UpdatedAt = updatedAt.Time
		}

		clinics = append(clinics, clinic)

	}

	return models.ClinicsResponse{
		Clinics: clinics,
		Count:   count,
	}, nil
}

func (c *clinicRepo) Update(ctx context.Context, request models.UpdateClinic) (string, error) {

	query := `update clinic set
	name = $1,
	description = $2,
    updated_at = $3 
	 where id = $4  
   `

	_, err := c.pool.Exec(ctx, query,
		request.Name,
		request.Description,
		time.Now(),
		request.ID)
	if err != nil {
		log.Println("error while updating clinic data...", err.Error())
		return "", err
	}

	return request.ID, nil
}

func (c *clinicRepo) Delete(ctx context.Context, id string) error {

	query := `
	update clinic
	 set deleted_at = $1
	  where id = $2
	`

	_, err := c.pool.Exec(ctx, query, time.Now(), id)
	if err != nil {
		log.Println("error while deleting clinic by id", err.Error())
		return err
	}

	return nil
}
