package main

import (
	"log"
	"zcelero/database"
	"zcelero/repository"
	"zcelero/routes"
	"zcelero/service"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

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
