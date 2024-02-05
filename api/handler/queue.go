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

// CreateQueue godoc
// @Router       /queue [POST]
// @Summary      Create a new Queue
// @Description  Create a new Queue
// @Tags         queue
// @Accept       json
// @Produce      json
// @Param        Queue body  models.CreateQueue true  "Queue data"
// @Success      201  {object}  models.Queue
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateQueue(c *gin.Context) {
	createQueue := models.CreateQueue{}

	if err := c.ShouldBindJSON(&createQueue); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
	}

	id, err := h.storage.Queue().Create(context.Background(), createQueue)
	if err != nil {
		handleResponse(c, "error while creating Queue ", http.StatusInternalServerError, err)
		return
	}

	Queue, err := h.storage.Queue().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while get Queue ", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, Queue)

}

// GetQueueByID godoc
// @Router       /queue/{id} [GET]
// @Summary      Get Queue by id
// @Description  Get Queue by id
// @Tags         queue
// @Accept       json
// @Produce      json
// @Param        id path string true "Queue"
// @Success      200  {object}  models.Queue
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetQueueByID(c *gin.Context) {

	var err error

	uid := c.Param("id")

	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "invalid uuid type ", http.StatusBadRequest, err.Error())
		return
	}

	Queue, err := h.storage.Queue().Get(context.Background(), models.PrimaryKey{
		ID: id.String(),
	})
	if err != nil {
		handleResponse(c, "error while get Queue by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, Queue)

}

// GetQueuesList godoc
// @Router       /queue [GET]
// @Summary      Get Queues list
// @Description  Get Queues list
// @Tags         queue
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      200  {object}  models.QueuesResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetQueuesList(c *gin.Context) {

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

	response, err := h.storage.Queue().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

	if err != nil {
		handleResponse(c, "error while getting Queue ", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, response)

}

// UpdateQueue godoc
// @Router       /queue/{id} [PUT]
// @Summary      Update Queue by id
// @Description  Update Queue by id
// @Tags         queue
// @Accept       json
// @Produce      json
// @Param        id path string true "Queue id"
// @Param        Queue body models.UpdateQueue true "Queue"
// @Success      200  {object}  models.Queue
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateQueue(c *gin.Context) {
	updateQueue := models.UpdateQueue{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	//updateAuthor.ID = uid

	if err := c.ShouldBindJSON(&updateQueue); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Queue().Update(context.Background(), updateQueue)
	if err != nil {
		handleResponse(c, "error while updating Queue ", http.StatusInternalServerError, err.Error())
		return
	}

	Queue, err := h.storage.Queue().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while getting Queue by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, Queue)

}

// DeleteQueue godoc
// @Router       /queue/{id} [DELETE]
// @Summary      Delete Queue
// @Description  Delete Queue
// @Tags         queue
// @Accept       json
// @Produce      json
// @Param        id path string true "Queue id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteQueue(c *gin.Context) {

	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err := h.storage.Queue().Delete(context.Background(), id.String()); err != nil {
		handleResponse(c, "error while deleting Queue  by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data succesfully deleted")

}
