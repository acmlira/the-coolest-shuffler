package main

import (
	"the-coolest-shuffler/configs"
	"the-coolest-shuffler/internal/repository"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Initialize configurations
	configs.Init()

	// Initialize repositories
	repository.Database(configs.GetPostgresDSN())
	repository.Cache(
		configs.GetRedisHost(),
		configs.GetRedisPort(),
		configs.GetRedisDatabase(),
		configs.GetRedisPassword())

	// Setup server
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true

	// Start procedure
	log.Info("Starting the-coolest-shuffler in http://" + configs.GetAppUrl())
	server.Start(configs.GetAppUrl())
}