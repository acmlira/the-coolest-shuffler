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
	database := repository.NewDatabase(configs.GetPostgresDSN())
	cache := repository.NewCache(
		configs.GetRedisUrl(),
		configs.GetRedisDatabase(),
		configs.GetRedisPassword())

	// Setup server
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true

	shuffler := service.NewShuffler(cache, database)

	// Register API routes
	decks := api.NewDecks(shuffler)
	decks.Register(server)

	// Start procedure
	log.Info("Starting the-coolest-shuffler in http://" + configs.GetAppUrl())
	server.Start(configs.GetAppUrl())
}
