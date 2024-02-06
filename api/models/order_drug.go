package models

import "time"

type OrderDrug struct {
	ID        string    `json:"id"`
	DrugID    string    `json:"drug_id"`
	OrdersID  string    `json:"orders_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type CreateOrderDrug struct {
	DrugID   string `json:"drug_id"`
	OrdersID string `json:"orders_id"`
}

type UpdateOrderDrug struct {
	ID       string `json:"id"`
	DrugID   string `json:"drug_id"`
	OrdersID string `json:"orders_id"`
}

type OrderDrugsResponse struct {
	OrderDrugs []OrderDrug `json:"order_drugs"`
	Count      int         `json:"count"`
}
