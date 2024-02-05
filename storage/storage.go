package storage

import (
	"context"
	"shifolink/api/models"
)

type IStorage interface {
	CloseDB()
	Author() IAuthorRepo
	ClinicAdmin() IClinicAdminRepo
	ClinicBranch() IClinicBranchRepo
	Clinic() IClinicRepo
	Customer() ICustomerRepo
	DoctorType() IDoctorTypeRepo
	Doctor() IDoctorRepo
	DrugStoreBranch() IDrugStoreBranchRepo
	DrugStore() IDrugStoreRepo
	Drug() IDrugRepo
	Journal() IJournalRepo
	OrderDrug() IOrderDrugRepo
	Orders() IOrdersRepo
	Pharmacist() IPharmacistRepo
	Queue() IQueueRepo
	SuperAdmin() ISuperAdminRepo
}

type IAuthorRepo interface {
	Create(context.Context, models.CreateAuthor) (string, error)
	Get(context.Context, models.PrimaryKey) (models.Author, error)
	GetList(context.Context, models.GetListRequest) (models.AuthorsResponse, error)
	Update(context.Context, models.UpdateAuthor) (string, error)
	Delete(context.Context, string) error
	UpdatePassword(context.Context, string) error
}

type IClinicAdminRepo interface {
	Create(context.Context, models.CreateClinicAdmin) (string, error)
	Get(context.Context, models.PrimaryKey) (models.ClinicAdmin, error)
	GetList(context.Context, models.GetListRequest) (models.ClinicAdminsResponse, error)
	Update(context.Context, models.UpdateClinicAdmin) (string, error)
	Delete(context.Context, string) error
	UpdatePassword(context.Context, string) error
}

type IClinicBranchRepo interface {
	Create(context.Context, models.CreateClinicBranch) (string, error)
	Get(context.Context, models.PrimaryKey) (models.ClinicBranch, error)
	GetList(context.Context, models.GetListRequest) (models.ClinicBranchsResponse, error)
	Update(context.Context, models.UpdateClinicBranch) (string, error)
	Delete(context.Context, string) error
}

type IClinicRepo interface {
	Create(context.Context, models.CreateClinic) (string, error)
	Get(context.Context, models.PrimaryKey) (models.Clinic, error)
	GetList(context.Context, models.GetListRequest) (models.ClinicsResponse, error)
	Update(context.Context, models.UpdateClinic) (string, error)
	Delete(context.Context, string) error
}

type ICustomerRepo interface {
	Create(context.Context, models.CreateCustomer) (string, error)
	Get(context.Context, models.PrimaryKey) (models.Customer, error)
	GetList(context.Context, models.GetListRequest) (models.CustomersResponse, error)
	Update(context.Context, models.UpdateCustomer) (string, error)
	Delete(context.Context, string) error
	UpdatePassword(context.Context, string) error
}

type IDoctorTypeRepo interface {
	Create(context.Context, models.CreateDoctorType) (string, error)
	Get(context.Context, models.PrimaryKey) (models.DoctorType, error)
	GetList(context.Context, models.GetListRequest) (models.DoctorTypesResponse, error)
	Update(context.Context, models.UpdateDoctorType) (string, error)
	Delete(context.Context, string) error
}

type IDoctorRepo interface {
	Create(context.Context, models.CreateDoctor) (string, error)
	Get(context.Context, models.PrimaryKey) (models.Doctor, error)
	GetList(context.Context, models.GetListRequest) (models.DoctorsResponse, error)
	Update(context.Context, models.UpdateDoctor) (string, error)
	Delete(context.Context, string) error
	UpdatePassword(context.Context, string) error
}

type IDrugStoreBranchRepo interface {
	Create(context.Context, models.CreateDrugStoreBranch) (string, error)
	Get(context.Context, models.PrimaryKey) (models.DrugStoreBranch, error)
	GetList(context.Context, models.GetListRequest) (models.DrugStoreBranchsResponse, error)
	Update(context.Context, models.UpdateDrugStoreBranch) (string, error)
	Delete(context.Context, string) error
}

type IDrugStoreRepo interface {
	Create(context.Context, models.CreateDrugStore) (string, error)
	Get(context.Context, models.PrimaryKey) (models.DrugStore, error)
	GetList(context.Context, models.GetListRequest) (models.DrugStoresResponse, error)
	Update(context.Context, models.UpdateDrugStore) (string, error)
	Delete(context.Context, string) error
}

type IDrugRepo interface {
	Create(context.Context, models.CreateDrug) (string, error)
	Get(context.Context, models.PrimaryKey) (models.Drug, error)
	GetList(context.Context, models.GetListRequest) (models.DrugsResponse, error)
	Update(context.Context, models.UpdateDrug) (string, error)
	Delete(context.Context, string) error
}

type IJournalRepo interface {
	Create(context.Context, models.CreateJournal) (string, error)
	Get(context.Context, models.PrimaryKey) (models.Journal, error)
	GetList(context.Context, models.GetListRequest) (models.JournalsResponse, error)
	Update(context.Context, models.UpdateJournal) (string, error)
	Delete(context.Context, string) error
}

type IOrderDrugRepo interface {
	Create(context.Context, models.CreateOrderDrug) (string, error)
	Get(context.Context, models.PrimaryKey) (models.OrderDrug, error)
	GetList(context.Context, models.GetListRequest) (models.OrderDrugsResponse, error)
	Update(context.Context, models.UpdateOrderDrug) (string, error)
	Delete(context.Context, string) error
}

type IOrdersRepo interface {
	Create(context.Context, models.CreateOrders) (string, error)
	Get(context.Context, models.PrimaryKey) (models.Orders, error)
	GetList(context.Context, models.GetListRequest) (models.OrdersResponse, error)
	Update(context.Context, models.UpdateOrders) (string, error)
	Delete(context.Context, string) error
}

type IPharmacistRepo interface {
	Create(context.Context, models.CreatePharmacist) (string, error)
	Get(context.Context, models.PrimaryKey) (models.Pharmacist, error)
	GetList(context.Context, models.GetListRequest) (models.PharmacistsResponse, error)
	Update(context.Context, models.UpdatePharmacist) (string, error)
	Delete(context.Context, string) error
	UpdatePassword(context.Context, string) error
}

type IQueueRepo interface {
	Create(context.Context, models.CreateQueue) (string, error)
	Get(context.Context, models.PrimaryKey) (models.Queue, error)
	GetList(context.Context, models.GetListRequest) (models.QueuesResponse, error)
	Update(context.Context, models.UpdateQueue) (string, error)
	Delete(context.Context, string) error
}

type ISuperAdminRepo interface {
	Create(context.Context, models.CreateSuperAdmin) (string, error)
	Get(context.Context, models.PrimaryKey) (models.SuperAdmin, error)
	GetList(context.Context, models.GetListRequest) (models.SuperAdminsResponse, error)
	Update(context.Context, models.UpdateSuperAdmin) (string, error)
	Delete(context.Context, string) error
	UpdatePassword(context.Context, string) error
}
