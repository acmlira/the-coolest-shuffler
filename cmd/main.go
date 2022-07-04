package main

import (
	"the-coolest-shuffler/configs"
	"the-coolest-shuffler/internal/api"
	"the-coolest-shuffler/internal/repository"
	"the-coolest-shuffler/internal/service"

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

	// Register API routes
	api.Deck{Shuffler: service.NewShuffler()}.Register(server)

	// Start procedure
	log.Info("Starting the-coolest-shuffler in http://" + configs.GetAppUrl())
	server.Start(configs.GetAppUrl())
}