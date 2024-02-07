package models

import "time"

type Author struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	Gender    string    `json:"gender"`
	BirthDate string    `json:"birth_date"`
	Age       int       `json:"age"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type CreateAuthor struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birth_date"`
	Address   string `json:"address"`
}

type UpdateAuthor struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}

type UpdateAuthorPassword struct {
	ID          string `json:"-"`
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}

type AuthorsResponse struct {
	Authors []Author `json:"authors"`
	Count   int      `json:"count"`
}
