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

// CreateClinicBranch godoc
// @Router       /clinic_branch [POST]
// @Summary      Create a new clinic branch
// @Description  Create a new clinic branch
// @Tags         clinic_branch
// @Accept       json
// @Produce      json
// @Param        clinic_branch body  models.CreateClinicBranch  true  "clinic branch data"
// @Success      201  {object}  models.ClinicBranch
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateClinicBranch(c *gin.Context) {
	createClinicBranch := models.CreateClinicBranch{}

	if err := c.ShouldBindJSON(&createClinicBranch); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
	}

	id, err := h.storage.ClinicBranch().Create(context.Background(), createClinicBranch)
	if err != nil {
		handleResponse(c, "error while creating clinic branch", http.StatusInternalServerError, err)
		return
	}

	clinicBranch, err := h.storage.ClinicBranch().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while get clinic branch", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, clinicBranch)

}

// GetClinicBranchByID godoc
// @Router       /clinic_branch/{id} [GET]
// @Summary      Get clinic branch by id
// @Description  Get clinic branch by id
// @Tags         clinic_branch
// @Accept       json
// @Produce      json
// @Param        id path string true "clinic_branch"
// @Success      200  {object}  models.ClinicBranch
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetClinicBranchByID(c *gin.Context) {

	var err error

	uid := c.Param("id")

	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "invalid uuid type ", http.StatusBadRequest, err.Error())
		return
	}

	clinicBranch, err := h.storage.ClinicBranch().Get(context.Background(), models.PrimaryKey{
		ID: id.String(),
	})
	if err != nil {
		handleResponse(c, "error while get clinic branch by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, clinicBranch)

}

// GetClinicBranchsList godoc
// @Router       /clinic_branchs [GET]
// @Summary      Get clinic branchs list
// @Description  Get clinic branchs list
// @Tags         clinic_branch
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      200  {object}  models.ClinicBranchsResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetClinicBranchsList(c *gin.Context) {

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

	response, err := h.storage.ClinicBranch().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

	if err != nil {
		handleResponse(c, "error while getting clinic branch", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, response)

}

// UpdateClinicBranch godoc
// @Router       /clinic_branch/{id} [PUT]
// @Summary      Update clinic branch by id
// @Description  Update clinic branch by id
// @Tags         clinic_branch
// @Accept       json
// @Produce      json
// @Param        id path string true "clinic branch id"
// @Param        clinic_branch body models.UpdateClinicBranch true "clinic_branch"
// @Success      200  {object}  models.ClinicBranch
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateClinicBranch(c *gin.Context) {
	updateClinicBranch := models.UpdateClinicBranch{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	//updateAuthor.ID = uid

	if err := c.ShouldBindJSON(&updateClinicBranch); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.ClinicBranch().Update(context.Background(), updateClinicBranch)
	if err != nil {
		handleResponse(c, "error while updating clinic branch", http.StatusInternalServerError, err.Error())
		return
	}

	clinicBranch, err := h.storage.ClinicBranch().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while getting clinic branch by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, clinicBranch)

}

// DeleteClinicBranch godoc
// @Router       /clinic_branch/{id} [DELETE]
// @Summary      Delete clinic branch
// @Description  Delete clinic branch
// @Tags         clinic_branch
// @Accept       json
// @Produce      json
// @Param        id path string true "clinic_branch id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteClinicBranch(c *gin.Context) {

	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err := h.storage.ClinicBranch().Delete(context.Background(), id.String()); err != nil {
		handleResponse(c, "error while deleting clinic branch by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data succesfully deleted")

}
