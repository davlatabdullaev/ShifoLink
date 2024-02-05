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

// CreateDrugStoreBranch godoc
// @Router       /drug_store_branch [POST]
// @Summary      Create a new drug store branch
// @Description  Create a new drug store branch
// @Tags         drug_store_branch
// @Accept       json
// @Produce      json
// @Param        drug_store_branch body  models.CreateDrugStoreBranch true  "drug_store_branch data"
// @Success      201  {object}  models.DrugStoreBranch
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateDrugStoreBranch(c *gin.Context) {
	createDrugStoreBranch := models.CreateDrugStoreBranch{}

	if err := c.ShouldBindJSON(&createDrugStoreBranch); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
	}

	id, err := h.storage.DrugStoreBranch().Create(context.Background(), createDrugStoreBranch)
	if err != nil {
		handleResponse(c, "error while creating drug store branch ", http.StatusInternalServerError, err)
		return
	}

	drugStoreBranch, err := h.storage.DrugStoreBranch().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while get drug store branch ", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, drugStoreBranch)

}

// GetDrugStoreBranchByID godoc
// @Router       /drug_store_branch/{id} [GET]
// @Summary      Get drug store branch by id
// @Description  Get drug store branch by id
// @Tags         drug_store_branch
// @Accept       json
// @Produce      json
// @Param        id path string true "drug_store_branch"
// @Success      200  {object}  models.DrugStoreBranch
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetDrugStoreBranchByID(c *gin.Context) {

	var err error

	uid := c.Param("id")

	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "invalid uuid type ", http.StatusBadRequest, err.Error())
		return
	}

	drugStoreBranch, err := h.storage.DrugStoreBranch().Get(context.Background(), models.PrimaryKey{
		ID: id.String(),
	})
	if err != nil {
		handleResponse(c, "error while get drug store branch by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, drugStoreBranch)

}

// GetDrugStoreBranchsList godoc
// @Router       /drug_store_branchs [GET]
// @Summary      Get drug store branchs list
// @Description  Get drug store branchs list
// @Tags         drug_store_branch
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      200  {object}  models.DrugStoreBranchsResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetDrugStoreBranchsList(c *gin.Context) {

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

	response, err := h.storage.DrugStoreBranch().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

	if err != nil {
		handleResponse(c, "error while getting drug store branch ", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, response)

}

// UpdateDrugStoreBranch godoc
// @Router       /drug_store_branch/{id} [PUT]
// @Summary      Update drug store branch by id
// @Description  Update drug store branch by id
// @Tags         drug_store_branch
// @Accept       json
// @Produce      json
// @Param        id path string true "drug_store_branch id"
// @Param        drug_store_branch body models.UpdateDrugStoreBranch true "drug_store_branch"
// @Success      200  {object}  models.DrugStoreBranch
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateDrugStoreBranch(c *gin.Context) {
	updateDrugStoreBranch := models.UpdateDrugStoreBranch{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	//updateAuthor.ID = uid

	if err := c.ShouldBindJSON(&updateDrugStoreBranch); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.DrugStoreBranch().Update(context.Background(), updateDrugStoreBranch)
	if err != nil {
		handleResponse(c, "error while updating drug store branch ", http.StatusInternalServerError, err.Error())
		return
	}

	drugStoreBranch, err := h.storage.DrugStoreBranch().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while getting drug store branch by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, drugStoreBranch)

}

// DeleteDrugStoreBranch godoc
// @Router       /drug_store_branch/{id} [DELETE]
// @Summary      Delete drug store branch
// @Description  Delete drug store branch
// @Tags         drug_store_branch
// @Accept       json
// @Produce      json
// @Param        id path string true "drug_store_branch id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteDrugStoreBranch(c *gin.Context) {

	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err := h.storage.DrugStoreBranch().Delete(context.Background(), id.String()); err != nil {
		handleResponse(c, "error while deleting drug store branch by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data succesfully deleted")

}
