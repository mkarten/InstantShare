package main

import (
	"InstantShare/controllers"
	"InstantShare/utils/db"
	"github.com/gin-gonic/gin"
	"github.com/tot0p/env"
)

func init() {
	gin.SetMode(gin.DebugMode)
	// TODO: need to load credentials from .env
	err := env.Load()
	if err != nil {
		panic(err)
	}
	db.InitDB(env.Get("HOST"), "3306", env.Get("USERNAME"), env.Get("PASSWORD"), env.Get("DATABASE"))
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
