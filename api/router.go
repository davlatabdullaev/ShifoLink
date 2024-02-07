package api

import (
	"shifolink/storage"

	_ "shifolink/api/docs"
	"shifolink/api/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title           ShifoLink
// @version         1.0.0
// @description     Online doctor appointments and drug orders
func New(store storage.IStorage) *gin.Engine {

	h := handler.New(store)

	r := gin.New()

	// 	AUTHOR

	r.POST("author", h.CreateAuthor)
	r.GET("author/:id", h.GetAuthorByID)
	r.GET("author", h.GetAuthorList)
	r.PUT("author/:id", h.UpdateAuthor)
	r.DELETE("author/:id", h.DeleteAuthor)

	// CLINIC ADMIN

	r.POST("clinic_admin", h.CreateClinicAdmin)
	r.GET("clinic_admin/:id", h.GetClinicAdminByID)
	r.GET("clinic_admin", h.GetClinicAdminsList)
	r.PUT("clinic_admin/:id", h.UpdateClinicAdmin)
	r.DELETE("clinic_admin/:id", h.DeleteClinicAdmin)

	// CLINIC BRANCH

	r.POST("clinic_branch", h.CreateClinicBranch)
	r.GET("clinic_branch/:id", h.GetClinicBranchByID)
	r.GET("clinic_branch", h.GetClinicBranchsList)
	r.PUT("clinic_branch/:id", h.UpdateClinicBranch)
	r.DELETE("clinic_branch/:id", h.DeleteClinicBranch)

	// CLINIC

	r.POST("clinic", h.CreateClinic)
	r.GET("clinic/:id", h.GetClinicByID)
	r.GET("clinic", h.GetClinicsList)
	r.PUT("clinic/:id", h.UpdateClinic)
	r.DELETE("clinic/:id", h.DeleteClinic)

	// CUSTOMER

	r.POST("customer", h.CreateCustomer)
	r.GET("customer/:id", h.GetCustomerByID)
	r.GET("customer", h.GetCustomersList)
	r.PUT("customer/:id", h.UpdateCustomer)
	r.DELETE("customer/:id", h.DeleteCustomer)

	// DOCTOR TYPE

	r.POST("doctor_type", h.CreateDoctorType)
	r.GET("doctor_type/:id", h.GetDoctorTypeByID)
	r.GET("doctor_type", h.GetDoctorTypesList)
	r.PUT("doctor_type/:id", h.UpdateDoctorType)
	r.DELETE("doctor_type/:id", h.DeleteDoctorType)

	// DOCTOR

	r.POST("doctor", h.CreateDoctor)
	r.GET("doctor/:id", h.GetDoctorByID)
	r.GET("doctor", h.GetDoctorsList)
	r.PUT("doctor/:id", h.UpdateDoctor)
	r.DELETE("doctor/:id", h.DeleteDoctor)

	// DRUG STORE BRANCH

	r.POST("drug_store_branch", h.CreateDrugStoreBranch)
	r.GET("drug_store_branch/:id", h.GetDrugStoreBranchByID)
	r.GET("drug_store_branch", h.GetDrugStoreBranchsList)
	r.PUT("drug_store_branch/:id", h.UpdateDrugStoreBranch)
	r.DELETE("drug_store_branch/:id", h.DeleteDrugStoreBranch)

	// DRUG STORE

	r.POST("drug_store", h.CreateDrugStore)
	r.GET("drug_store/:id", h.GetDrugStoreByID)
	r.GET("drug_store", h.GetDrugStoresList)
	r.PUT("drug_store/:id", h.UpdateDrugStore)
	r.DELETE("drug_store/:id", h.DeleteDrugStore)

	// DRUG

	r.POST("drug", h.CreateDrug)
	r.GET("drug/:id", h.GetDrugByID)
	r.GET("drug", h.GetDrugsList)
	r.PUT("drug/:id", h.UpdateDrug)
	r.DELETE("drug/:id", h.DeleteDrug)

	// JOURNAL

	r.POST("journal", h.CreateJournal)
	r.GET("journal/:id", h.GetJournalByID)
	r.GET("journal", h.GetJournalsList)
	r.PUT("journal/:id", h.UpdateJournal)
	r.DELETE("journal/:id", h.DeleteJournal)

	// ORDER DRUG

	r.POST("order_drug", h.CreateOrderDrug)
	r.GET("order_drug/:id", h.GetOrderDrugByID)
	r.GET("order_drug", h.GetOrderDrugsList)
	r.PUT("order_drug/:id", h.UpdateOrderDrug)
	r.DELETE("order_drug/:id", h.DeleteOrderDrug)

	// ORDERS

	r.POST("orders", h.CreateOrders)
	r.GET("orders/:id", h.GetOrdersByID)
	r.GET("orders", h.GetOrderssList)
	r.PUT("orders/:id", h.UpdateOrders)
	r.DELETE("orders/:id", h.DeleteOrders)

	// PHARMACIST

	r.POST("pharmacist", h.CreatePharmacist)
	r.GET("pharmacist/:id", h.GetPharmacistByID)
	r.GET("pharmacist", h.GetPharmacistsList)
	r.PUT("pharmacist/:id", h.UpdatePharmacist)
	r.DELETE("pharmacist/:id", h.DeletePharmacist)

	// QUEUE

	r.POST("queue", h.CreateQueue)
	r.GET("queue/:id", h.GetQueueByID)
	r.GET("queue", h.GetQueuesList)
	r.PUT("queue/:id", h.UpdateQueue)
	r.DELETE("queue/:id", h.DeleteQueue)

	// SUPER ADMIN

	r.POST("super_admin", h.CreateSuperAdmin)
	r.GET("super_admin/:id", h.GetSuperAdminByID)
	r.GET("super_admin", h.GetSuperAdminsList)
	r.PUT("super_admin/:id", h.UpdateSuperAdmin)
	r.DELETE("super_admin/:id", h.DeleteSuperAdmin)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r

}
