package main

import (
	"InstantShare/controllers"
	"InstantShare/utils/db"
	"github.com/gin-gonic/gin"
	"github.com/tot0p/env"
)

func init() {
	gin.SetMode(gin.DebugMode)
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

	// profile
	r.GET("/profile", controllers.GetProfile)

	// register
	r.GET("/register", controllers.GetRegisterPage)
	r.POST("/register", controllers.PostRegisterPage)

	// login
	r.GET("/login", controllers.GetLoginPage)
	r.POST("/login", controllers.PostLoginPage)

	// logout
	r.GET("/logout", controllers.Logout)

	// create event
	r.GET("/createEvent", controllers.GetCreateEvent)
	r.POST("/createEvent", controllers.PostCreateEvent)

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{"path": c.Request.URL.Path})
	})

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
