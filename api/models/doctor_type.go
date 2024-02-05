package models

import "time"

type DoctorType struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	ClinicBranchID string    `json:"clinic_branch_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}

type CreateDoctorType struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	ClinicBranchID string `json:"clinic_branch_id"`
}

type UpdateDoctorType struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	ClinicBranchID string `json:"clinic_branch_id"`
}

type DoctorTypesResponse struct {
	DoctorTypes []DoctorType `json:"doctor_types"`
	Count       int          `json:"count"`
}
