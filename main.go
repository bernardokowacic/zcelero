package main

import (
	"fmt"
	"os"
	"zcelero/repository"
	"zcelero/routes"
	"zcelero/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	zerolog.SetGlobalLevel(zerologLogLevel(os.Getenv("INFO_LEVEL")))
}

func main() {
	textManagementRepository := repository.NewRepository()
	textManagementService := service.NewService(textManagementRepository)

	router := StartAPI(textManagementService)
	log.Info().Msg("API Started")
	router.Run()
}

func StartAPI(textManagementService service.TextManagementServiceInteface) *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.Default()

	routes.GetRoutes(router, textManagementService)
	return router
}

func zerologLogLevel(level string) zerolog.Level {
	switch level {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "error":
		return zerolog.ErrorLevel
	case "warn":
		return zerolog.WarnLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.PanicLevel
	default:
		panic(fmt.Sprintf("the specified %s log level is not supported", level))
	}
}
