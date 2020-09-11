package routes

import (
	"github.com/gin-gonic/gin"

	"gogin/src/api/controller"
)

// Home - entry points to the home page
func Home(app *gin.Engine) {
	app.GET("/", controller.Home)
	app.GET("/index", controller.Home)
}
