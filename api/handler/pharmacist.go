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

// CreatePharmacist godoc
// @Router       /pharmacist [POST]
// @Summary      Create a new Pharmacist
// @Description  Create a new Pharmacist
// @Tags         pharmacist
// @Accept       json
// @Produce      json
// @Param        Pharmacist body  models.CreatePharmacist true  "Pharmacist data"
// @Success      201  {object}  models.Pharmacist
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreatePharmacist(c *gin.Context) {
	createPharmacist := models.CreatePharmacist{}

	if err := c.ShouldBindJSON(&createPharmacist); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
	}

	id, err := h.storage.Pharmacist().Create(context.Background(), createPharmacist)
	if err != nil {
		handleResponse(c, "error while creating Pharmacist ", http.StatusInternalServerError, err)
		return
	}

	Pharmacist, err := h.storage.Pharmacist().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while get Pharmacist ", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, Pharmacist)

}

// GetPharmacistByID godoc
// @Router       /pharmacist/{id} [GET]
// @Summary      Get Pharmacist by id
// @Description  Get Pharmacist by id
// @Tags         pharmacist
// @Accept       json
// @Produce      json
// @Param        id path string true "Pharmacist"
// @Success      200  {object}  models.Pharmacist
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetPharmacistByID(c *gin.Context) {

	var err error

	uid := c.Param("id")

	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "invalid uuid type ", http.StatusBadRequest, err.Error())
		return
	}

	Pharmacist, err := h.storage.Pharmacist().Get(context.Background(), models.PrimaryKey{
		ID: id.String(),
	})
	if err != nil {
		handleResponse(c, "error while get Pharmacist by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, Pharmacist)

}

// GetPharmacistsList godoc
// @Router       /pharmacist [GET]
// @Summary      Get Pharmacists list
// @Description  Get Pharmacists list
// @Tags         pharmacist
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      200  {object}  models.PharmacistsResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetPharmacistsList(c *gin.Context) {

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

	response, err := h.storage.Pharmacist().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

	if err != nil {
		handleResponse(c, "error while getting Pharmacist ", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, response)

}

// UpdatePharmacist godoc
// @Router       /pharmacist/{id} [PUT]
// @Summary      Update Pharmacist by id
// @Description  Update Pharmacist by id
// @Tags         pharmacist
// @Accept       json
// @Produce      json
// @Param        id path string true "Pharmacist id"
// @Param        Pharmacist body models.UpdatePharmacist true "Pharmacist"
// @Success      200  {object}  models.Pharmacist
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdatePharmacist(c *gin.Context) {
	updatePharmacist := models.UpdatePharmacist{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	//updateAuthor.ID = uid

	if err := c.ShouldBindJSON(&updatePharmacist); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Pharmacist().Update(context.Background(), updatePharmacist)
	if err != nil {
		handleResponse(c, "error while updating Pharmacist ", http.StatusInternalServerError, err.Error())
		return
	}

	Pharmacist, err := h.storage.Pharmacist().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while getting Pharmacist by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, Pharmacist)

}

// DeletePharmacist godoc
// @Router       /pharmacist/{id} [DELETE]
// @Summary      Delete Pharmacist
// @Description  Delete Pharmacist
// @Tags         pharmacist
// @Accept       json
// @Produce      json
// @Param        id path string true "Pharmacist id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeletePharmacist(c *gin.Context) {

	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err := h.storage.Pharmacist().Delete(context.Background(), id.String()); err != nil {
		handleResponse(c, "error while deleting Pharmacist  by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data succesfully deleted")

}
