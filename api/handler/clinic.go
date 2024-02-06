package handler

import (
	"context"
	"errors"
	"net/http"
	"shifolink/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateClinic godoc
// @Router       /clinic [POST]
// @Summary      Create a new clinic 
// @Description  Create a new clinic 
// @Tags         clinic
// @Accept       json
// @Produce      json
// @Param        clinic body  models.CreateClinic  true  "clinic data"
// @Success      201  {object}  models.Clinic
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateClinic(c *gin.Context) {
	createClinic := models.CreateClinic{}

	if err := c.ShouldBindJSON(&createClinic); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
	}

	id, err := h.storage.Clinic().Create(context.Background(), createClinic)
	if err != nil {
		handleResponse(c, "error while creating clinic", http.StatusInternalServerError, err)
		return
	}

	clinic, err := h.storage.Clinic().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while get clinic ", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, clinic)

}

// GetClinicByID godoc
// @Router       /clinic/{id} [GET]
// @Summary      Get clinic by id
// @Description  Get clinic by id
// @Tags         clinic
// @Accept       json
// @Produce      json
// @Param        id path string true "clinic"
// @Success      200  {object}  models.Clinic
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetClinicByID(c *gin.Context) {

	var err error

	uid := c.Param("id")

	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "invalid uuid type ", http.StatusBadRequest, err.Error())
		return
	}

	clinic, err := h.storage.Clinic().Get(context.Background(), models.PrimaryKey{
		ID: id.String(),
	})
	if err != nil {
		handleResponse(c, "error while get clinic by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, clinic)

}

// GetClinicsList godoc
// @Router       /clinic [GET]
// @Summary      Get clinics list
// @Description  Get clinics list
// @Tags         clinic
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      200  {object}  models.ClinicsResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetClinicsList(c *gin.Context) {

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

	response, err := h.storage.Clinic().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

	if err != nil {
		handleResponse(c, "error while getting clinic", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, response)

}

// UpdateClinic godoc
// @Router       /clinic/{id} [PUT]
// @Summary      Update clinic by id
// @Description  Update clinic by id
// @Tags         clinic
// @Accept       json
// @Produce      json
// @Param        id path string true "clinic id"
// @Param        clinic body models.UpdateClinic true "clinic"
// @Success      200  {object}  models.Clinic
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateClinic(c *gin.Context) {
	updateClinic := models.UpdateClinic{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	//updateAuthor.ID = uid

	if err := c.ShouldBindJSON(&updateClinic); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Clinic().Update(context.Background(), updateClinic)
	if err != nil {
		handleResponse(c, "error while updating clinic ", http.StatusInternalServerError, err.Error())
		return
	}

	clinic, err := h.storage.Clinic().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while getting clinic by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, clinic)

}

// DeleteClinic godoc
// @Router       /clinic/{id} [DELETE]
// @Summary      Delete clinic 
// @Description  Delete clinic 
// @Tags         clinic
// @Accept       json
// @Produce      json
// @Param        id path string true "clinic id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteClinic(c *gin.Context) {

	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err := h.storage.Clinic().Delete(context.Background(), id.String()); err != nil {
		handleResponse(c, "error while deleting clinic by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data succesfully deleted")

}
