package controllers

import "github.com/gin-gonic/gin"

func MainPage(c *gin.Context) {
	c.HTML(200, "home.html", gin.H{})
}

// create an event
// view events you have created
