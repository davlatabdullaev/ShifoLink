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

type authorRepo struct {
	pool *pgxpool.Pool
}

func NewAuthorRepo(pool *pgxpool.Pool) storage.IAuthorRepo {
	return &authorRepo{
		pool: pool,
	}
}

func (a *authorRepo) Create(ctx context.Context, request models.CreateAuthor) (string, error) {

	id := uuid.New()

	query := `insert into author (id, first_name, last_name, email, password, phone, gender, birth_date, age, address) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	rowsAffected, err := a.pool.Exec(ctx, query,
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

	if r := rowsAffected.RowsAffected(); r == 0 {
		log.Println("error is while rows affected ", err.Error())
		return "", err
	}

	if err != nil {
		log.Println("error while inserting author", err.Error())
		return "", err
	}

	return id.String(), nil

}

func (a *authorRepo) Get(ctx context.Context, request models.PrimaryKey) (models.Author, error) {

	var updatedAt = sql.NullTime{}

	author := models.Author{}

	query := `select 
	 id,
	 first_name, 
	 last_name, 
	 email, 
	 password, 
	 phone, 
	 gender, 
	 birth_date::text, 
	 age, 
	 address, 
	 created_at, 
	 updated_at 
	 from author where deleted_at is null and id = $1`

	row := a.pool.QueryRow(ctx, query, request.ID)

	err := row.Scan(
		&author.ID,
		&author.FirstName,
		&author.LastName,
		&author.Email,
		&author.Password,
		&author.Phone,
		&author.Gender,
		&author.BirthDate,
		&author.Age,
		&author.Address,
		&author.CreatedAt,
		&updatedAt,
	)

	if err != nil {
		log.Println("error while selecting author", err.Error())
		return models.Author{}, err
	}

	if updatedAt.Valid {
		author.UpdatedAt = updatedAt.Time
	}

	return author, nil
}

func (a *authorRepo) GetList(ctx context.Context, request models.GetListRequest) (models.AuthorsResponse, error) {

	var (
		updatedAt         = sql.NullTime{}
		authors           = []models.Author{}
		count             = 0
		query, countQuery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)

	countQuery = `select count(1) from author where deleted_at is null`

	if search != "" {
		countQuery += fmt.Sprintf(` and (first_name ilike '%%%s%%' or last_name ilike '%%%s%%')`, search, search)
	}
	if err := a.pool.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		fmt.Println("error is while selecting count", err.Error())
		return models.AuthorsResponse{}, err
	}

	query = `select 
	 id,
	 first_name, 
	 last_name, 
	 email, 
	 password, 
	 phone, 
	 gender, 
	 birth_date::text, 
	 age, 
	 address, 
	 created_at, 
	 updated_at from author where deleted_at is null`

	if search != "" {
		query += fmt.Sprintf(` and (first_name ilike '%%%s%%' or last_name ilike '%%%s%%')`, search, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := a.pool.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting author", err.Error())
		return models.AuthorsResponse{}, err
	}

	for rows.Next() {
		author := models.Author{}
		if err = rows.Scan(
			&author.ID,
			&author.FirstName,
			&author.LastName,
			&author.Email,
			&author.Password,
			&author.Phone,
			&author.Gender,
			&author.BirthDate,
			&author.Age,
			&author.Address,
			&author.CreatedAt,
			&updatedAt,
		); err != nil {
			fmt.Println("error is while scanning author data", err.Error())
			return models.AuthorsResponse{}, err
		}

		if updatedAt.Valid {
			author.UpdatedAt = updatedAt.Time
		}

		authors = append(authors, author)

	}

	return models.AuthorsResponse{
		Authors: authors,
		Count:   count,
	}, nil
}

func (a *authorRepo) Update(ctx context.Context, request models.UpdateAuthor) (string, error) {

	query := `update author set
    first_name = $1,
    last_name = $2, 
	email = $3,
	phone = $4,
    address = $5,
	updated_at = $6
   where id = $7  
   `

	rowsAffected, err := a.pool.Exec(ctx, query,
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
		log.Println("error while updating author data...", err.Error())
		return "", err
	}

	return request.ID, nil
}

func (a *authorRepo) Delete(ctx context.Context, id string) error {

	query := `
	update author
	 set deleted_at = $1
	  where id = $2
	`

	rowsAffected, err := a.pool.Exec(ctx, query, time.Now(), id)

	if r := rowsAffected.RowsAffected(); r == 0 {
		log.Println("error is while rows affected ", err.Error())
		return err
	}

	if err != nil {
		log.Println("error while deleting author by id", err.Error())
		return err
	}

	return nil
}

func (a *authorRepo) GetPassword(ctx context.Context, id string) (string, error) {
	password := ""

	query := `
		select password from author 
		                where id = $1`

	if err := a.pool.QueryRow(ctx, query, id).Scan(&password); err != nil {
		fmt.Println("Error while scanning password from author", err.Error())
		return "", err
	}

	return password, nil
}

func (a *authorRepo) UpdatePassword(ctx context.Context, request models.UpdateAuthorPassword) error {

	query := `
		update author 
				set password = $1, updated_at = now()
					where id = $2`

	rowsAffected, err := a.pool.Exec(ctx, query, request.NewPassword, request.ID)

	if r := rowsAffected.RowsAffected(); r == 0 {
		log.Println("error is while rows affected ", err.Error())
		return err
	}

	if err != nil {
		log.Println("error while updating password for author", err.Error())
		return err
	}

	return nil
}
