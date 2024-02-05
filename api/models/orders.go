package models

import "time"

type Orders struct {
	ID           string    `json:"id"`
	PharmacistID string    `json:"pharmacist_id"`
	CustomerID   string    `json:"customer_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

type CreateOrders struct {
	PharmacistID string `json:"pharmacist_id"`
	CustomerID   string `json:"customer_id"`
}

type UpdateOrders struct {
	PharmacistID string `json:"pharmacist_id"`
	CustomerID   string `json:"customer_id"`
}

type OrdersResponse struct {
	Orderss []Orders `json:"orderss"`
	Count  int     `json:"count"`
}
