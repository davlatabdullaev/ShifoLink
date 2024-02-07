package handler

import (
	"context"
	"errors"
	"net/http"
	"shifolink/api/models"
	"shifolink/pkg/check"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateClinicAdmin godoc
// @Router       /clinic_admin [POST]
// @Summary      Create a new clinic admin
// @Description  Create a new clinic admin
// @Tags         clinic_admin
// @Accept       json
// @Produce      json
// @Param        clinic_admin  body  models.CreateClinicAdmin  true  "clinic admin data"
// @Success      201  {object}  models.ClinicAdmin
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateClinicAdmin(c *gin.Context) {
	createClinicAdmin := models.CreateClinicAdmin{}

	if err := c.ShouldBindJSON(&createClinicAdmin); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
	}

	id, err := h.storage.ClinicAdmin().Create(context.Background(), createClinicAdmin)
	if err != nil {
		handleResponse(c, "error while creating clinic admin", http.StatusInternalServerError, err)
		return
	}

	clinicAdmin, err := h.storage.ClinicAdmin().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while get clinic admin", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, clinicAdmin)

}

// GetClinicAdminByID godoc
// @Router       /clinic_admin/{id} [GET]
// @Summary      Get clinic admin by id
// @Description  Get clinic admin by id
// @Tags         clinic_admin
// @Accept       json
// @Produce      json
// @Param        id path string true "clinic_admin"
// @Success      200  {object}  models.ClinicAdmin
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetClinicAdminByID(c *gin.Context) {

	var err error

	uid := c.Param("id")

	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "invalid uuid type ", http.StatusBadRequest, err.Error())
		return
	}

	clinicAdmin, err := h.storage.ClinicAdmin().Get(context.Background(), models.PrimaryKey{
		ID: id.String(),
	})
	if err != nil {
		handleResponse(c, "error while get clinic admin by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, clinicAdmin)

}

// GetClinicAdminsList godoc
// @Router       /clinic_admin [GET]
// @Summary      Get clinic admins list
// @Description  Get clinic admins list
// @Tags         clinic_admin
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      200  {object}  models.ClinicAdminsResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetClinicAdminsList(c *gin.Context) {

	var (
		page, limit int
		search      string
		err         error
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, "error while parsing page ", http.StatusBadRequest, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}

	search = c.Query("search")

	response, err := h.storage.ClinicAdmin().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

	if err != nil {
		handleResponse(c, "error while getting clinic admin", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, response)

}

// UpdateClinicAdmin godoc
// @Router       /clinic_admin/{id} [PUT]
// @Summary      Update clinic admin by id
// @Description  Update clinic admin by id
// @Tags         clinic_admin
// @Accept       json
// @Produce      json
// @Param        id path string true "clinic admin id"
// @Param        clinic_admin body models.UpdateClinicAdmin true "clinic_admin"
// @Success      200  {object}  models.ClinicAdmin
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateClinicAdmin(c *gin.Context) {
	updateClinicAdmin := models.UpdateClinicAdmin{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	//updateAuthor.ID = uid

	if err := c.ShouldBindJSON(&updateClinicAdmin); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.ClinicAdmin().Update(context.Background(), updateClinicAdmin)
	if err != nil {
		handleResponse(c, "error while updating clinic admin", http.StatusInternalServerError, err.Error())
		return
	}

	clinicAdmin, err := h.storage.ClinicAdmin().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while getting clinic admin by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, clinicAdmin)

}

// DeleteClinicAdmin godoc
// @Router       /clinic_admin/{id} [DELETE]
// @Summary      Delete clinic admin
// @Description  Delete clinic admin
// @Tags         clinic_admin
// @Accept       json
// @Produce      json
// @Param        id path string true "clinic_admin id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteClinicAdmin(c *gin.Context) {

	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err := h.storage.ClinicAdmin().Delete(context.Background(), id.String()); err != nil {
		handleResponse(c, "error while deleting clinic admin by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data succesfully deleted")

}

// UpdateClinicAdminPassword godoc
// @Router       /clinic_admin/{id} [PATCH]
// @Summary      Update clinic admin password
// @Description  update clinic admin password
// @Tags         clinic_admin
// @Accept       json
// @Produce      json
// @Param 		 id path string true "clinic admin"
// @Param        clinic_admin body models.UpdateClinicAdminPassword true "clinic_admin"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateClinicAdminPassword(c *gin.Context) {

	updateClinicAdminPassword := models.UpdateClinicAdminPassword{}

	if err := c.ShouldBindJSON(&updateClinicAdminPassword); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	uid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		handleResponse(c, "error while parsing uuid", http.StatusBadRequest, err.Error())
		return
	}

	updateClinicAdminPassword.ID = uid.String()

	oldPassword, err := h.storage.ClinicAdmin().GetPassword(c, updateClinicAdminPassword.ID)
	if err != nil {
		handleResponse(c, "error while getting password by id", http.StatusInternalServerError, err.Error())
		return
	}

	if oldPassword != updateClinicAdminPassword.OldPassword{
		handleResponse(c, "old password is not correct", http.StatusBadRequest, "old password is not correct")
		return
	}

	if err = check.ValidatePassword(updateClinicAdminPassword.NewPassword); err != nil {
		handleResponse(c, "new password is weak", http.StatusBadRequest, err.Error())
		return
	} 


	if err = h.storage.ClinicAdmin().UpdatePassword(context.Background(), updateClinicAdminPassword); err != nil {
		handleResponse(c, "error while updating clinic admin password", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "password successfully updated")
}
