package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logrus.Debug("Starting API service...")

	router := gin.Default()

	// Register APIs
	logrus.Debug("Registering Welcome API...")
	RegisterWelcomeAPI(router)
	logrus.Debug("Registering Fruits API...")
	RegisterFruitsAPI(router)

	logrus.Debug("Service running at http://localhost:8080")
	router.Run(":8080")
}
