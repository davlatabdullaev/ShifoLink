package models

import "time"

type DrugStoreBranch struct {
	ID          string    `json:"id"`
	DrugStoreID string    `json:"drug_store_id"`
	Address     string    `json:"address"`
	Phone       string    `json:"phone"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type CreateDrugStoreBranch struct {
	DrugStoreID string `json:"drug_store_id"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
}

type UpdateDrugStoreBranch struct {
	DrugStoreID string `json:"drug_store_id"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
}

type DrugStoreBranchsResponse struct {
	DrugStoreBranchs []DrugStoreBranch `json:"drug_store_branchs"`
	Count            int               `json:"count"`
}
