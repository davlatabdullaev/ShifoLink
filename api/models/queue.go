package models

import "time"

type Queue struct {
	ID          string    `json:"id"`
	CustomerID  string    `json:"customer_id"`
	DoctorID    string    `json:"doctor_id"`
	QueueNumber string    `json:"queue_number"`
	QueueTime   string    `json:"queue_time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type CreateQueue struct {
	CustomerID  string `json:"customer_id"`
	DoctorID    string `json:"doctor_id"`
	QueueNumber string `json:"queue_number"`
	QueueTime   string `json:"queue_time"`
}

type UpdateQueue struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	DoctorID   string `json:"doctor_id"`
	QueueTime  string `json:"queue_time"`
}

type QueuesResponse struct {
	Queues []Queue `json:"queues"`
	Count  int     `json:"count"`
}
