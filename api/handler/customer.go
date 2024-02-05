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

// CreateCustomer godoc
// @Router       /customer [POST]
// @Summary      Create a new customer 
// @Description  Create a new customer 
// @Tags         customer
// @Accept       json
// @Produce      json
// @Param        customer body  models.CreateCustomer true  "customer data"
// @Success      201  {object}  models.Customer
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateCustomer(c *gin.Context) {
	createCustomer := models.CreateCustomer{}

	if err := c.ShouldBindJSON(&createCustomer); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
	}

	id, err := h.storage.Customer().Create(context.Background(), createCustomer)
	if err != nil {
		handleResponse(c, "error while creating customer", http.StatusInternalServerError, err)
		return
	}

	customer, err := h.storage.Customer().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while get customer ", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, customer)

}

// GetCustomerByID godoc
// @Router       /customer/{id} [GET]
// @Summary      Get customer by id
// @Description  Get customer by id
// @Tags         customer
// @Accept       json
// @Produce      json
// @Param        id path string true "customer"
// @Success      200  {object}  models.Customer
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetCustomerByID(c *gin.Context) {

	var err error

	uid := c.Param("id")

	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "invalid uuid type ", http.StatusBadRequest, err.Error())
		return
	}

	customer, err := h.storage.Customer().Get(context.Background(), models.PrimaryKey{
		ID: id.String(),
	})
	if err != nil {
		handleResponse(c, "error while get customer by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, customer)

}

// GetCustomersList godoc
// @Router       /customers [GET]
// @Summary      Get customers list
// @Description  Get customers list
// @Tags         customer
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      200  {object}  models.CustomersResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetCustomersList(c *gin.Context) {

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

	response, err := h.storage.Customer().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

	if err != nil {
		handleResponse(c, "error while getting customer", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, response)

}

// UpdateCustomer godoc
// @Router       /customer/{id} [PUT]
// @Summary      Update customer by id
// @Description  Update customer by id
// @Tags         customer
// @Accept       json
// @Produce      json
// @Param        id path string true "customer id"
// @Param        customer body models.UpdateCustomer true "customer"
// @Success      200  {object}  models.Customer
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateCustomer(c *gin.Context) {
	updateCustomer := models.UpdateCustomer{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	//updateAuthor.ID = uid

	if err := c.ShouldBindJSON(&updateCustomer); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Customer().Update(context.Background(), updateCustomer)
	if err != nil {
		handleResponse(c, "error while updating customer ", http.StatusInternalServerError, err.Error())
		return
	}

	customer, err := h.storage.Customer().Get(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(c, "error while getting customer by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, customer)

}

// DeleteCustomer godoc
// @Router       /customer/{id} [DELETE]
// @Summary      Delete customer 
// @Description  Delete customer
// @Tags         customer
// @Accept       json
// @Produce      json
// @Param        id path string true "customer id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteCustomer(c *gin.Context) {

	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err := h.storage.Customer().Delete(context.Background(), id.String()); err != nil {
		handleResponse(c, "error while deleting customer by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data succesfully deleted")

}
