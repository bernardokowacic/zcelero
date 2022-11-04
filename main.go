package main

import (
	"zcelero/repository"
	"zcelero/routes"
	"zcelero/service"
)

func main() {
	textManagementRepository := repository.NewRepository()
	textManagementService := service.NewService(textManagementRepository)

	router := routes.StartAPI(textManagementService)
	router.Run()
}
