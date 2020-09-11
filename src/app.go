package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gogin/src/internal/dir"
)

var router *gin.Engine

func main() {
	var htmlToLoad = dir.ListDir("templates")

	router = gin.Default()

	// router.LoadHTMLGlob("templates/*") -> don't working well with multilevel filestructure
	router.LoadHTMLFiles(htmlToLoad...)

	router.GET("/", func(c *gin.Context) {

		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "Home Page",
			},
		)

	})

	router.Run()

}
