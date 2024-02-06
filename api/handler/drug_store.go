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

// CreateDrugStore godoc
// @Router       /drug_store [POST]
// @Summary      Create a new drug store
// @Description  Create a new drug store
// @Tags         drug_store
// @Accept       json
// @Produce      json
// @Param        drug_store body  models.CreateDrugStore true  "drug_store data"
// @Success      201  {object}  models.DrugStore
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateDrugStore(c *gin.Context) {
	createDrugStore := models.CreateDrugStore{}

	if err := c.ShouldBindJSON(&createDrugStore); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
	}

	id, err := h.storage.DrugStore().Create(context.Background(), createDrugStore)
	if err != nil {
		handleResponse(c, "error while creating drug store ", http.StatusInternalServerError, err)
		return
	}

	drugStore, err := h.storage.DrugStore().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while get drug store ", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, drugStore)

}

// GetDrugStoreByID godoc
// @Router       /drug_store/{id} [GET]
// @Summary      Get drug store by id
// @Description  Get drug store by id
// @Tags         drug_store
// @Accept       json
// @Produce      json
// @Param        id path string true "drug_store"
// @Success      200  {object}  models.DrugStore
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetDrugStoreByID(c *gin.Context) {

	var err error

	uid := c.Param("id")

	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "invalid uuid type ", http.StatusBadRequest, err.Error())
		return
	}

	drugStore, err := h.storage.DrugStore().Get(context.Background(), models.PrimaryKey{
		ID: id.String(),
	})
	if err != nil {
		handleResponse(c, "error while get drug store by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, drugStore)

}

// GetDrugStoresList godoc
// @Router       /drug_store [GET]
// @Summary      Get drug stores list
// @Description  Get drug stores list
// @Tags         drug_store
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      200  {object}  models.DrugStoresResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetDrugStoresList(c *gin.Context) {

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

	response, err := h.storage.DrugStore().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

	if err != nil {
		handleResponse(c, "error while getting drug store ", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, response)

}

// UpdateDrugStore godoc
// @Router       /drug_store/{id} [PUT]
// @Summary      Update drug store by id
// @Description  Update drug store by id
// @Tags         drug_store
// @Accept       json
// @Produce      json
// @Param        id path string true "drug_store id"
// @Param        drug_store body models.UpdateDrugStore true "drug_store"
// @Success      200  {object}  models.DrugStore
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateDrugStore(c *gin.Context) {
	updateDrugStore := models.UpdateDrugStore{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	//updateAuthor.ID = uid

	if err := c.ShouldBindJSON(&updateDrugStore); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.DrugStore().Update(context.Background(), updateDrugStore)
	if err != nil {
		handleResponse(c, "error while updating drug store ", http.StatusInternalServerError, err.Error())
		return
	}

	drugStore, err := h.storage.DrugStore().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while getting drug store by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, drugStore)

}

// DeleteDrugStore godoc
// @Router       /drug_store/{id} [DELETE]
// @Summary      Delete drug store
// @Description  Delete drug store
// @Tags         drug_store
// @Accept       json
// @Produce      json
// @Param        id path string true "drug_store id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteDrugStore(c *gin.Context) {

	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err := h.storage.DrugStore().Delete(context.Background(), id.String()); err != nil {
		handleResponse(c, "error while deleting drug store  by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data succesfully deleted")

}
