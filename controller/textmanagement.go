package controller

import (
	"net/http"
	"zcelero/entity"
	"zcelero/service"

	"github.com/gin-gonic/gin"
)

func Get(textManagementService service.TextManagementServiceInteface) gin.HandlerFunc {
	return func(c *gin.Context) {
		textId, exists := c.GetQuery("text_id")
		if !exists {
			c.JSON(http.StatusNotAcceptable, gin.H{"message": "text_id param is required"})
			return
		}

		json := struct {
			PrivateKey         string `json:"private_key"`
			PrivateKeyPassword string `json:"private_key_password"`
		}{}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response, err := textManagementService.Get(textId, json.PrivateKey, json.PrivateKeyPassword)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"text": response})
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
