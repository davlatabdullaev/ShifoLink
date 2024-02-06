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

// CreateJournal godoc
// @Router       /journal [POST]
// @Summary      Create a new journal
// @Description  Create a new journal
// @Tags         journal
// @Accept       json
// @Produce      json
// @Param        journal body  models.CreateJournal true  "journal data"
// @Success      201  {object}  models.Journal
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateJournal(c *gin.Context) {
	createJournal := models.CreateJournal{}

	if err := c.ShouldBindJSON(&createJournal); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
	}

	id, err := h.storage.Journal().Create(context.Background(), createJournal)
	if err != nil {
		handleResponse(c, "error while creating journal ", http.StatusInternalServerError, err)
		return
	}

	journal, err := h.storage.Journal().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while get journal ", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, journal)

}

// GetJournalByID godoc
// @Router       /journal/{id} [GET]
// @Summary      Get journal by id
// @Description  Get journal by id
// @Tags         journal
// @Accept       json
// @Produce      json
// @Param        id path string true "journal"
// @Success      200  {object}  models.Journal
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetJournalByID(c *gin.Context) {

	var err error

	uid := c.Param("id")

	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "invalid uuid type ", http.StatusBadRequest, err.Error())
		return
	}

	journal, err := h.storage.Journal().Get(context.Background(), models.PrimaryKey{
		ID: id.String(),
	})
	if err != nil {
		handleResponse(c, "error while get journal by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, journal)

}

// GetJournalsList godoc
// @Router       /journal [GET]
// @Summary      Get journals list
// @Description  Get journals list
// @Tags         journal
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      200  {object}  models.JournalsResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetJournalsList(c *gin.Context) {

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

	response, err := h.storage.Journal().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

	if err != nil {
		handleResponse(c, "error while getting journal ", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, response)

}

// UpdateJournal godoc
// @Router       /journal/{id} [PUT]
// @Summary      Update journal by id
// @Description  Update journal by id
// @Tags         journal
// @Accept       json
// @Produce      json
// @Param        id path string true "journal id"
// @Param        journal body models.UpdateJournal true "journal"
// @Success      200  {object}  models.Journal
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateJournal(c *gin.Context) {
	updateJournal := models.UpdateJournal{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	//updateAuthor.ID = uid

	if err := c.ShouldBindJSON(&updateJournal); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Journal().Update(context.Background(), updateJournal)
	if err != nil {
		handleResponse(c, "error while updating journal ", http.StatusInternalServerError, err.Error())
		return
	}

	journal, err := h.storage.Journal().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while getting journal by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, journal)

}

// DeleteJournal godoc
// @Router       /journal/{id} [DELETE]
// @Summary      Delete journal
// @Description  Delete journal
// @Tags         journal
// @Accept       json
// @Produce      json
// @Param        id path string true "journal id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteJournal(c *gin.Context) {

	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err := h.storage.Journal().Delete(context.Background(), id.String()); err != nil {
		handleResponse(c, "error while deleting journal  by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data succesfully deleted")

}
