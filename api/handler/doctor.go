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

// CreateDoctor godoc
// @Router       /doctor [POST]
// @Summary      Create a new doctor
// @Description  Create a new doctor
// @Tags         doctor
// @Accept       json
// @Produce      json
// @Param        doctor body  models.CreateDoctor true  "doctor data"
// @Success      201  {object}  models.Doctor
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateDoctor(c *gin.Context) {
	createDoctor := models.CreateDoctor{}

	if err := c.ShouldBindJSON(&createDoctor); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
	}

	id, err := h.storage.Doctor().Create(context.Background(), createDoctor)
	if err != nil {
		handleResponse(c, "error while creating doctor ", http.StatusInternalServerError, err)
		return
	}

	doctor, err := h.storage.Doctor().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while get doctor ", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, doctor)

}

// GetDoctorByID godoc
// @Router       /doctor/{id} [GET]
// @Summary      Get doctor by id
// @Description  Get doctor by id
// @Tags         doctor
// @Accept       json
// @Produce      json
// @Param        id path string true "doctor"
// @Success      200  {object}  models.Doctor
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetDoctorByID(c *gin.Context) {

	var err error

	uid := c.Param("id")

	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "invalid uuid type ", http.StatusBadRequest, err.Error())
		return
	}

	doctor, err := h.storage.Doctor().Get(context.Background(), models.PrimaryKey{
		ID: id.String(),
	})
	if err != nil {
		handleResponse(c, "error while get doctor by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, doctor)

}

// GetDoctorsList godoc
// @Router       /doctor [GET]
// @Summary      Get doctors list
// @Description  Get doctors list
// @Tags         doctor
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      200  {object}  models.DoctorsResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetDoctorsList(c *gin.Context) {

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

	response, err := h.storage.Doctor().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

	if err != nil {
		handleResponse(c, "error while getting doctor ", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, response)

}

// UpdateDoctor godoc
// @Router       /doctor/{id} [PUT]
// @Summary      Update doctor by id
// @Description  Update doctor by id
// @Tags         doctor
// @Accept       json
// @Produce      json
// @Param        id path string true "doctor id"
// @Param        doctor body models.UpdateDoctor true "doctor"
// @Success      200  {object}  models.Doctor
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateDoctor(c *gin.Context) {
	updateDoctor := models.UpdateDoctor{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	//updateAuthor.ID = uid

	if err := c.ShouldBindJSON(&updateDoctor); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Doctor().Update(context.Background(), updateDoctor)
	if err != nil {
		handleResponse(c, "error while updating doctor ", http.StatusInternalServerError, err.Error())
		return
	}

	doctor, err := h.storage.Doctor().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while getting doctor by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, doctor)

}

// DeleteDoctor godoc
// @Router       /doctor/{id} [DELETE]
// @Summary      Delete doctor
// @Description  Delete doctor
// @Tags         doctor
// @Accept       json
// @Produce      json
// @Param        id path string true "doctor id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteDoctor(c *gin.Context) {

	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err := h.storage.Doctor().Delete(context.Background(), id.String()); err != nil {
		handleResponse(c, "error while deleting doctor by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data succesfully deleted")

}

// UpdateDoctorPassword godoc
// @Router       /doctor/{id} [PATCH]
// @Summary      Update doctor password
// @Description  update doctor password
// @Tags         doctor
// @Accept       json
// @Produce      json
// @Param 		 id path string true "doctor"
// @Param        doctor body models.UpdateDoctorPassword true "doctor"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateDoctorPassword(c *gin.Context) {
	updateDoctorPassword := models.UpdateDoctorPassword{}

	if err := c.ShouldBindJSON(&updateDoctorPassword); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	uid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		handleResponse(c, "error while parsing uuid", http.StatusBadRequest, err.Error())
		return
	}

	updateDoctorPassword.ID = uid.String()

	if err = h.storage.Doctor().UpdatePassword(context.Background(), updateDoctorPassword); err != nil {
		handleResponse(c, "error while updating doctor password", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "password successfully updated")
}
