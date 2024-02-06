package models

import "time"

type DrugStore struct {
	ID          string    `json:"drug_store"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type CreateDrugStore struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateDrugStore struct {
	ID          string `json:"drug_store"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DrugStoresResponse struct {
	DrugStores []DrugStore `json:"drug_stores"`
	Count      int         `json:"count"`
}
