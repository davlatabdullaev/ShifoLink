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

type pharmacistRepo struct {
	pool *pgxpool.Pool
}

func NewPharmacistRepo(pool *pgxpool.Pool) storage.IPharmacistRepo {
	return &pharmacistRepo{
		pool: pool,
	}
}

func (p *pharmacistRepo) Create(ctx context.Context, request models.CreatePharmacist) (string, error) {

	id := uuid.New()

	query := `insert into pharmacist (
		id, 
		drug_store_branch_id, 
		first_name,
	    last_name, 
	    email, 
		password, 
		phone, 
		gender, 
		birth_date, 
		age, 
		address) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := p.pool.Exec(ctx, query,
		id,
		request.DrugStoreBranchID,
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
		log.Println("error while inserting pharmacist", err.Error())
		return "", err
	}

	return id.String(), nil

}

func (p *pharmacistRepo) Get(ctx context.Context, request models.PrimaryKey) (models.Pharmacist, error) {

	pharmacist := models.Pharmacist{}

	query := `select 
	 id,
	 drug_store_branch_id,
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
	 from doctor where deleted_at is null and id = $1`

	row := p.pool.QueryRow(ctx, query, request.ID)

	err := row.Scan(
		&pharmacist.ID,
		&pharmacist.DrugStoreBranchID,
		&pharmacist.FirstName,
		&pharmacist.LastName,
		&pharmacist.Email,
		&pharmacist.Password,
		&pharmacist.Phone,
		&pharmacist.Gender,
		&pharmacist.BirthDate,
		&pharmacist.Age,
		&pharmacist.Address,
		&pharmacist.CreatedAt,
		&pharmacist.UpdatedAt,
	)

	if err != nil {
		log.Println("error while selecting pharmacist", err.Error())
		return models.Pharmacist{}, err
	}

	return pharmacist, nil

}

func (p *pharmacistRepo) GetList(ctx context.Context, request models.GetListRequest) (models.PharmacistsResponse, error) {

	var (
		pharmacists       = []models.Pharmacist{}
		count             = 0
		query, countQuery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)

	countQuery = `select count(1) from pharmacist where deleted_at is null`

	if search != "" {
		countQuery += fmt.Sprintf(` and (first_name ilike '%%%s%%' or last_name ilike '%%%s%%')`, search, search)
	}
	if err := p.pool.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		fmt.Println("error is while selecting count", err.Error())
		return models.PharmacistsResponse{}, err
	}

	query = `select 
	 id,
	 drug_store_branch_id,
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
	 updated_at from pharmacist where deleted_at is null`

	if search != "" {
		query += fmt.Sprintf(` and (first_name ilike '%%%s%%' or last_name ilike '%%%s%%')`, search, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := p.pool.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting pharmacist", err.Error())
		return models.PharmacistsResponse{}, err
	}

	for rows.Next() {
		pharmacist := models.Pharmacist{}
		if err = rows.Scan(
			&pharmacist.ID,
			&pharmacist.DrugStoreBranchID,
			&pharmacist.FirstName,
			&pharmacist.LastName,
			&pharmacist.Email,
			&pharmacist.Password,
			&pharmacist.Phone,
			&pharmacist.Gender,
			&pharmacist.BirthDate,
			&pharmacist.Age,
			&pharmacist.Address,
			&pharmacist.CreatedAt,
			&pharmacist.UpdatedAt,
		); err != nil {
			fmt.Println("error is while scanning pharmacist data", err.Error())
			return models.PharmacistsResponse{}, err
		}

		pharmacists = append(pharmacists, pharmacist)

	}

	return models.PharmacistsResponse{
		Pharmacists: pharmacists,
		Count:       count,
	}, nil
}

func (p *pharmacistRepo) Update(ctx context.Context, request models.UpdatePharmacist) (string, error) {

	query := `update pharmacist set
	drug_store_branch_id = $1,
    first_name = $2,
    last_name = $3, 
	email = $4,
	phone = $5,
    address = $6,
	updated_at = $7
   where id = $8
   `

	_, err := p.pool.Exec(ctx, query,
		request.DrugStoreBranchID,
		request.FirstName,
		request.LastName,
		request.Email,
		request.Phone,
		request.Address,
		time.Now(),
		request.ID)
	if err != nil {
		log.Println("error while updating pharmacist data...", err.Error())
		return "", err
	}

	return request.ID, nil

}

func (p *pharmacistRepo) Delete(ctx context.Context, id string) error {

	query := `
	update pharmacist
	 set deleted_at = $1
	  where id = $2
	`

	_, err := p.pool.Exec(ctx, query, time.Now(), id)
	if err != nil {
		log.Println("error while deleting pharmacist by id", err.Error())
		return err
	}

	return nil
}

func (p *pharmacistRepo) UpdatePassword(ctx context.Context, id string) error {

	return nil
}
