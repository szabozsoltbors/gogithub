package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home - handle requests when the main accesspoint is invoked
func Home(c *gin.Context) {

	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Home Page",
		},
	)

	c.Next()

}
