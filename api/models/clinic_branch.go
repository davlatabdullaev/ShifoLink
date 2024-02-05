package models

import "time"

type ClinicBranch struct {
	ID        string    `json:"id"`
	ClinicID  string    `json:"clinic_id"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type CreateClinicBranch struct {
	ClinicID string `json:"clinic_id"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type UpdateClinicBranch struct {
	ClinicID string `json:"clinic_id"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type ClinicBranchsResponse struct {
	ClinicBranchs []ClinicBranch `json:"clinic_branch"`
	Count         int            `json:"count"`
}
