package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"shifolink/api/models"
	"shifolink/pkg/check"
	"shifolink/storage"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type clinicAdminRepo struct {
	pool *pgxpool.Pool
}

func NewClinicAdminRepo(pool *pgxpool.Pool) storage.IClinicAdminRepo {
	return &clinicAdminRepo{
		pool: pool,
	}
}

func (c *clinicAdminRepo) Create(ctx context.Context, request models.CreateClinicAdmin) (string, error) {

	id := uuid.New()

	query := `insert into clinic_admin
	 (id, 
	  clinic_branch_id,
	  doctor_type_id,
	  first_name, 
	  last_name, 
	  email, 
	  password, 
	  phone, 
	  gender, 
	  birth_date, 
	  age, 
	  address) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	rowsAffected, err := c.pool.Exec(ctx, query,
		id,
		request.ClinicBranchID,
		request.DoctorTypeID,
		request.FirstName,
		request.LastName,
		request.Email,
		request.Password,
		request.Phone,
		request.Gender,
		request.BirthDate,
		check.CalculateAge(request.BirthDate),
		request.Address,
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

func (c *clinicAdminRepo) Get(ctx context.Context, request models.PrimaryKey) (models.ClinicAdmin, error) {

	var updatedAt = sql.NullTime{}

	clinicAdmin := models.ClinicAdmin{}

	query := `select 
	 id,
	 clinic_branch_id,
	 doctor_type_id,
	 first_name, 
	 last_name, 
	 email, 
	 password, 
	 phone, 
	 gender, 
	 birth_date, 
	 age, 
	 address, 
	 created_at, 
	 updated_at 
	 from clinic_admin where deleted_at is null and id = $1`

	row := c.pool.QueryRow(ctx, query, request.ID)

	err := row.Scan(
		&clinicAdmin.ID,
		&clinicAdmin.ClinicBranchID,
		&clinicAdmin.DoctorTypeID,
		&clinicAdmin.FirstName,
		&clinicAdmin.LastName,
		&clinicAdmin.Email,
		&clinicAdmin.Password,
		&clinicAdmin.Phone,
		&clinicAdmin.Gender,
		&clinicAdmin.BirthDate,
		&clinicAdmin.Age,
		&clinicAdmin.Address,
		&clinicAdmin.CreatedAt,
		&updatedAt,
	)

	if err != nil {
		log.Println("error while selecting clinic admin", err.Error())
		return models.ClinicAdmin{}, err
	}

	if updatedAt.Valid {
		clinicAdmin.UpdatedAt = updatedAt.Time
	}

	return clinicAdmin, nil
}

func (c *clinicAdminRepo) GetList(ctx context.Context, request models.GetListRequest) (models.ClinicAdminsResponse, error) {

	var (
		updatedAt         = sql.NullTime{}
		clinicAdmins      = []models.ClinicAdmin{}
		count             = 0
		query, countQuery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)

	countQuery = `select count(1) from clinic_admin where deleted_at is null`

	if search != "" {
		countQuery += fmt.Sprintf(` and (first_name ilike '%%%s%%' or last_name ilike '%%%s%%')`, search, search)
	}
	if err := c.pool.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		fmt.Println("error is while selecting count", err.Error())
		return models.ClinicAdminsResponse{}, err
	}

	query = `select 
	 id,
	 clinic_branch_id,
	 doctor_type_id,
	 first_name, 
	 last_name, 
	 email, 
	 password, 
	 phone, 
	 gender, 
	 birth_date, 
	 age, 
	 address, 
	 created_at, 
	 updated_at from clinic_admin where deleted_at is null`

	if search != "" {
		query += fmt.Sprintf(` and (first_name ilike '%%%s%%' or last_name ilike '%%%s%%')`, search, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := c.pool.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting clinic admin", err.Error())
		return models.ClinicAdminsResponse{}, err
	}

	for rows.Next() {
		clinicAdmin := models.ClinicAdmin{}
		if err = rows.Scan(
			&clinicAdmin.ID,
			&clinicAdmin.ClinicBranchID,
			&clinicAdmin.DoctorTypeID,
			&clinicAdmin.FirstName,
			&clinicAdmin.LastName,
			&clinicAdmin.Email,
			&clinicAdmin.Password,
			&clinicAdmin.Phone,
			&clinicAdmin.Gender,
			&clinicAdmin.BirthDate,
			&clinicAdmin.Age,
			&clinicAdmin.Address,
			&clinicAdmin.CreatedAt,
			&updatedAt,
		); err != nil {
			fmt.Println("error is while scanning clinic admin data", err.Error())
			return models.ClinicAdminsResponse{}, err
		}

		if updatedAt.Valid {
			clinicAdmin.UpdatedAt = updatedAt.Time
		}

		clinicAdmins = append(clinicAdmins, clinicAdmin)

	}

	return models.ClinicAdminsResponse{
		ClinicAdmins: clinicAdmins,
		Count:        count,
	}, nil
}

func (c *clinicAdminRepo) Update(ctx context.Context, request models.UpdateClinicAdmin) (string, error) {

	query := `update author set
	clinic_branch_id = $1,
	doctor_type_id = $2,
    first_name = $3,
    last_name = $4, 
	email = $5,
	phone = $6,
    address = $7,
	updated_at = $8
   where id = $9  
   `

	rowsAffected, err := c.pool.Exec(ctx, query,
		request.ClinicBranchID,
		request.DoctorTypeID,
		request.FirstName,
		request.LastName,
		request.Email,
		request.Phone,
		request.Address,
		time.Now(),
		request.ID)

	if r := rowsAffected.RowsAffected(); r == 0 {
		log.Println("error is while rows affected ", err.Error())
		return "", err
	}

	if err != nil {
		log.Println("error while updating clinic admin data...", err.Error())
		return "", err
	}

	return request.ID, nil
}

func (c *clinicAdminRepo) Delete(ctx context.Context, id string) error {

	query := `
	update clinic_admin
	 set deleted_at = $1
	  where id = $2
	`

	rowsAffected, err := c.pool.Exec(ctx, query, time.Now(), id)

	if r := rowsAffected.RowsAffected(); r == 0 {
		log.Println("error is while rows affected ", err.Error())
		return err
	}

	if err != nil {
		log.Println("error while deleting clinic admin by id", err.Error())
		return err
	}

	return nil
}

func (c *clinicAdminRepo) UpdatePassword(ctx context.Context, request models.UpdateClinicAdminPassword) error {

	query := `
		update clinic_admin 
				set password = $1, updated_at = now()
					where id = $2`

	rowsAffected, err := c.pool.Exec(ctx, query, request.NewPassword, request.ID)

	if r := rowsAffected.RowsAffected(); r == 0 {
		log.Println("error is while rows affected ", err.Error())
		return err
	}

	if err != nil {
		fmt.Println("error while updating password for clinic admin", err.Error())
		return err
	}

	return nil
}
