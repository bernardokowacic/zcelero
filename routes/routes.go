package routes

import (
	"zcelero/controller"
	"zcelero/service"

	"github.com/gin-gonic/gin"
)

func GetRoutes(router *gin.Engine, textManagementService service.TextManagementService) {
	router.POST("/v1/text-management", controller.Insert(textManagementService))
	router.GET("/v1/text-management", controller.Get(textManagementService))
}

func StartAPI(textManagementService service.TextManagementService) *gin.Engine {
	router := gin.Default()

	GetRoutes(router, textManagementService)
	return router
}
