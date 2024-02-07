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

// CreateSuperAdmin godoc
// @Router       /super_admin [POST]
// @Summary      Create a new SuperAdmin
// @Description  Create a new SuperAdmin
// @Tags         super_admin
// @Accept       json
// @Produce      json
// @Param        SuperAdmin body  models.CreateSuperAdmin true  "SuperAdmin data"
// @Success      201  {object}  models.SuperAdmin
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateSuperAdmin(c *gin.Context) {
	createSuperAdmin := models.CreateSuperAdmin{}

	if err := c.ShouldBindJSON(&createSuperAdmin); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
	}

	id, err := h.storage.SuperAdmin().Create(context.Background(), createSuperAdmin)
	if err != nil {
		handleResponse(c, "error while creating SuperAdmin ", http.StatusInternalServerError, err)
		return
	}

	SuperAdmin, err := h.storage.SuperAdmin().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while get SuperAdmin ", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, SuperAdmin)

}

// GetSuperAdminByID godoc
// @Router       /super_admin/{id} [GET]
// @Summary      Get SuperAdmin by id
// @Description  Get SuperAdmin by id
// @Tags         super_admin
// @Accept       json
// @Produce      json
// @Param        id path string true "SuperAdmin"
// @Success      200  {object}  models.SuperAdmin
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetSuperAdminByID(c *gin.Context) {

	var err error

	uid := c.Param("id")

	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "invalid uuid type ", http.StatusBadRequest, err.Error())
		return
	}

	SuperAdmin, err := h.storage.SuperAdmin().Get(context.Background(), models.PrimaryKey{
		ID: id.String(),
	})
	if err != nil {
		handleResponse(c, "error while get SuperAdmin by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, SuperAdmin)

}

// GetSuperAdminsList godoc
// @Router       /super_admin [GET]
// @Summary      Get SuperAdmins list
// @Description  Get SuperAdmins list
// @Tags         super_admin
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      200  {object}  models.SuperAdminsResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetSuperAdminsList(c *gin.Context) {

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

	response, err := h.storage.SuperAdmin().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

	if err != nil {
		handleResponse(c, "error while getting SuperAdmin ", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, response)

}

// UpdateSuperAdmin godoc
// @Router       /super_admin/{id} [PUT]
// @Summary      Update SuperAdmin by id
// @Description  Update SuperAdmin by id
// @Tags         super_admin
// @Accept       json
// @Produce      json
// @Param        id path string true "SuperAdmin id"
// @Param        SuperAdmin body models.UpdateSuperAdmin true "SuperAdmin"
// @Success      200  {object}  models.SuperAdmin
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateSuperAdmin(c *gin.Context) {
	updateSuperAdmin := models.UpdateSuperAdmin{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	//updateAuthor.ID = uid

	if err := c.ShouldBindJSON(&updateSuperAdmin); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.SuperAdmin().Update(context.Background(), updateSuperAdmin)
	if err != nil {
		handleResponse(c, "error while updating SuperAdmin ", http.StatusInternalServerError, err.Error())
		return
	}

	SuperAdmin, err := h.storage.SuperAdmin().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while getting SuperAdmin by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, SuperAdmin)

}

// DeleteSuperAdmin godoc
// @Router       /super_admin/{id} [DELETE]
// @Summary      Delete SuperAdmin
// @Description  Delete SuperAdmin
// @Tags         super_admin
// @Accept       json
// @Produce      json
// @Param        id path string true "SuperAdmin id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteSuperAdmin(c *gin.Context) {

	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err := h.storage.SuperAdmin().Delete(context.Background(), id.String()); err != nil {
		handleResponse(c, "error while deleting SuperAdmin  by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data succesfully deleted")

}

// UpdateSuperAdminPassword godoc
// @Router       /super_admin/{id} [PATCH]
// @Summary      Update super_admin password
// @Description  update super_admin password
// @Tags         super_admin
// @Accept       json
// @Produce      json
// @Param 		 id path string true "super_admin"
// @Param        super_admin body models.UpdateSuperAdminPassword true "super_admin"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateSuperAdminPassword(c *gin.Context) {
	updateSuperAdminPassword := models.UpdateSuperAdminPassword{}

	if err := c.ShouldBindJSON(&updateSuperAdminPassword); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	uid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		handleResponse(c, "error while parsing uuid", http.StatusBadRequest, err.Error())
		return
	}

	updateSuperAdminPassword.ID = uid.String()

	oldPassword, err := h.storage.SuperAdmin().GetPassword(c, updateSuperAdminPassword.ID)
	if err != nil {
		handleResponse(c, "error while getting password by id", http.StatusInternalServerError, err.Error())
		return
	}

	if oldPassword != updateSuperAdminPassword.OldPassword {
		handleResponse(c, "old password is not correct", http.StatusBadRequest, "old password is not correct")
		return
	}

	if err = check.ValidatePassword(updateSuperAdminPassword.NewPassword); err != nil {
		handleResponse(c, "new password is weak", http.StatusBadRequest, err.Error())
		return
	}

	if err = h.storage.SuperAdmin().UpdatePassword(context.Background(), updateSuperAdminPassword); err != nil {
		handleResponse(c, "error while updating super_admin password", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "password successfully updated")
}
