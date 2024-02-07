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

type journalRepo struct {
	pool *pgxpool.Pool
}

func NewJournalRepo(pool *pgxpool.Pool) storage.IJournalRepo {
	return &journalRepo{
		pool: pool,
	}
}

func (j *journalRepo) Create(ctx context.Context, request models.CreateJournal) (string, error) {

	id := uuid.New()

	query := `insert into journal
	 (id, 
	  author_id,
	  theme,
	  article) 
	  values ($1, $2, $3, $4)`

	rowsAffected, err := j.pool.Exec(ctx, query,
		id,
		request.AuthorID,
		request.Theme,
		request.Article,
	)

	if r := rowsAffected.RowsAffected(); r == 0 {
		log.Println("error is while rows affected ", err.Error())
		return "", err
	}

	if err != nil {
		log.Println("error while inserting journal ", err.Error())
		return "", err
	}

	return id.String(), nil

}

func (j *journalRepo) Get(ctx context.Context, request models.PrimaryKey) (models.Journal, error) {

	var updatedAt = sql.NullTime{}

	journal := models.Journal{}

	query := `select 
	 id,
	 author_id,
	 theme,
	 article,
	 created_at,
	 updated_at
	 from journal where deleted_at is null and id = $1`

	row := j.pool.QueryRow(ctx, query, request.ID)

	err := row.Scan(
		&journal.ID,
		&journal.AuthorID,
		&journal.Theme,
		&journal.Article,
		&journal.CreatedAt,
		&updatedAt,
	)

	if err != nil {
		log.Println("error while selecting journal", err.Error())
		return models.Journal{}, err
	}

	if updatedAt.Valid {
		journal.UpdatedAt = updatedAt.Time
	}

	return journal, nil
}

func (j *journalRepo) GetList(ctx context.Context, request models.GetListRequest) (models.JournalsResponse, error) {

	var (
		updatedAt         = sql.NullTime{}
		journals          = []models.Journal{}
		count             = 0
		query, countQuery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)

	countQuery = `select count(1) from journal where deleted_at is null`

	if search != "" {
		countQuery += fmt.Sprintf(` and (theme ilike '%%%s%%' or article ilike '%%%s%%')`, search, search)
	}
	if err := j.pool.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		fmt.Println("error is while selecting count", err.Error())
		return models.JournalsResponse{}, err
	}

	query = `select 
	 id,
	 author_id,
	 theme,
	 article,
	 created_at, 
	 updated_at from journal where deleted_at is null`

	if search != "" {
		query += fmt.Sprintf(` and (theme ilike '%%%s%%' or article ilike '%%%s%%')`, search, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := j.pool.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting journal ", err.Error())
		return models.JournalsResponse{}, err
	}

	for rows.Next() {
		journal := models.Journal{}
		if err = rows.Scan(
			&journal.ID,
			&journal.AuthorID,
			&journal.Theme,
			&journal.Article,
			&journal.CreatedAt,
			&updatedAt,
		); err != nil {
			fmt.Println("error is while scanning journal data", err.Error())
			return models.JournalsResponse{}, err
		}

		if updatedAt.Valid {
			journal.UpdatedAt = updatedAt.Time
		}

		journals = append(journals, journal)

	}

	return models.JournalsResponse{
		Journals: journals,
		Count:    count,
	}, nil
}

func (j *journalRepo) Update(ctx context.Context, request models.UpdateJournal) (string, error) {

	query := `update journal set
	author_id = $1,
	theme = $2,
	article = $3,
    updated_at = $4 
	 where id = $5  
   `

	rowsAffected, err := j.pool.Exec(ctx, query,
		request.AuthorID,
		request.Theme,
		request.Article,
		time.Now(),
		request.ID)

	if r := rowsAffected.RowsAffected(); r == 0 {
		log.Println("error is while rows affected ", err.Error())
		return "", err
	}

	if err != nil {
		log.Println("error while updating journal data...", err.Error())
		return "", err
	}

	return request.ID, nil

}

func (j *journalRepo) Delete(ctx context.Context, id string) error {

	query := `
	update journal
	 set deleted_at = $1
	  where id = $2
	`

	rowsAffected, err := j.pool.Exec(ctx, query, time.Now(), id)

	if r := rowsAffected.RowsAffected(); r == 0 {
		log.Println("error is while rows affected ", err.Error())
		return err
	}

	if err != nil {
		log.Println("error while deleting journal  by id", err.Error())
		return err
	}

	return nil
}
