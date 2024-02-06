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

type drugStoreBranchRepo struct {
	pool *pgxpool.Pool
}

func NewDrugStoreBranchRepo(pool *pgxpool.Pool) storage.IDrugStoreBranchRepo {
	return &drugStoreBranchRepo{
		pool: pool,
	}
}

func (d *drugStoreBranchRepo) Create(ctx context.Context, request models.CreateDrugStoreBranch) (string, error) {

	id := uuid.New()

	query := `insert into drug_store_branch
	 (id, 
	  drug_store_id,
	  address,
	  phone,
	  working_time) 
	  values ($1, $2, $3, $4, $5)`

	_, err := d.pool.Exec(ctx, query,
		id,
		request.DrugStoreID,
		request.Address,
		request.Phone,
		request.WorkingTime,
	)
	if err != nil {
		log.Println("error while inserting drug store branch", err.Error())
		return "", err
	}

	return id.String(), nil

}

func (d *drugStoreBranchRepo) Get(ctx context.Context, request models.PrimaryKey) (models.DrugStoreBranch, error) {

	drugStoreBranch := models.DrugStoreBranch{}

	query := `select 
	 id,
	 drug_store_id,
	 address,
	 phone,
	 working_time,
	 created_at,
	 updated_at
	 from drug_store_branch where deleted_at is null and id = $1`

	row := d.pool.QueryRow(ctx, query, request.ID)

	err := row.Scan(
		&drugStoreBranch.ID,
		&drugStoreBranch.DrugStoreID,
		&drugStoreBranch.Address,
		&drugStoreBranch.Phone,
		&drugStoreBranch.WorkingTime,
		&drugStoreBranch.CreatedAt,
		&drugStoreBranch.UpdatedAt,
	)

	if err != nil {
		log.Println("error while selecting drug store branch ", err.Error())
		return models.DrugStoreBranch{}, err
	}

	return drugStoreBranch, nil

}

func (d *drugStoreBranchRepo) GetList(ctx context.Context, request models.GetListRequest) (models.DrugStoreBranchsResponse, error) {

	var (
		drugStoreBranchs  = []models.DrugStoreBranch{}
		count             = 0
		query, countQuery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)

	countQuery = `select count(1) from drug_store_branch where deleted_at is null`

	if search != "" {
		countQuery += fmt.Sprintf(` and (address ilike '%%%s%%' or phone ilike '%%%s%%')`, search, search)
	}
	if err := d.pool.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		fmt.Println("error is while selecting count", err.Error())
		return models.DrugStoreBranchsResponse{}, err
	}

	query = `select 
	 id,
	 drug_store_id,
	 address,
	 phone,
	 working_time,
	 created_at, 
	 updated_at from drug_store_branch where deleted_at is null`

	if search != "" {
		query += fmt.Sprintf(` and (address ilike '%%%s%%' or phone ilike '%%%s%%')`, search, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := d.pool.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting drug store branch ", err.Error())
		return models.DrugStoreBranchsResponse{}, err
	}

	for rows.Next() {
		drugStoreBranch := models.DrugStoreBranch{}
		if err = rows.Scan(
			&drugStoreBranch.ID,
			&drugStoreBranch.DrugStoreID,
			&drugStoreBranch.Address,
			&drugStoreBranch.Phone,
			&drugStoreBranch.WorkingTime,
			&drugStoreBranch.CreatedAt,
			&drugStoreBranch.UpdatedAt,
		); err != nil {
			fmt.Println("error is while scanning drug store branch data", err.Error())
			return models.DrugStoreBranchsResponse{}, err
		}

		drugStoreBranchs = append(drugStoreBranchs, drugStoreBranch)

	}

	return models.DrugStoreBranchsResponse{
		DrugStoreBranchs: drugStoreBranchs,
		Count:            count,
	}, nil

}

func (d *drugStoreBranchRepo) Update(ctx context.Context, request models.UpdateDrugStoreBranch) (string, error) {

	query := `update drug_store_branch set
	drug_store_id = $1,
	address = $2,
	phone = $3,
	working_time = $4,
    updated_at = $5 
	 where id = $6  
   `

	_, err := d.pool.Exec(ctx, query,
		request.DrugStoreID,
		request.Address,
		request.Phone,
		request.WorkingTime,
		time.Now(),
		request.ID)
	if err != nil {
		log.Println("error while updating drug store branch data...", err.Error())
		return "", err
	}

	return request.ID, nil
}

func (d *drugStoreBranchRepo) Delete(ctx context.Context, id string) error {

	query := `
	update drug_store_branch
	 set deleted_at = $1
	  where id = $2
	`

	_, err := d.pool.Exec(ctx, query, time.Now(), id)
	if err != nil {
		log.Println("error while deleting drug store branch by id", err.Error())
		return err
	}

	return nil
}
