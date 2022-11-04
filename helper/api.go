package helper

import (
	"os"
	"zcelero/routes"
	"zcelero/service"

	"github.com/gin-gonic/gin"
)

// StartAPI initializes Gin API
func StartAPI(textManagementService service.TextManagementServiceInteface) *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.Default()

	routes.GetRoutes(router, textManagementService)
	return router
}
