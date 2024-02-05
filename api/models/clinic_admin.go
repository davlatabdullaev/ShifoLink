package models

import "time"

type ClinicAdmin struct {
	ID             string    `json:"id"`
	ClinicBranchID string    `json:"clinic_branch_id"`
	DoctorTypeID   string    `json:"doctor_type_id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	Phone          string    `json:"phone"`
	Gender         string    `json:"gender"`
	BirthDate      string    `json:"birth_date"`
	Age            int       `json:"age"`
	Address        string    `json:"address"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}

type CreateClinicAdmin struct {
	ClinicBranchID string `json:"clinic_branch_id"`
	DoctorTypeID   string `json:"doctor_type_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Phone          string `json:"phone"`
	Gender         string `json:"gender"`
	BirthDate      string `json:"birth_date"`
	Address        string `json:"address"`
}

type UpdateClinicAdmin struct {
	ClinicBranchID string `json:"clinic_branch_id"`
	DoctorTypeID   string `json:"doctor_type_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Address        string `json:"address"`
}

type UpdateClinicAdminPassword struct {
	Password string `json:"password"`
}

type ClinicAdminsResponse struct {
	ClinicAdmins []ClinicAdmin `json:"clinic_admins"`
	Count        int           `json:"count"`
}
