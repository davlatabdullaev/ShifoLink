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

// CreateOrders godoc
// @Router       /orders [POST]
// @Summary      Create a new Orders
// @Description  Create a new Orders
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        Orders body  models.CreateOrders true  "Orders data"
// @Success      201  {object}  models.Orders
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateOrders(c *gin.Context) {
	createOrders := models.CreateOrders{}

	if err := c.ShouldBindJSON(&createOrders); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
	}

	id, err := h.storage.Orders().Create(context.Background(), createOrders)
	if err != nil {
		handleResponse(c, "error while creating Orders ", http.StatusInternalServerError, err)
		return
	}

	orders, err := h.storage.Orders().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while get Orders ", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, orders)

}

// GetOrdersByID godoc
// @Router       /orders/{id} [GET]
// @Summary      Get Orders by id
// @Description  Get Orders by id
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        id path string true "Orders"
// @Success      200  {object}  models.Orders
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetOrdersByID(c *gin.Context) {

	var err error

	uid := c.Param("id")

	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "invalid uuid type ", http.StatusBadRequest, err.Error())
		return
	}

	orders, err := h.storage.Orders().Get(context.Background(), models.PrimaryKey{
		ID: id.String(),
	})
	if err != nil {
		handleResponse(c, "error while get Orders by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, orders)

}

// GetOrdersList godoc
// @Router       /orders [GET]
// @Summary      Get Orders list
// @Description  Get Orders list
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      200  {object}  models.OrdersResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetOrderssList(c *gin.Context) {

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

	response, err := h.storage.Orders().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

	if err != nil {
		handleResponse(c, "error while getting Orders ", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, response)

}

// UpdateOrders godoc
// @Router       /orders/{id} [PUT]
// @Summary      Update Orders by id
// @Description  Update Orders by id
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        id path string true "Orders id"
// @Param        Orders body models.UpdateOrders true "Orders"
// @Success      200  {object}  models.Orders
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateOrders(c *gin.Context) {
	updateOrders := models.UpdateOrders{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	//updateAuthor.ID = uid

	if err := c.ShouldBindJSON(&updateOrders); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Orders().Update(context.Background(), updateOrders)
	if err != nil {
		handleResponse(c, "error while updating Orders ", http.StatusInternalServerError, err.Error())
		return
	}

	orders, err := h.storage.Orders().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while getting Orders by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, orders)

}

// DeleteOrders godoc
// @Router       /orders/{id} [DELETE]
// @Summary      Delete Orders
// @Description  Delete Orders
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        id path string true "Orders id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteOrders(c *gin.Context) {

	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err := h.storage.Orders().Delete(context.Background(), id.String()); err != nil {
		handleResponse(c, "error while deleting Orders  by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data succesfully deleted")

}
