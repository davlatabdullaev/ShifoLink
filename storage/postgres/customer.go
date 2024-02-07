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

type customerRepo struct {
	pool *pgxpool.Pool
}

func NewCustomerRepo(pool *pgxpool.Pool) storage.ICustomerRepo {
	return &customerRepo{
		pool: pool,
	}
}

func (c *customerRepo) Create(ctx context.Context, request models.CreateCustomer) (string, error) {

	id := uuid.New()

	query := `insert into customer (id, first_name, last_name, email, password, phone, gender, birth_date, age, address) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := c.pool.Exec(ctx, query,
		id,
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
		log.Println("error while inserting customer", err.Error())
		return "", err
	}

	return id.String(), nil

}

func (c *customerRepo) Get(ctx context.Context, request models.PrimaryKey) (models.Customer, error) {

	customer := models.Customer{}

	query := `select 
	 id,
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
	 from customer where deleted_at is null and id = $1`

	row := c.pool.QueryRow(ctx, query, request.ID)

	err := row.Scan(
		&customer.ID,
		&customer.FirstName,
		&customer.LastName,
		&customer.Email,
		&customer.Password,
		&customer.Phone,
		&customer.Gender,
		&customer.BirthDate,
		&customer.Age,
		&customer.Address,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)

	if err != nil {
		log.Println("error while selecting customer", err.Error())
		return models.Customer{}, err
	}

	return customer, nil

}

func (c *customerRepo) GetList(ctx context.Context, request models.GetListRequest) (models.CustomersResponse, error) {

	var (
		customers         = []models.Customer{}
		count             = 0
		query, countQuery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)

	countQuery = `select count(1) from customer where deleted_at is null`

	if search != "" {
		countQuery += fmt.Sprintf(` and (first_name ilike '%%%s%%' or last_name ilike '%%%s%%')`, search, search)
	}
	if err := c.pool.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		fmt.Println("error is while selecting count", err.Error())
		return models.CustomersResponse{}, err
	}

	query = `select 
	 id,
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
	 updated_at from customer where deleted_at is null`

	if search != "" {
		query += fmt.Sprintf(` and (first_name ilike '%%%s%%' or last_name ilike '%%%s%%')`, search, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := c.pool.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting customer", err.Error())
		return models.CustomersResponse{}, err
	}

	for rows.Next() {
		customer := models.Customer{}
		if err = rows.Scan(
			&customer.ID,
			&customer.FirstName,
			&customer.LastName,
			&customer.Email,
			&customer.Password,
			&customer.Phone,
			&customer.Gender,
			&customer.BirthDate,
			&customer.Age,
			&customer.Address,
			&customer.CreatedAt,
			&customer.UpdatedAt,
		); err != nil {
			fmt.Println("error is while scanning customer data", err.Error())
			return models.CustomersResponse{}, err
		}

		customers = append(customers, customer)

	}

	return models.CustomersResponse{
		Customers: customers,
		Count:     count,
	}, nil
}

func (c *customerRepo) Update(ctx context.Context, request models.UpdateCustomer) (string, error) {

	query := `update customer set
    first_name = $1,
    last_name = $2, 
	email = $3,
	phone = $4,
    address = $5,
	updated_at = $6
   where id = $7  
   `

	_, err := c.pool.Exec(ctx, query,
		request.FirstName,
		request.LastName,
		request.Email,
		request.Phone,
		request.Address,
		time.Now(),
		request.ID)
	if err != nil {
		log.Println("error while updating customer data...", err.Error())
		return "", err
	}

	return request.ID, nil

}

func (c *customerRepo) Delete(ctx context.Context, id string) error {

	query := `
	update customer
	 set deleted_at = $1
	  where id = $2
	`

	_, err := c.pool.Exec(ctx, query, time.Now(), id)
	if err != nil {
		log.Println("error while deleting customer by id", err.Error())
		return err
	}

	return nil
}

func (c *customerRepo) UpdatePassword(ctx context.Context, request models.UpdateCustomerPassword) error {

	query := `
		update customer
				set password = $1, updated_at = now()
					where id = $2`

	if _, err := c.pool.Exec(ctx, query, request.NewPassword, request.ID); err != nil {
		fmt.Println("error while updating password for customer", err.Error())
		return err
	}

	return nil
}
