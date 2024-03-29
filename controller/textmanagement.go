package controller

import (
	"net/http"
	"zcelero/entity"
	"zcelero/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Get(textManagementService service.TextManagementServiceInteface) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Debug().Msg("end-point GET /v1/text-management requested")

		textId, exists := c.GetQuery("id")
		if !exists {
			log.Info().Msg("id not sent")
			c.JSON(http.StatusNotAcceptable, gin.H{"message": "id param is required"})
			return
		}

		json := struct {
			PrivateKey         string `json:"private_key"`
			PrivateKeyPassword string `json:"private_key_password"`
		}{}
		if err := c.ShouldBindJSON(&json); err != nil {
			log.Error().Msg(err.Error())
			c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
			return
		}

		response, err := textManagementService.Get(textId, json.PrivateKey, json.PrivateKeyPassword)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Debug().Msg("end-point GET /v1/text-management finished")

		c.JSON(http.StatusOK, gin.H{"text": response})
	}
}

func Insert(textManagementService service.TextManagementServiceInteface) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Debug().Msg("end-point POST /v1/text-management requested")

		var json entity.TextManagement
		if err := c.ShouldBindJSON(&json); err != nil {
			log.Error().Msg(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if *json.Encryption && json.PrivateKeyPassword == "" {
			log.Info().Msg("private_key_password not sent when encryption is required")
			c.JSON(http.StatusNotAcceptable, gin.H{"error": "private_key_password is required when encryption is true"})
			return
		}
		if *json.Encryption && (json.KeySize != 1024 && json.KeySize != 2048 && json.KeySize != 4096) {
			log.Info().Msg("keysize must be 1024 or 2048 or 4096")
			c.JSON(http.StatusNotAcceptable, gin.H{"error": "keysize must be 1024 or 2048 or 4096"})
			return
		}

		response, err := textManagementService.Insert(json)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Debug().Msg("end-point GET /v1/text-management finished")

		c.JSON(http.StatusOK, gin.H{"uuid": response.Uuid, "private_key": response.PrivateKey})
	}
}
