package main

import (
	"the-coolest-shuffler/configs"

	log "github.com/sirupsen/logrus"
)

func main()  {
	// Initialize configurations
	configs.Init()

	// Start procedure
	log.Info("Starting the-coolest-shuffler in http://" + configs.GetAppUrl())
}