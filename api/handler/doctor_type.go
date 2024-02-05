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

// CreateDoctorType godoc
// @Router       /doctor_type [POST]
// @Summary      Create a new doctor type
// @Description  Create a new doctor type 
// @Tags         doctor_type
// @Accept       json
// @Produce      json
// @Param        doctor_type body  models.CreateDoctorType true  "doctor type data"
// @Success      201  {object}  models.DoctorType
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateDoctorType(c *gin.Context) {
	createDoctorType := models.CreateDoctorType{}

	if err := c.ShouldBindJSON(&createDoctorType); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
	}

	id, err := h.storage.DoctorType().Create(context.Background(), createDoctorType)
	if err != nil {
		handleResponse(c, "error while creating doctor type", http.StatusInternalServerError, err)
		return
	}

	dtype, err := h.storage.DoctorType().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while get doctor type ", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, dtype)

}

// GetDoctorTypeByID godoc
// @Router       /doctor_type/{id} [GET]
// @Summary      Get doctor type by id
// @Description  Get doctor type by id
// @Tags         doctor_type
// @Accept       json
// @Produce      json
// @Param        id path string true "doctor_type"
// @Success      200  {object}  models.DoctorType
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetDoctorTypeByID(c *gin.Context) {

	var err error

	uid := c.Param("id")

	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "invalid uuid type ", http.StatusBadRequest, err.Error())
		return
	}

	dtype, err := h.storage.DoctorType().Get(context.Background(), models.PrimaryKey{
		ID: id.String(),
	})
	if err != nil {
		handleResponse(c, "error while get doctor type by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, dtype)

}

// GetDoctorTypesList godoc
// @Router       /doctor_types [GET]
// @Summary      Get doctor_types list
// @Description  Get doctor_types list
// @Tags         doctor_type
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      200  {object}  models.DoctorTypesResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetDoctorTypesList(c *gin.Context) {

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

	response, err := h.storage.DoctorType().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

	if err != nil {
		handleResponse(c, "error while getting doctor type", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, response)

}

// UpdateDoctorType godoc
// @Router       /doctor_type/{id} [PUT]
// @Summary      Update doctor type by id
// @Description  Update doctor type by id
// @Tags         doctor_type
// @Accept       json
// @Produce      json
// @Param        id path string true "doctor type id"
// @Param        doctor_type body models.UpdateDoctorType true "doctor_type"
// @Success      200  {object}  models.DoctorType
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateDoctorType(c *gin.Context) {
	updateDoctorType := models.UpdateDoctorType{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	//updateAuthor.ID = uid

	if err := c.ShouldBindJSON(&updateDoctorType); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.DoctorType().Update(context.Background(), updateDoctorType)
	if err != nil {
		handleResponse(c, "error while updating doctor type ", http.StatusInternalServerError, err.Error())
		return
	}

	dtype, err := h.storage.DoctorType().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while getting doctor type by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, dtype)

}

// DeleteDoctorType godoc
// @Router       /doctor_type/{id} [DELETE]
// @Summary      Delete doctor type
// @Description  Delete doctor type
// @Tags         doctor_type
// @Accept       json
// @Produce      json
// @Param        id path string true "doctor_type id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteDoctorType(c *gin.Context) {

	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err := h.storage.DoctorType().Delete(context.Background(), id.String()); err != nil {
		handleResponse(c, "error while deleting doctor type by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data succesfully deleted")

}
