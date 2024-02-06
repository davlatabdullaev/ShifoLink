package models

import "time"

type Drug struct {
	ID                string    `json:"id"`
	DrugStoreBranchID string    `json:"drug_store_branch_id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Count             int       `json:"count"`
	Price             string    `json:"price"`
	DateOfManufacture string    `json:"date_of_manufacture"`
	BestBefore        string    `json:"best_before"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         time.Time `json:"deleted_at"`
}

type CreateDrug struct {
	DrugStoreBranchID string `json:"drug_store_branch_id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	Count             int    `json:"count"`
	Price             string `json:"price"`
	DateOfManufacture string `json:"date_of_manufacture"`
	BestBefore        string `json:"best_before"`
}

type UpdateDrug struct {
	ID                string `json:"id"`
	DrugStoreBranchID string `json:"drug_store_branch_id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	Count             int    `json:"count"`
	Price             string `json:"price"`
}

type DrugsResponse struct {
	Drugs []Drug `json:"drugs"`
	Count int    `json:"count"`
}
