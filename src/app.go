package main

import (
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"gogin/src/api/config"
)

func main() {
	app := gin.Default()

	if config.InitEnv() != nil ||
		config.InitLogger() != nil ||
		config.InitRoutes(app) != nil {
		return
	}

	log.WithFields(log.Fields{
		"profile":       config.ENV.PROFILE,
		"port":          config.ENV.PORT,
		"logging level": config.ENV.LOGLEVEL,
	}).Info("The application is configured:")

	app.Run(config.ENV.PORT)
}
