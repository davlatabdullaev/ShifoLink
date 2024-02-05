package models

import "time"

type Clinic struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type CreateClinic struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateClinic struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ClinicsResponse struct {
	Clinics []Clinic `json:"clinics"`
	Count   int      `json:"count"`
}
