package controller

import (
	"net/http"
	"zcelero/database/entity"
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
		var json entity.TextManagement
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if *json.Encryption && json.PrivateKeyPassword == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "private_key_password is required when encryption is true"})
			return
		}

		response, err := textManagementService.Insert(json)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"uuid": response.Uuid, "private_key": response.PrivateKey})
	}
}
