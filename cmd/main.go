package main

import (
	"the-coolest-shuffler/configs"
	docs "the-coolest-shuffler/docs"
	"the-coolest-shuffler/internal/api"
	"the-coolest-shuffler/internal/middleware"
	"the-coolest-shuffler/internal/repository"
	"the-coolest-shuffler/internal/service"

	"github.com/labstack/echo/v4"
)

// @title The Coolest Shuffler
// @version 1.0
// @description API to handle the deck and cards to be used in any game like Poker or Blackjack

// @contact.name acmlira
// @contact.url https://github.com/acmlira/the-coolest-shuffler

// @license.name MIT
// @license.url https://www.mit.edu/~amini/LICENSE.md

// @host http://localhost:8916
// @BasePath /the-coolest-shuffler/v1
func main() {
	// Initialize configurations
	configs.Init()

	// Init docs
	docs.SwaggerInfo.Host = configs.GetAppUrl()
	docs.SwaggerInfo.Version = configs.GetAppVersion()

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

	// Middlewares
	server.Use(middleware.Logger)

	// Start procedure
	server.Start(configs.GetAppUrl())
}
