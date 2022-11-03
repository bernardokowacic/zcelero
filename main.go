package main

import (
	"log"
	"zcelero/database"
	"zcelero/repository"
	"zcelero/routes"
	"zcelero/service"
)

func main() {
	dbConn, err := database.CreatePGConn()
	if err != nil {
		log.Fatal(err.Error())
	}

	database.Migrate(dbConn)

	textManagementRepository := repository.NewRepository(dbConn)
	textManagementService := service.NewService(textManagementRepository)

	router := routes.StartAPI(textManagementService)
	router.Run()
}
