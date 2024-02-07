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

type orderDrugRepo struct {
	pool *pgxpool.Pool
}

func NewOrderDrugRepo(pool *pgxpool.Pool) storage.IOrderDrugRepo {
	return &orderDrugRepo{
		pool: pool,
	}
}

func (o *orderDrugRepo) Create(ctx context.Context, request models.CreateOrderDrug) (string, error) {

	id := uuid.New()

	query := `insert into order_drug
	 (id, 
	  drug_id,
	  orders_id) 
	  values ($1, $2, $3)`

	_, err := o.pool.Exec(ctx, query,
		id,
		request.DrugID,
		request.OrdersID,
	)
	if err != nil {
		log.Println("error while inserting order drug ", err.Error())
		return "", err
	}

	return id.String(), nil

}

func (o *orderDrugRepo) Get(ctx context.Context, request models.PrimaryKey) (models.OrderDrug, error) {

	var updatedAt = sql.NullTime{}

	orderDrug := models.OrderDrug{}

	query := `select 
	 id,
	 drug_id,
	 orders_id,
	 created_at,
	 updated_at
	 from order_drug where deleted_at is null and id = $1`

	row := o.pool.QueryRow(ctx, query, request.ID)

	err := row.Scan(
		&orderDrug.ID,
		&orderDrug.DrugID,
		&orderDrug.OrdersID,
		&orderDrug.CreatedAt,
		&updatedAt,
	)

	if err != nil {
		log.Println("error while selecting order drug", err.Error())
		return models.OrderDrug{}, err
	}

	if updatedAt.Valid {
		orderDrug.UpdatedAt = updatedAt.Time
	}

	return orderDrug, nil

}

func (o *orderDrugRepo) GetList(ctx context.Context, request models.GetListRequest) (models.OrderDrugsResponse, error) {

	var (
		updatedAt         = sql.NullTime{}
		orderDrugs        = []models.OrderDrug{}
		count             = 0
		query, countQuery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)

	countQuery = `select count(1) from order_drug where deleted_at is null`

	if search != "" {
		countQuery += fmt.Sprintf(` and (drug_id ilike '%%%s%%' or orders_id ilike '%%%s%%')`, search, search)
	}
	if err := o.pool.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		fmt.Println("error is while selecting count", err.Error())
		return models.OrderDrugsResponse{}, err
	}

	query = `select 
	 id,
	 drug_id,
	 orders_id,
	 created_at, 
	 updated_at from order_drug where deleted_at is null`

	if search != "" {
		query += fmt.Sprintf(` and (drug_id ilike '%%%s%%' or orders_id ilike '%%%s%%')`, search, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := o.pool.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting order drug ", err.Error())
		return models.OrderDrugsResponse{}, err
	}

	for rows.Next() {
		orderDrug := models.OrderDrug{}
		if err = rows.Scan(
			&orderDrug.ID,
			&orderDrug.DrugID,
			&orderDrug.OrdersID,
			&orderDrug.CreatedAt,
			&updatedAt,
		); err != nil {
			fmt.Println("error is while scanning order drug data", err.Error())
			return models.OrderDrugsResponse{}, err
		}

		if updatedAt.Valid {
			orderDrug.UpdatedAt = updatedAt.Time
		}

		orderDrugs = append(orderDrugs, orderDrug)

	}

	return models.OrderDrugsResponse{
		OrderDrugs: orderDrugs,
		Count:      count,
	}, nil
}

func (o *orderDrugRepo) Update(ctx context.Context, request models.UpdateOrderDrug) (string, error) {

	query := `update order_drug set
	drug_id = $1,
	orders_id = $2,
    updated_at = $3 
	 where id = $4  
   `

	_, err := o.pool.Exec(ctx, query,
		request.DrugID,
		request.OrdersID,
		time.Now(),
		request.ID)
	if err != nil {
		log.Println("error while updating order drug data...", err.Error())
		return "", err
	}

	return request.ID, nil
}

func (o *orderDrugRepo) Delete(ctx context.Context, id string) error {

	query := `
	update order_drug
	 set deleted_at = $1
	  where id = $2
	`

	_, err := o.pool.Exec(ctx, query, time.Now(), id)
	if err != nil {
		log.Println("error while deleting order_drug  by id", err.Error())
		return err
	}

	return nil

}
