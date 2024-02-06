package models

import "time"

type Journal struct {
	ID        string    `json:"id"`
	AuthorID  string    `json:"author_id"`
	Theme     string    `json:"theme"`
	Article   string    `json:"article"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type CreateJournal struct {
	AuthorID string `json:"author_id"`
	Theme    string `json:"theme"`
	Article  string `json:"article"`
}

type UpdateJournal struct {
	ID       string `json:"id"`
	AuthorID string `json:"author_id"`
	Theme    string `json:"theme"`
	Article  string `json:"article"`
}

type JournalsResponse struct {
	Journals []Journal `json:"journals"`
	Count    int       `json:"count"`
}
