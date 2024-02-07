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

type doctorRepo struct {
	pool *pgxpool.Pool
}

func NewDoctorRepo(pool *pgxpool.Pool) storage.IDoctorRepo {
	return &doctorRepo{
		pool: pool,
	}
}

func (d *doctorRepo) Create(ctx context.Context, request models.CreateDoctor) (string, error) {

	id := uuid.New()

	query := `insert into doctor (
		id, 
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
		working_time,
		status) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	rowsAffected, err := d.pool.Exec(ctx, query,
		id,
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
		request.WorkingTime,
		request.Status,
	)

	if r := rowsAffected.RowsAffected(); r == 0 {
		log.Println("error is while rows affected ", err.Error())
		return "", err
	}

	if err != nil {
		log.Println("error while inserting doctor", err.Error())
		return "", err
	}

	return id.String(), nil
}

func (d *doctorRepo) Get(ctx context.Context, request models.PrimaryKey) (models.Doctor, error) {

	var updatedAt = sql.NullTime{}

	doctor := models.Doctor{}

	query := `select 
	 id,
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
	 working_time,
	 status,
	 created_at, 
	 updated_at 
	 from doctor where deleted_at is null and id = $1`

	row := d.pool.QueryRow(ctx, query, request.ID)

	err := row.Scan(
		&doctor.ID,
		&doctor.DoctorTypeID,
		&doctor.FirstName,
		&doctor.LastName,
		&doctor.Email,
		&doctor.Password,
		&doctor.Phone,
		&doctor.Gender,
		&doctor.BirthDate,
		&doctor.Age,
		&doctor.Address,
		&doctor.WorkingTime,
		&doctor.Status,
		&doctor.CreatedAt,
		&updatedAt,
	)

	if err != nil {
		log.Println("error while selecting doctor", err.Error())
		return models.Doctor{}, err
	}

	if updatedAt.Valid {
		doctor.UpdatedAt = updatedAt.Time
	}

	return doctor, nil

}

func (d *doctorRepo) GetList(ctx context.Context, request models.GetListRequest) (models.DoctorsResponse, error) {

	var (
		updatedAt         = sql.NullTime{}
		doctors           = []models.Doctor{}
		count             = 0
		query, countQuery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)

	countQuery = `select count(1) from doctor where deleted_at is null`

	if search != "" {
		countQuery += fmt.Sprintf(` and (first_name ilike '%%%s%%' or last_name ilike '%%%s%%')`, search, search)
	}
	if err := d.pool.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		fmt.Println("error is while selecting count", err.Error())
		return models.DoctorsResponse{}, err
	}

	query = `select 
	 id,
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
	 working_time,
	 status,
	 created_at, 
	 updated_at from doctor where deleted_at is null`

	if search != "" {
		query += fmt.Sprintf(` and (first_name ilike '%%%s%%' or last_name ilike '%%%s%%')`, search, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := d.pool.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting doctor", err.Error())
		return models.DoctorsResponse{}, err
	}

	for rows.Next() {
		doctor := models.Doctor{}
		if err = rows.Scan(
			&doctor.ID,
			&doctor.DoctorTypeID,
			&doctor.FirstName,
			&doctor.LastName,
			&doctor.Email,
			&doctor.Password,
			&doctor.Phone,
			&doctor.Gender,
			&doctor.BirthDate,
			&doctor.Age,
			&doctor.Address,
			&doctor.WorkingTime,
			&doctor.Status,
			&doctor.CreatedAt,
			&updatedAt,
		); err != nil {
			fmt.Println("error is while scanning doctor data", err.Error())
			return models.DoctorsResponse{}, err
		}

		if updatedAt.Valid {
			doctor.UpdatedAt = updatedAt.Time
		}

		doctors = append(doctors, doctor)
	}

	return models.DoctorsResponse{
		Doctors: doctors,
		Count:   count,
	}, nil
}

func (d *doctorRepo) Update(ctx context.Context, request models.UpdateDoctor) (string, error) {

	query := `update doctor set
	doctor_type_id = $1,
    first_name = $2,
    last_name = $3, 
	email = $4,
	phone = $5,
    address = $6,
	working_time = $7,
	status = $8,
	updated_at = $9
   where id = $10
   `

	rowsAffected, err := d.pool.Exec(ctx, query,
		request.DoctorTypeID,
		request.FirstName,
		request.LastName,
		request.Email,
		request.Phone,
		request.Address,
		request.WorkingTime,
		request.Status,
		time.Now(),
		request.ID)

	if r := rowsAffected.RowsAffected(); r == 0 {
		log.Println("error is while rows affected ", err.Error())
		return "", err
	}

	if err != nil {
		log.Println("error while updating doctor data...", err.Error())
		return "", err
	}

	return request.ID, nil

}

func (d *doctorRepo) Delete(ctx context.Context, id string) error {

	query := `
	update doctor
	 set deleted_at = $1
	  where id = $2
	`

	rowsAffected, err := d.pool.Exec(ctx, query, time.Now(), id)

	if r := rowsAffected.RowsAffected(); r == 0 {
		log.Println("error is while rows affected ", err.Error())
		return err
	}

	if err != nil {
		log.Println("error while deleting doctor by id", err.Error())
		return err
	}

	return nil
}

func (d *doctorRepo) UpdatePassword(ctx context.Context, request models.UpdateDoctorPassword) error {

	query := `
		update doctor
				set password = $1, updated_at = now()
					where id = $2`

	rowsAffected, err := d.pool.Exec(ctx, query, request.NewPassword, request.ID)

	if r := rowsAffected.RowsAffected(); r == 0 {
		log.Println("error is while rows affected ", err.Error())
		return err
	}

	if err != nil {
		fmt.Println("error while updating password for doctor", err.Error())
		return err
	}

	return nil
}
