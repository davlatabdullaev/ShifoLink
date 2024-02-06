package models

import "time"

type Doctor struct {
	ID           string    `json:"id"`
	DoctorTypeID string    `json:"doctor_type_id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Phone        string    `json:"phone"`
	Gender       string    `json:"gender"`
	BirthDate    string    `json:"birth_date"`
	Age          int       `json:"age"`
	Address      string    `json:"address"`
	WorkingTime  string    `json:"working_time"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

type CreateDoctor struct {
	DoctorTypeID string `json:"doctor_type_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Phone        string `json:"phone"`
	Gender       string `json:"gender"`
	BirthDate    string `json:"birth_date"`
	Address      string `json:"address"`
	WorkingTime  string `json:"working_time"`
	Status       string `json:"status"`
}

type UpdateDoctor struct {
	ID           string `json:"id"`
	DoctorTypeID string `json:"doctor_type_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	WorkingTime  string `json:"working_time"`
	Status       string `json:"status"`
}

type UpdateDoctorPassword struct {
	Password string `json:"password"`
}

type DoctorsResponse struct {
	Doctors []Doctor `json:"doctors"`
	Count   int      `json:"count"`
}
