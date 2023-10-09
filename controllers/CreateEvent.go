package controllers

import (
	"InstantShare/utils"
	"InstantShare/utils/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetCreateEvent(c *gin.Context) {
	if utils.SessionManager.CheckSession(c) {
		c.HTML(200, "createEvent.html", gin.H{"user": utils.SessionManager.GetSessionUser(c)})
	} else {
		c.Redirect(302, "/")
	}
}

func PostCreateEvent(c *gin.Context) {
	// get the session user
	user := utils.SessionManager.GetSessionUser(c)
	// get the event name
	eventName := c.PostForm("name")

	// create the event
	_, err := db.CreateEvent(eventName, user.UUID)
	if err != nil {
		fmt.Println(err)
		c.HTML(400, "createEvent.html", gin.H{
			"error": "could not create event",
		})
		return
	}
	c.Redirect(302, "/profile")
}
