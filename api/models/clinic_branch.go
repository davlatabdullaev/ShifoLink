package models

import "time"

type ClinicBranch struct {
	ID          string    `json:"id"`
	ClinicID    string    `json:"clinic_id"`
	Address     string    `json:"address"`
	Phone       string    `json:"phone"`
	WorkingTime string    `json:"working_time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type CreateClinicBranch struct {
	ClinicID    string `json:"clinic_id"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	WorkingTime string `json:"working_time"`
}

type UpdateClinicBranch struct {
	ID          string `json:"id"`
	ClinicID    string `json:"clinic_id"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	WorkingTime string `json:"working_time"`
}

type ClinicBranchsResponse struct {
	ClinicBranchs []ClinicBranch `json:"clinic_branch"`
	Count         int            `json:"count"`
}
