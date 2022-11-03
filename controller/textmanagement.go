package controller

import (
	"net/http"
	"zcelero/service"

	"github.com/gin-gonic/gin"
)

func Get(textManagementService service.TextManagementServiceInteface) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	}
}

func Insert(textManagementService service.TextManagementServiceInteface) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	}
}
