package main

import (
	"InstantShare/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// add the routes
	r.GET("/", controllers.MainPage)

	// create api group
	api := r.Group("/api")
	api.GET("/ping", controllers.Ping)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

//InstantShare
//Gvjk2ABYaLa8Pbg2
