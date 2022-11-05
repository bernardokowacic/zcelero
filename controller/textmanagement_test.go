package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"zcelero/entity"
	"zcelero/helper"
	serviceMock "zcelero/mocks/service"

	"github.com/go-playground/assert/v2"
)

func TestGetUserRoute(t *testing.T) {
	service := &serviceMock.TextManagementServiceInteface{}
	router := helper.StartAPI(service)

	uuid := "154ad8a0-1e42-4cf6-9d7b-e49f71dcc4ec"
	args := struct {
		PrivateKey         string `json:"private_key"`
		PrivateKeyPassword string `json:"private_key_password"`
	}{
		PrivateKey:         "private_key",
		PrivateKeyPassword: "aaa",
	}
	body, _ := json.Marshal(args)

	service.On("Get", uuid, args.PrivateKey, args.PrivateKeyPassword).Return("message", nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/text-management?text_id="+uuid, bytes.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetUserRouteWithoutTextID(t *testing.T) {
	service := &serviceMock.TextManagementServiceInteface{}
	router := helper.StartAPI(service)

	args := struct {
		PrivateKey         string `json:"private_key"`
		PrivateKeyPassword string `json:"private_key_password"`
	}{
		PrivateKey:         "private_key",
		PrivateKeyPassword: "aaa",
	}
	body, _ := json.Marshal(args)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/text-management", bytes.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotAcceptable, w.Code)
}

func TestGetUserRouteWithServiceError(t *testing.T) {
	service := &serviceMock.TextManagementServiceInteface{}
	router := helper.StartAPI(service)

	uuid := "154ad8a0-1e42-4cf6-9d7b-e49f71dcc4ec"
	args := struct {
		PrivateKey         string `json:"private_key"`
		PrivateKeyPassword string `json:"private_key_password"`
	}{
		PrivateKey:         "private_key",
		PrivateKeyPassword: "aaa",
	}
	body, _ := json.Marshal(args)

	service.On("Get", uuid, args.PrivateKey, args.PrivateKeyPassword).Return("", errors.New("some error"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/text-management?text_id="+uuid, bytes.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetUserRouteBidingError(t *testing.T) {
	service := &serviceMock.TextManagementServiceInteface{}
	router := helper.StartAPI(service)

	uuid := "154ad8a0-1e42-4cf6-9d7b-e49f71dcc4ec"
	args := struct {
		PrivateKey         string `json:"private_key"`
		PrivateKeyPassword int    `json:"private_key_password"`
	}{
		PrivateKey:         "private_key",
		PrivateKeyPassword: 12,
	}
	body, _ := json.Marshal(args)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/text-management?text_id="+uuid, bytes.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotAcceptable, w.Code)
}

func TestPostUserRouteWithoutEncryptation(t *testing.T) {
	service := &serviceMock.TextManagementServiceInteface{}
	router := helper.StartAPI(service)

	encryptation := false
	args := entity.TextManagement{
		TextData:   "text data",
		Encryption: &encryptation,
	}
	body, _ := json.Marshal(args)

	response := entity.TextManagement{
		TextData:   args.TextData,
		Encryption: args.Encryption,
		Uuid:       "uuid",
	}

	service.On("Insert", args).Return(response, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/text-management", bytes.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPostUserRouteWithEncryptation(t *testing.T) {
	service := &serviceMock.TextManagementServiceInteface{}
	router := helper.StartAPI(service)

	encryptation := true
	args := entity.TextManagement{
		TextData:           "text data",
		Encryption:         &encryptation,
		KeySize:            1024,
		PrivateKeyPassword: "aaa",
	}
	body, _ := json.Marshal(args)

	response := entity.TextManagement{
		TextData:           args.TextData,
		Encryption:         args.Encryption,
		KeySize:            args.KeySize,
		PrivateKeyPassword: args.PrivateKeyPassword,
		Uuid:               "uuid",
		PrivateKey:         "private_key",
	}

	service.On("Insert", args).Return(response, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/text-management", bytes.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPostUserRouteWithBindingError(t *testing.T) {
	service := &serviceMock.TextManagementServiceInteface{}
	router := helper.StartAPI(service)

	encryptation := true
	args := entity.TextManagement{
		Encryption:         &encryptation,
		KeySize:            1024,
		PrivateKeyPassword: "aaa",
	}
	body, _ := json.Marshal(args)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/text-management", bytes.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPostUserRouteWithoutPassword(t *testing.T) {
	service := &serviceMock.TextManagementServiceInteface{}
	router := helper.StartAPI(service)

	encryptation := true
	args := entity.TextManagement{
		TextData:   "text data",
		Encryption: &encryptation,
		KeySize:    1024,
	}
	body, _ := json.Marshal(args)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/text-management", bytes.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotAcceptable, w.Code)
}

func TestPostUserRouteWithWrongKeySize(t *testing.T) {
	service := &serviceMock.TextManagementServiceInteface{}
	router := helper.StartAPI(service)

	encryptation := true
	args := entity.TextManagement{
		TextData:           "text data",
		Encryption:         &encryptation,
		PrivateKeyPassword: "aaa",
		KeySize:            12,
	}
	body, _ := json.Marshal(args)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/text-management", bytes.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotAcceptable, w.Code)
}

func TestPostUserRouteWithInsertError(t *testing.T) {
	service := &serviceMock.TextManagementServiceInteface{}
	router := helper.StartAPI(service)

	encryptation := true
	args := entity.TextManagement{
		TextData:           "text data",
		Encryption:         &encryptation,
		KeySize:            1024,
		PrivateKeyPassword: "aaa",
	}
	body, _ := json.Marshal(args)

	service.On("Insert", args).Return(entity.TextManagement{}, errors.New("error"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/text-management", bytes.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
