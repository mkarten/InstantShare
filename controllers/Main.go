package controllers

import (
	"InstantShare/utils"
	"github.com/gin-gonic/gin"
)

func MainPage(c *gin.Context) {
	if utils.SessionManager.CheckSession(c) {
		c.HTML(200, "home.html", gin.H{"user": utils.SessionManager.GetSessionUser(c)})
	} else {
		c.HTML(200, "home.html", gin.H{})
	}
}

// create an event
// view events you have created
