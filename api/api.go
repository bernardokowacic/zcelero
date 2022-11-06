package api

import (
	"os"
	"zcelero/routes"
	"zcelero/service"

	"github.com/gin-gonic/gin"
)

// Start initializes Gin API
func Start(textManagementService service.TextManagementServiceInteface) *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.Default()

	routes.GetRoutes(router, textManagementService)
	return router
}
