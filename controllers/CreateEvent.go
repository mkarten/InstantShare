package controllers

import (
	"InstantShare/utils"
	"github.com/gin-gonic/gin"
)

func GetCreateEvent(c *gin.Context) {
	if utils.SessionManager.CheckSession(c) {
		c.HTML(200, "createEvent.html", gin.H{"user": utils.SessionManager.GetSessionUser(c)})
	} else {
		c.HTML(200, "createEvent.html", gin.H{})
	}
}

func PostCreateEvent(c *gin.Context) {
	// TODO: implement event creation
}
