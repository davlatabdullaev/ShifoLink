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

type queueRepo struct {
	pool *pgxpool.Pool
}

func NewQueueRepo(pool *pgxpool.Pool) storage.IQueueRepo {
	return &queueRepo{
		pool: pool,
	}
}

func (q *queueRepo) Create(ctx context.Context, request models.CreateQueue) (string, error) {

	id := uuid.New()

	query := `insert into queue
	 (id, 
	  customer_id,
	  doctor_id,
	  queue_time) 
	  values ($1, $2, $3, $4)`

	rowsAffected, err := q.pool.Exec(ctx, query,
		id,
		request.CustomerID,
		request.DoctorID,
		request.QueueTime,
	)

	if r := rowsAffected.RowsAffected(); r == 0 {
		log.Println("error is while rows affected ", err.Error())
		return "", err
	}

	if err != nil {
		log.Println("error while inserting queue ", err.Error())
		return "", err
	}

	return id.String(), nil

}

func (q *queueRepo) Get(ctx context.Context, request models.PrimaryKey) (models.Queue, error) {

	var updatedAt = sql.NullTime{}

	queue := models.Queue{}

	query := `select 
	 id,
	 customer_id,
	 doctor_id,
	 queue_number,
	 queue_time,
	 created_at,
	 updated_at
	 from queue where deleted_at is null and id = $1`

	row := q.pool.QueryRow(ctx, query, request.ID)

	err := row.Scan(
		&queue.ID,
		&queue.CustomerID,
		&queue.DoctorID,
		&queue.QueueNumber,
		&queue.QueueTime,
		&queue.CreatedAt,
		&updatedAt,
	)

	if err != nil {
		log.Println("error while selecting queue", err.Error())
		return models.Queue{}, err
	}

	if updatedAt.Valid {
		queue.UpdatedAt = updatedAt.Time
	}

	return queue, nil

}

func (q *queueRepo) GetList(ctx context.Context, request models.GetListRequest) (models.QueuesResponse, error) {

	var (
		updatedAt         = sql.NullTime{}
		queues            = []models.Queue{}
		count             = 0
		query, countQuery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)

	countQuery = `select count(1) from queue where deleted_at is null`

	if search != "" {
		countQuery += fmt.Sprintf(` and (queue_number ilike '%%%s%%')`, search)
	}
	if err := q.pool.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		fmt.Println("error is while selecting count", err.Error())
		return models.QueuesResponse{}, err
	}

	query = `select 
	id,
	customer_id,
	doctor_id,
	queue_number,
	queue_time,
	created_at,
	updated_at from queue where deleted_at is null`

	if search != "" {
		query += fmt.Sprintf(`  and (queue_number ilike '%%%s%%')`, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := q.pool.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting queue ", err.Error())
		return models.QueuesResponse{}, err
	}

	for rows.Next() {
		queue := models.Queue{}
		if err = rows.Scan(
			&queue.ID,
			&queue.CustomerID,
			&queue.DoctorID,
			&queue.QueueNumber,
			&queue.QueueTime,
			&queue.CreatedAt,
			&updatedAt); err != nil {
			fmt.Println("error is while scanning queues data", err.Error())
			return models.QueuesResponse{}, err
		}

		if updatedAt.Valid {
			queue.UpdatedAt = updatedAt.Time
		}

		queues = append(queues, queue)

	}

	return models.QueuesResponse{
		Queues: queues,
		Count:  count,
	}, nil
}

func (q *queueRepo) Update(ctx context.Context, request models.UpdateQueue) (string, error) {

	query := `update queue set
	customer_id = $1,
	doctor_id = $2,
	queue_time = $3,
    updated_at = $4
	 where id = $5  
   `

	rowsAffected, err := q.pool.Exec(ctx, query,
		request.CustomerID,
		request.DoctorID,
		request.QueueTime,
		time.Now(),
		request.ID)

	if r := rowsAffected.RowsAffected(); r == 0 {
		log.Println("error is while rows affected ", err.Error())
		return "", err
	}

	if err != nil {
		log.Println("error while updating queue data...", err.Error())
		return "", err
	}

	return request.ID, nil

}

func (q *queueRepo) Delete(ctx context.Context, id string) error {

	query := `
	update queue
	 set deleted_at = $1
	  where id = $2
	`

	rowsAffected, err := q.pool.Exec(ctx, query, time.Now(), id)

	if r := rowsAffected.RowsAffected(); r == 0 {
		log.Println("error is while rows affected ", err.Error())
		return err
	}

	if err != nil {
		log.Println("error while deleting queue  by id", err.Error())
		return err
	}

	return nil

}
