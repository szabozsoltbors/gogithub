package config

import (
	"github.com/gin-gonic/gin"

	"gogin/src/util"

	"gogin/src/api/routes"

	log "github.com/sirupsen/logrus"
)

// InitRoutes - initializes the API routes
func InitRoutes(app *gin.Engine) error {

	htmlFiles, err := util.ListDir("templates")

	if err != nil {
		log.WithError(err).Error("Unable to load the HTML templates")
		return err
	}

	app.LoadHTMLFiles(htmlFiles...)

	for _, file := range htmlFiles {
		log.WithFields(log.Fields{
			"name": file,
		}).Debug("HTML template imported")
	}

	routes.Home(app)

	return nil

}
