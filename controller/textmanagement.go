package controllers

import (
	"net/http"
	"zcelero/service"

	"github.com/gin-gonic/gin"
)

func GetUser(textManagementService service.TextManagementService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	}
}
