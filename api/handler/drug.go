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

// CreateDrug godoc
// @Router       /drug [POST]
// @Summary      Create a new drug
// @Description  Create a new drug
// @Tags         drug
// @Accept       json
// @Produce      json
// @Param        drug body  models.CreateDrug true  "drug data"
// @Success      201  {object}  models.Drug
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateDrug(c *gin.Context) {
	createDrug := models.CreateDrug{}

	if err := c.ShouldBindJSON(&createDrug); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
	}

	id, err := h.storage.Drug().Create(context.Background(), createDrug)
	if err != nil {
		handleResponse(c, "error while creating drug ", http.StatusInternalServerError, err)
		return
	}

	drug, err := h.storage.Drug().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while get drug ", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, drug)

}

// GetDrugByID godoc
// @Router       /drug/{id} [GET]
// @Summary      Get drug by id
// @Description  Get drug by id
// @Tags         drug
// @Accept       json
// @Produce      json
// @Param        id path string true "drug"
// @Success      200  {object}  models.Drug
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetDrugByID(c *gin.Context) {

	var err error

	uid := c.Param("id")

	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "invalid uuid type ", http.StatusBadRequest, err.Error())
		return
	}

	drug, err := h.storage.Drug().Get(context.Background(), models.PrimaryKey{
		ID: id.String(),
	})
	if err != nil {
		handleResponse(c, "error while get drug by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, drug)

}

// GetDrugsList godoc
// @Router       /drugs [GET]
// @Summary      Get drugs list
// @Description  Get drugs list
// @Tags         drug
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      200  {object}  models.DrugsResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetDrugsList(c *gin.Context) {

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

	response, err := h.storage.Drug().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

	if err != nil {
		handleResponse(c, "error while getting drug ", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, response)

}

// UpdateDrug godoc
// @Router       /drug/{id} [PUT]
// @Summary      Update drug by id
// @Description  Update drug by id
// @Tags         drug
// @Accept       json
// @Produce      json
// @Param        id path string true "drug id"
// @Param        drug body models.UpdateDrug true "drug"
// @Success      200  {object}  models.Drug
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateDrug(c *gin.Context) {
	updateDrug := models.UpdateDrug{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	//updateAuthor.ID = uid

	if err := c.ShouldBindJSON(&updateDrug); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Drug().Update(context.Background(), updateDrug)
	if err != nil {
		handleResponse(c, "error while updating drug ", http.StatusInternalServerError, err.Error())
		return
	}

	Drug, err := h.storage.Drug().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while getting drug by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, Drug)

}

// DeleteDrug godoc
// @Router       /drug/{id} [DELETE]
// @Summary      Delete drug
// @Description  Delete drug
// @Tags         drug
// @Accept       json
// @Produce      json
// @Param        id path string true "drug id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteDrug(c *gin.Context) {

	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err := h.storage.Drug().Delete(context.Background(), id.String()); err != nil {
		handleResponse(c, "error while deleting drug  by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data succesfully deleted")

}
