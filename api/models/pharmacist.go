package models

import "time"

type Pharmacist struct {
	ID                string    `json:"id"`
	DrugStoreBranchID string    `json:"drug_store_branch_id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Email             string    `json:"email"`
	Password          string    `json:"password"`
	Phone             string    `json:"phone"`
	Gender            string    `json:"gender"`
	BirthDate         string    `json:"birth_date"`
	Age               int       `json:"age"`
	Address           string    `json:"address"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         time.Time `json:"deleted_at"`
}

type CreatePharmacist struct {
	DrugStoreBranchID string `json:"drug_store_branch_id"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	Phone             string `json:"phone"`
	Gender            string `json:"gender"`
	BirthDate         string `json:"birth_date"`
	Address           string `json:"address"`
}

type UpdatePharmacist struct {
	ID                string `json:"id"`
	DrugStoreBranchID string `json:"drug_store_branch_id"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	Address           string `json:"address"`
}

type UpdatePharmacistPassword struct {
	Password string `json:"password"`
}

type PharmacistsResponse struct {
	Pharmacists []Pharmacist `json:"pharmacists"`
	Count       int          `json:"count"`
}
