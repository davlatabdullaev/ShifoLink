package models

import "time"

type SuperAdmin struct {
	ID          string    `json:"id"`
	ClinicID    string    `json:"clinic_id"`
	DrugStoreID string    `json:"drug_store_id"`
	AuthorID    string    `json:"author_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Phone       string    `json:"phone"`
	Gender      string    `json:"gender"`
	BirthDate   string    `json:"birth_date"`
	Age         int       `json:"age"`
	Address     string    `json:"address"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type CreateSuperAdmin struct {
	ClinicID    string `json:"clinic_id"`
	DrugStoreID string `json:"drug_store_id"`
	AuthorID    string `json:"author_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Phone       string `json:"phone"`
	Gender      string `json:"gender"`
	BirthDate   string `json:"birth_date"`
	Address     string `json:"address"`
}

type UpdateSuperAdmin struct {
	ID          string `json:"id"`
	ClinicID    string `json:"clinic_id"`
	DrugStoreID string `json:"drug_store_id"`
	AuthorID    string `json:"author_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
}

type UpdateSuperAdminPassword struct {
	Password string `json:"password"`
}

type SuperAdminsResponse struct {
	SuperAdmins []SuperAdmin `json:"super_admins"`
	Count       int          `json:"count"`
}
