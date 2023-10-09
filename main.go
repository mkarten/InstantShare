package main

import (
	"InstantShare/controllers"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.DebugMode)
	// TODO: need to load credentials from .env
	//db.InitDB("", "", "", "", "") // "host","port","user","password","database
}

func main() {
	r := gin.Default()

	// set the html template
	r.LoadHTMLGlob("./src/templates/*")

	// add the routes
	r.GET("/", controllers.MainPage)

	// register
	r.GET("/register", controllers.GetRegisterPage)
	r.POST("/register", controllers.PostRegisterPage)

	// login
	r.GET("/login", controllers.GetLoginPage)
	r.POST("/login", controllers.PostLoginPage)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
