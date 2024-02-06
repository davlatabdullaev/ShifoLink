package postgres

import (
	"context"
	"fmt"
	"log"
	"shifolink/api/models"
	"shifolink/storage"
	"time"

	"github.com/google/uuid"
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

	id := uuid.New()

	query := `insert into doctor_type
	 (id, 
	  name,
	  description,
	  clinic_branch_id ) 
	  values ($1, $2, $3, $4)`

	_, err := d.pool.Exec(ctx, query,
		id,
		request.Name,
		request.Description,
		request.ClinicBranchID,
	)
	if err != nil {
		log.Println("error while inserting doctor type", err.Error())
		return "", err
	}

	return id.String(), nil

}

func (d *doctorTypeRepo) Get(ctx context.Context, request models.PrimaryKey) (models.DoctorType, error) {

	doctorType := models.DoctorType{}

	query := `select 
	 id,
	 name,
	 description,
	 clinic_branch_id,
	 created_at,
	 updated_at
	 from doctor_type where deleted_at is null and id = $1`

	row := d.pool.QueryRow(ctx, query, request.ID)

	err := row.Scan(
		&doctorType.ID,
		&doctorType.Name,
		&doctorType.Description,
		&doctorType.ClinicBranchID,
		&doctorType.CreatedAt,
		&doctorType.UpdatedAt,
	)

	if err != nil {
		log.Println("error while selecting doctor type ", err.Error())
		return models.DoctorType{}, err
	}

	return doctorType, nil
}

func (d *doctorTypeRepo) GetList(ctx context.Context, request models.GetListRequest) (models.DoctorTypesResponse, error) {

	var (
		doctorTypes       = []models.DoctorType{}
		count             = 0
		query, countQuery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)

	countQuery = `select count(1) from doctor_type where deleted_at is null`

	if search != "" {
		countQuery += fmt.Sprintf(` and (name ilike '%%%s%%')`, search)
	}
	if err := d.pool.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		fmt.Println("error is while selecting count", err.Error())
		return models.DoctorTypesResponse{}, err
	}

	query = `select 
	 id,
	 name,
	 description,
	 clinic_branch_id,
	 created_at, 
	 updated_at from doctor_type where deleted_at is null`

	if search != "" {
		query += fmt.Sprintf(` and (name ilike '%%%s%%')`, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := d.pool.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting doctor type ", err.Error())
		return models.DoctorTypesResponse{}, err
	}

	for rows.Next() {
		doctorType := models.DoctorType{}
		if err = rows.Scan(
			&doctorType.ID,
			&doctorType.Name,
			&doctorType.Description,
			&doctorType.ClinicBranchID,
			&doctorType.CreatedAt,
			&doctorType.UpdatedAt,
		); err != nil {
			fmt.Println("error is while scanning doctor_type data", err.Error())
			return models.DoctorTypesResponse{}, err
		}

		doctorTypes = append(doctorTypes, doctorType)

	}

	return models.DoctorTypesResponse{
		DoctorTypes: doctorTypes,
		Count:       count,
	}, nil
}

func (d *doctorTypeRepo) Update(ctx context.Context, request models.UpdateDoctorType) (string, error) {

	query := `update doctor_type set
	name = $1,
	description = $2,
	clinic_branch_id = $3
    updated_at = $4, 
	 where id = $5  
   `

	_, err := d.pool.Exec(ctx, query,
		request.Name,
		request.Description,
		request.ClinicBranchID,
		time.Now(),
		request.ID)
	if err != nil {
		log.Println("error while updating doctor type data...", err.Error())
		return "", err
	}

	return request.ID, nil

}

func (d *doctorTypeRepo) Delete(ctx context.Context, id string) error {

	query := `
	update doctor_type
	 set deleted_at = $1
	  where id = $2
	`

	_, err := d.pool.Exec(ctx, query, time.Now(), id)
	if err != nil {
		log.Println("error while deleting doctor type by id", err.Error())
		return err
	}

	return nil
}
