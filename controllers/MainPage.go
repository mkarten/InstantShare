package controllers

import "github.com/gin-gonic/gin"

func MainPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// create an event
// view events you have created
