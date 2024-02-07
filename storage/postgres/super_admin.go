package postgres

import (
	"context"
	"fmt"
	"log"
	"shifolink/api/models"
	"shifolink/pkg/check"
	"shifolink/storage"
	"time"

	"github.com/google/uuid"
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

	id := uuid.New()

	query := `insert into super_admin (
		id, 
		clinic_id,
		drug_store_id,
		author_id, 
		first_name,
	    last_name, 
	    email, 
		password, 
		phone, 
		gender, 
		birth_date, 
		age, 
		address) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	_, err := s.pool.Exec(ctx, query,
		id,
		request.ClinicID,
		request.DrugStoreID,
		request.AuthorID,
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
	if err != nil {
		log.Println("error while inserting super_admin", err.Error())
		return "", err
	}

	return id.String(), nil
}

func (s *superAdminRepo) Get(ctx context.Context, request models.PrimaryKey) (models.SuperAdmin, error) {

	superAdmin := models.SuperAdmin{}

	query := `select 
	 id,
	 clinic_id,
	 drug_store_id,
	 author_id,
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
	 from super_admin where deleted_at is null and id = $1`

	row := s.pool.QueryRow(ctx, query, request.ID)

	err := row.Scan(
		&superAdmin.ID,
		&superAdmin.ClinicID,
		&superAdmin.DrugStoreID,
		&superAdmin.AuthorID,
		&superAdmin.FirstName,
		&superAdmin.LastName,
		&superAdmin.Email,
		&superAdmin.Password,
		&superAdmin.Phone,
		&superAdmin.Gender,
		&superAdmin.BirthDate,
		&superAdmin.Age,
		&superAdmin.Address,
		&superAdmin.CreatedAt,
		&superAdmin.UpdatedAt,
	)

	if err != nil {
		log.Println("error while selecting superAdmin", err.Error())
		return models.SuperAdmin{}, err
	}

	return superAdmin, nil

}

func (s *superAdminRepo) GetList(ctx context.Context, request models.GetListRequest) (models.SuperAdminsResponse, error) {

	var (
		superAdmins       = []models.SuperAdmin{}
		count             = 0
		query, countQuery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)

	countQuery = `select count(1) from super_admin where deleted_at is null`

	if search != "" {
		countQuery += fmt.Sprintf(` and (first_name ilike '%%%s%%' or last_name ilike '%%%s%%')`, search, search)
	}
	if err := s.pool.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		fmt.Println("error is while selecting count", err.Error())
		return models.SuperAdminsResponse{}, err
	}

	query = `select 
	 id,
	 clinic_id,
	 drug_store_id,
	 author_id,
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
	 updated_at from super_admin where deleted_at is null`

	if search != "" {
		query += fmt.Sprintf(` and (first_name ilike '%%%s%%' or last_name ilike '%%%s%%')`, search, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := s.pool.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting super admin", err.Error())
		return models.SuperAdminsResponse{}, err
	}

	for rows.Next() {
		superAdmin := models.SuperAdmin{}
		if err = rows.Scan(
			&superAdmin.ID,
			&superAdmin.ClinicID,
			&superAdmin.DrugStoreID,
			&superAdmin.AuthorID,
			&superAdmin.FirstName,
			&superAdmin.LastName,
			&superAdmin.Email,
			&superAdmin.Password,
			&superAdmin.Phone,
			&superAdmin.Gender,
			&superAdmin.BirthDate,
			&superAdmin.Age,
			&superAdmin.Address,
			&superAdmin.CreatedAt,
			&superAdmin.UpdatedAt,
		); err != nil {
			fmt.Println("error is while scanning super admin data", err.Error())
			return models.SuperAdminsResponse{}, err
		}

		superAdmins = append(superAdmins, superAdmin)

	}

	return models.SuperAdminsResponse{
		SuperAdmins: superAdmins,
		Count:       count,
	}, nil
}

func (s *superAdminRepo) Update(ctx context.Context, request models.UpdateSuperAdmin) (string, error) {

	query := `update super_admin set
	clinic_id = $1,
	drug_store_id = $2,
	author_id = $3,
    first_name = $4,
    last_name = $5, 
	email = $6,
	phone = $7,
    address = $8,
	updated_at = $9
   where id = $10
   `

	_, err := s.pool.Exec(ctx, query,
		request.ClinicID,
		request.DrugStoreID,
		request.AuthorID,
		request.FirstName,
		request.LastName,
		request.Email,
		request.Phone,
		request.Address,
		time.Now(),
		request.ID)
	if err != nil {
		log.Println("error while updating super admin data...", err.Error())
		return "", err
	}

	return request.ID, nil

}

func (s *superAdminRepo) Delete(ctx context.Context, id string) error {

	query := `
	update super_admin
	 set deleted_at = $1
	  where id = $2
	`

	_, err := s.pool.Exec(ctx, query, time.Now(), id)
	if err != nil {
		log.Println("error while deleting super admin by id", err.Error())
		return err
	}

	return nil

}

func (s *superAdminRepo) UpdatePassword(ctx context.Context, request models.UpdateSuperAdminPassword) error {

	query := `
		update super_admin 
				set password = $1, updated_at = now()
					where id = $2`

	if _, err := s.pool.Exec(ctx, query, request.NewPassword, request.ID); err != nil {
		fmt.Println("error while updating password for super_admin", err.Error())
		return err
	}

	return nil
}
