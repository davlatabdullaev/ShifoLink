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

// CreateOrderDrug godoc
// @Router       /order_drug [POST]
// @Summary      Create a new OrderDrug
// @Description  Create a new OrderDrug
// @Tags         order_drug
// @Accept       json
// @Produce      json
// @Param        orderDrug body  models.CreateOrderDrug true  "OrderDrug data"
// @Success      201  {object}  models.OrderDrug
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateOrderDrug(c *gin.Context) {
	createOrderDrug := models.CreateOrderDrug{}

	if err := c.ShouldBindJSON(&createOrderDrug); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
	}

	id, err := h.storage.OrderDrug().Create(context.Background(), createOrderDrug)
	if err != nil {
		handleResponse(c, "error while creating orderDrug ", http.StatusInternalServerError, err)
		return
	}

	orderDrug, err := h.storage.OrderDrug().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while get orderDrug ", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, orderDrug)

}

// GetOrderDrugByID godoc
// @Router       /order_drug/{id} [GET]
// @Summary      Get OrderDrug by id
// @Description  Get OrderDrug by id
// @Tags         order_drug
// @Accept       json
// @Produce      json
// @Param        id path string true "OrderDrug"
// @Success      200  {object}  models.OrderDrug
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetOrderDrugByID(c *gin.Context) {

	var err error

	uid := c.Param("id")

	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "invalid uuid type ", http.StatusBadRequest, err.Error())
		return
	}

	orderDrug, err := h.storage.OrderDrug().Get(context.Background(), models.PrimaryKey{
		ID: id.String(),
	})
	if err != nil {
		handleResponse(c, "error while get OrderDrug by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, orderDrug)

}

// GetOrderDrugsList godoc
// @Router       /order_drug [GET]
// @Summary      Get OrderDrugs list
// @Description  Get OrderDrugs list
// @Tags         order_drug
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      200  {object}  models.OrderDrugsResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetOrderDrugsList(c *gin.Context) {

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

	response, err := h.storage.OrderDrug().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

	if err != nil {
		handleResponse(c, "error while getting OrderDrug ", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, response)

}

// UpdateOrderDrug godoc
// @Router       /order_drug/{id} [PUT]
// @Summary      Update OrderDrug by id
// @Description  Update OrderDrug by id
// @Tags         order_drug
// @Accept       json
// @Produce      json
// @Param        id path string true "OrderDrug id"
// @Param        OrderDrug body models.UpdateOrderDrug true "OrderDrug"
// @Success      200  {object}  models.OrderDrug
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateOrderDrug(c *gin.Context) {
	updateOrderDrug := models.UpdateOrderDrug{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	//updateAuthor.ID = uid

	if err := c.ShouldBindJSON(&updateOrderDrug); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.OrderDrug().Update(context.Background(), updateOrderDrug)
	if err != nil {
		handleResponse(c, "error while updating OrderDrug ", http.StatusInternalServerError, err.Error())
		return
	}

	OrderDrug, err := h.storage.OrderDrug().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while getting OrderDrug by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, OrderDrug)

}

// DeleteOrderDrug godoc
// @Router       /order_drug/{id} [DELETE]
// @Summary      Delete OrderDrug
// @Description  Delete OrderDrug
// @Tags         order_drug
// @Accept       json
// @Produce      json
// @Param        id path string true "OrderDrug id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteOrderDrug(c *gin.Context) {

	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err := h.storage.OrderDrug().Delete(context.Background(), id.String()); err != nil {
		handleResponse(c, "error while deleting OrderDrug  by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data succesfully deleted")

}
