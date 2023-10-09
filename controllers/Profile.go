package controllers

import (
	"InstantShare/utils"
	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	if utils.SessionManager.CheckSession(c) {
		c.HTML(200, "profile.html", gin.H{"user": utils.SessionManager.GetSessionUser(c)})
	} else {
		c.HTML(200, "profile.html", gin.H{})
	}
}
