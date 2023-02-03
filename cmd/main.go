package main

import (
	"github.com/martingarrettl/link-shortener/cmd/server"
	"github.com/martingarrettl/link-shortener/internal/controllers"
)

func main() {
	// initialize database connection
	// todo

	// create router
	r := server.NewServer()

	// Register routes
	r.GET("/:shortcode", controllers.ShortURL)
	r.POST("/", controllers.NewURL)

	// listen and serve on 0.0.0.0:8080
	r.Run()
}
