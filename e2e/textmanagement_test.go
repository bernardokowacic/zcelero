package e2e_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"zcelero/api"
	"zcelero/entity"
	"zcelero/helper"
	"zcelero/repository"
	"zcelero/service"

	"github.com/go-playground/assert/v2"
)

func TestEndToEndEncrypted(t *testing.T) {
	os.Mkdir("storage", 0777)
	helper := helper.NewHelper()
	textManagementRepository := repository.NewRepository(helper)
	textManagementService := service.NewService(textManagementRepository, helper)
	router := api.Start(textManagementService)
	encryptation := true

	postArgs := entity.TextManagement{
		TextData:           "encrypted text data",
		Encryption:         &encryptation,
		KeySize:            1024,
		PrivateKeyPassword: "123456",
	}
	postBody, _ := json.Marshal(postArgs)

	req, _ := http.NewRequest(http.MethodPost, "/v1/text-management", bytes.NewReader(postBody))
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	fmt.Println(res.Body)

	postResponse := struct {
		PrivateKey string `json:"private_key"`
		Uuid       string `json:"uuid"`
	}{}
	json.Unmarshal(res.Body.Bytes(), &postResponse)

	getArgs := entity.TextManagement{
		PrivateKey:         postResponse.PrivateKey,
		PrivateKeyPassword: postArgs.PrivateKeyPassword,
	}
	getBody, _ := json.Marshal(getArgs)

	req, _ = http.NewRequest(http.MethodGet, "/v1/text-management?text_id="+postResponse.Uuid, bytes.NewReader(getBody))
	res = httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, `{"text":"encrypted text data"}`, res.Body.String())
}

func TestEndToEndNonEncrypted(t *testing.T) {
	os.Mkdir("storage", 0777)
	helper := helper.NewHelper()
	textManagementRepository := repository.NewRepository(helper)
	textManagementService := service.NewService(textManagementRepository, helper)
	router := api.Start(textManagementService)
	encryptation := false

	postArgs := entity.TextManagement{
		TextData:   "text data",
		Encryption: &encryptation,
	}
	postBody, _ := json.Marshal(postArgs)

	req, _ := http.NewRequest(http.MethodPost, "/v1/text-management", bytes.NewReader(postBody))
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	fmt.Println(res.Body)

	postResponse := struct {
		PrivateKey string `json:"private_key"`
		Uuid       string `json:"uuid"`
	}{}
	json.Unmarshal(res.Body.Bytes(), &postResponse)

	getArgs := entity.TextManagement{
		PrivateKey:         postResponse.PrivateKey,
		PrivateKeyPassword: postArgs.PrivateKeyPassword,
	}
	getBody, _ := json.Marshal(getArgs)

	req, _ = http.NewRequest(http.MethodGet, "/v1/text-management?text_id="+postResponse.Uuid, bytes.NewReader(getBody))
	res = httptest.NewRecorder()
	router.ServeHTTP(res, req)
	fmt.Println(res.Body)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, `{"text":"text data"}`, res.Body.String())
}
