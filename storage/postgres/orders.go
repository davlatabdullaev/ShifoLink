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

type ordersRepo struct {
	pool *pgxpool.Pool
}

func NewOrdersRepo(pool *pgxpool.Pool) storage.IOrdersRepo {
	return &ordersRepo{
		pool: pool,
	}
}

func (o *ordersRepo) Create(ctx context.Context, request models.CreateOrders) (string, error) {

	id := uuid.New()

	query := `insert into orders
	 (id, 
	  pharmacist_id,
	  customer_id) 
	  values ($1, $2, $3)`

	rowsAffected, err := o.pool.Exec(ctx, query,
		id,
		request.PharmacistID,
		request.CustomerID,
	)

	if r := rowsAffected.RowsAffected(); r == 0 {
		log.Println("error is while rows affected ", err.Error())
		return "", err
	}

	if err != nil {
		log.Println("error while inserting orders ", err.Error())
		return "", err
	}

	return id.String(), nil

}

func (o *ordersRepo) Get(ctx context.Context, request models.PrimaryKey) (models.Orders, error) {

	var updatedAt = sql.NullTime{}

	orders := models.Orders{}

	query := `select 
	 id,
	 pharmacist_id,
	 customer_id,
	 created_at,
	 updated_at
	 from orders where deleted_at is null and id = $1`

	row := o.pool.QueryRow(ctx, query, request.ID)

	err := row.Scan(
		&orders.ID,
		&orders.PharmacistID,
		&orders.CustomerID,
		&orders.CreatedAt,
		&updatedAt,
	)

	if err != nil {
		log.Println("error while selecting orders", err.Error())
		return models.Orders{}, err
	}

	if updatedAt.Valid {
		orders.UpdatedAt = updatedAt.Time
	}

	return orders, nil

}

func (o *ordersRepo) GetList(ctx context.Context, request models.GetListRequest) (models.OrdersResponse, error) {

	var (
		updatedAt         = sql.NullTime{}
		orders            = []models.Orders{}
		count             = 0
		query, countQuery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)

	countQuery = `select count(1) from orders where deleted_at is null`

	if search != "" {
		countQuery += fmt.Sprintf(` and (pharmacist_id ilike '%%%s%%' or customer_id ilike '%%%s%%')`, search, search)
	}
	if err := o.pool.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		fmt.Println("error is while selecting count", err.Error())
		return models.OrdersResponse{}, err
	}

	query = `select 
	 id,
	 pharmacist_id,
	 customer_id,
	 created_at, 
	 updated_at from orders where deleted_at is null`

	if search != "" {
		query += fmt.Sprintf(` and (pharmacist_id ilike '%%%s%%' or customer_id ilike '%%%s%%')`, search, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := o.pool.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting orders ", err.Error())
		return models.OrdersResponse{}, err
	}

	for rows.Next() {
		order := models.Orders{}
		if err = rows.Scan(
			&order.ID,
			&order.PharmacistID,
			&order.CustomerID,
			&order.CreatedAt,
			&updatedAt,
		); err != nil {
			fmt.Println("error is while scanning orders data", err.Error())
			return models.OrdersResponse{}, err
		}

		if updatedAt.Valid {
			order.UpdatedAt = updatedAt.Time
		}

		orders = append(orders, order)

	}

	return models.OrdersResponse{
		Orderss: orders,
		Count:   count,
	}, nil
}

func (o *ordersRepo) Update(ctx context.Context, request models.UpdateOrders) (string, error) {

	query := `update orders set
	pharmacist_id = $1,
	customer_id = $2,
    updated_at = $3 
	 where id = $4  
   `

	rowsAffected, err := o.pool.Exec(ctx, query,
		request.PharmacistID,
		request.CustomerID,
		time.Now(),
		request.ID)

	if r := rowsAffected.RowsAffected(); r == 0 {
		log.Println("error is while rows affected ", err.Error())
		return "", err
	}

	if err != nil {
		log.Println("error while updating orders data...", err.Error())
		return "", err
	}

	return request.ID, nil

}

func (o *ordersRepo) Delete(ctx context.Context, id string) error {

	query := `
	update orders
	 set deleted_at = $1
	  where id = $2
	`

	rowsAffected, err := o.pool.Exec(ctx, query, time.Now(), id)

	if r := rowsAffected.RowsAffected(); r == 0 {
		log.Println("error is while rows affected ", err.Error())
		return err
	}

	if err != nil {
		log.Println("error while deleting orders  by id", err.Error())
		return err
	}

	return nil
}
