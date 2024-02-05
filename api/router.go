package api

import (
	"shifolink/storage"

	_ "shifolink/api/docs"
	"shifolink/api/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func New(store storage.IStorage) *gin.Engine {

	handler.New(store)

	r := gin.New()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r

}
