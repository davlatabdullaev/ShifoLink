package handler

import (
	"shifolink/api/models"
	"shifolink/service"
	"shifolink/storage"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	storage storage.IStorage
	services service.IServiceManager
}

func New(store storage.IStorage, services service.IServiceManager) Handler {
	return Handler{
		storage: store,
		services: services,
	}
}

func handleResponse(c *gin.Context, msg string, statusCode int, data interface{}) {
	response := models.Response{}

	switch code := statusCode; {
	case code < 400:
		response.Description = "succes"
	case code < 500:
		response.Description = "bad request"
	default:
		response.Description = "internal server error"

	}

	response.StatusCode = statusCode
	response.Data = data

	c.JSON(response.StatusCode, response)

}
