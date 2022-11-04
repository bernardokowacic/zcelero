package main

import (
	"fmt"
	"os"
	"zcelero/helper"
	"zcelero/repository"
	"zcelero/service"

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

	router := helper.StartAPI(textManagementService)
	log.Info().Msg("API Started")
	router.Run()
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
