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

type drugStoreRepo struct {
	pool *pgxpool.Pool
}

func NewDrugStoreRepo(pool *pgxpool.Pool) storage.IDrugStoreRepo {
	return &drugStoreRepo{
		pool: pool,
	}
}

func (d *drugStoreRepo) Create(ctx context.Context, request models.CreateDrugStore) (string, error) {

	id := uuid.New()

	query := `insert into drug_store
	 (id, 
	  name,
	  description) 
	  values ($1, $2, $3)`

	_, err := d.pool.Exec(ctx, query,
		id,
		request.Name,
		request.Description,
	)
	if err != nil {
		log.Println("error while inserting drug store ", err.Error())
		return "", err
	}

	return id.String(), nil

}

func (d *drugStoreRepo) Get(ctx context.Context, request models.PrimaryKey) (models.DrugStore, error) {

	var updatedAt = sql.NullTime{}

	drugStore := models.DrugStore{}

	query := `select 
	 id,
	 name,
	 description,
	 created_at,
	 updated_at
	 from drug_store where deleted_at is null and id = $1`

	row := d.pool.QueryRow(ctx, query, request.ID)

	err := row.Scan(
		&drugStore.ID,
		&drugStore.Name,
		&drugStore.Description,
		&drugStore.CreatedAt,
		&updatedAt,
	)

	if err != nil {
		log.Println("error while selecting drug store ", err.Error())
		return models.DrugStore{}, err
	}

	if updatedAt.Valid {
		drugStore.UpdatedAt = updatedAt.Time
	}

	return drugStore, nil

}

func (d *drugStoreRepo) GetList(ctx context.Context, request models.GetListRequest) (models.DrugStoresResponse, error) {

	var (
		updatedAt         = sql.NullTime{}
		drugStores        = []models.DrugStore{}
		count             = 0
		query, countQuery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)

	countQuery = `select count(1) from drug_store where deleted_at is null`

	if search != "" {
		countQuery += fmt.Sprintf(` and (name ilike '%%%s%%' or description ilike '%%%s%%')`, search, search)
	}
	if err := d.pool.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		fmt.Println("error is while selecting count", err.Error())
		return models.DrugStoresResponse{}, err
	}

	query = `select 
	 id,
	 name,
	 description,
	 created_at, 
	 updated_at from drug_store where deleted_at is null`

	if search != "" {
		query += fmt.Sprintf(` and (name ilike '%%%s%%' or description ilike '%%%s%%')`, search, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := d.pool.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting drug store ", err.Error())
		return models.DrugStoresResponse{}, err
	}

	for rows.Next() {
		drugStore := models.DrugStore{}
		if err = rows.Scan(
			&drugStore.ID,
			&drugStore.Name,
			&drugStore.Description,
			&drugStore.CreatedAt,
			&updatedAt,
		); err != nil {
			fmt.Println("error is while scanning drug store data", err.Error())
			return models.DrugStoresResponse{}, err
		}

		if updatedAt.Valid {
			drugStore.UpdatedAt = updatedAt.Time
		}

		drugStores = append(drugStores, drugStore)

	}

	return models.DrugStoresResponse{
		DrugStores: drugStores,
		Count:      count,
	}, nil
}

func (d *drugStoreRepo) Update(ctx context.Context, request models.UpdateDrugStore) (string, error) {

	query := `update drug_store set
	name = $1,
	description = $2,
    updated_at = $3 
	 where id = $4  
   `

	_, err := d.pool.Exec(ctx, query,
		request.Name,
		request.Description,
		time.Now(),
		request.ID)
	if err != nil {
		log.Println("error while updating drug store data...", err.Error())
		return "", err
	}

	return request.ID, nil

}

func (d *drugStoreRepo) Delete(ctx context.Context, id string) error {

	query := `
	update drug_store
	 set deleted_at = $1
	  where id = $2
	`

	_, err := d.pool.Exec(ctx, query, time.Now(), id)
	if err != nil {
		log.Println("error while deleting drug store  by id", err.Error())
		return err
	}

	return nil
}
