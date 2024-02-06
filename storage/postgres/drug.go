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

type drugRepo struct {
	pool *pgxpool.Pool
}

func NewDrugRepo(pool *pgxpool.Pool) storage.IDrugRepo {
	return &drugRepo{
		pool: pool,
	}
}

func (d *drugRepo) Create(ctx context.Context, request models.CreateDrug) (string, error) {

	id := uuid.New()

	query := `insert into drug
	 (id, 
	  drug_store_branch_id,
	  name,
	  description,
	  count,
	  price,
	  date_of_manufacture,
	  best_before) 
	  values ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := d.pool.Exec(ctx, query,
		id,
		request.DrugStoreBranchID,
		request.Name,
		request.Description,
		request.Count,
		request.Price,
		request.DateOfManufacture,
		request.BestBefore,
	)
	if err != nil {
		log.Println("error while inserting drug ", err.Error())
		return "", err
	}

	return id.String(), nil

}

func (d *drugRepo) Get(ctx context.Context, request models.PrimaryKey) (models.Drug, error) {

	drug := models.Drug{}

	query := `select 
	 id,
	 drug_store_branch_id,
	 name,
	 description,
	 count,
	 price,
	 date_of_manufacture,
	 best_before,
	 created_at,
	 updated_at
	 from drug where deleted_at is null and id = $1`

	row := d.pool.QueryRow(ctx, query, request.ID)

	err := row.Scan(
		&drug.ID,
		&drug.DrugStoreBranchID,
		&drug.Name,
		&drug.Description,
		&drug.Count,
		&drug.Price,
		&drug.DateOfManufacture,
		&drug.BestBefore,
		&drug.CreatedAt,
		&drug.UpdatedAt,
	)

	if err != nil {
		log.Println("error while selecting drug ", err.Error())
		return models.Drug{}, err
	}

	return drug, nil

}

func (d *drugRepo) GetList(ctx context.Context, request models.GetListRequest) (models.DrugsResponse, error) {

	var (
		drugs             = []models.Drug{}
		count             = 0
		query, countQuery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)

	countQuery = `select count(1) from drug where deleted_at is null`

	if search != "" {
		countQuery += fmt.Sprintf(` and (name ilike '%%%s%%' or description ilike '%%%s%%')`, search, search)
	}
	if err := d.pool.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		fmt.Println("error is while selecting count", err.Error())
		return models.DrugsResponse{}, err
	}

	query = `select 
	 id,
	 drug_store_branch_id,
	 name,
	 description,
	 count,
	 price,
	 date_of_manufacture,
	 best_before,
	 created_at,
	 updated_at
	 from drug where deleted_at is null`

	if search != "" {
		query += fmt.Sprintf(` and (name ilike '%%%s%%' or description ilike '%%%s%%')`, search, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := d.pool.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting drug ", err.Error())
		return models.DrugsResponse{}, err
	}

	for rows.Next() {
		drug := models.Drug{}
		if err = rows.Scan(
			&drug.ID,
			&drug.DrugStoreBranchID,
			&drug.Name,
			&drug.Description,
			&drug.Count,
			&drug.Price,
			&drug.DateOfManufacture,
			&drug.BestBefore,
			&drug.CreatedAt,
			&drug.UpdatedAt,
		); err != nil {
			fmt.Println("error is while scanning drug data", err.Error())
			return models.DrugsResponse{}, err
		}

		drugs = append(drugs, drug)

	}

	return models.DrugsResponse{
		Drugs: drugs,
		Count: count,
	}, nil
}

func (d *drugRepo) Update(ctx context.Context, request models.UpdateDrug) (string, error) {

	query := `update drug set
	drug_store_branch_id = $1,
	name = $2,
	description = $3,
	count = $4,
	price = $5
    updated_at = $6, 
	 where id = $7  
   `

	_, err := d.pool.Exec(ctx, query,
		request.DrugStoreBranchID,
		request.Name,
		request.Description,
		request.Count,
		request.Count,
		time.Now(),
		request.ID)
	if err != nil {
		log.Println("error while updating drug data...", err.Error())
		return "", err
	}

	return request.ID, nil

}

func (d *drugRepo) Delete(ctx context.Context, id string) error {

	query := `
	update drug
	 set deleted_at = $1
	  where id = $2
	`

	_, err := d.pool.Exec(ctx, query, time.Now(), id)
	if err != nil {
		log.Println("error while deleting drug by id", err.Error())
		return err
	}

	return nil

}
