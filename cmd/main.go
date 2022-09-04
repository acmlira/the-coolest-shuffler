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
	cardsRepository := repository.NewCardsRepository(configs.GetPostgresDSN())
	decksRepository := repository.NewDecksRepository(
		configs.GetRedisUrl(),
		configs.GetRedisDatabase(),
		configs.GetRedisPassword())

	// Setup server
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true

	shufflerService := service.NewShufflerService(cardsRepository, decksRepository)

	// Register API routes
	deckAPI := api.NewDeckAPI(shufflerService)
	deckAPI.Register(server)

	// Start procedure
	log.Info("Starting the-coolest-shuffler in http://" + configs.GetAppUrl())
	server.Start(configs.GetAppUrl())
}
