package controllers

import (
	"InstantShare/utils"
	"InstantShare/utils/db"
	"github.com/gin-gonic/gin"
)

func GetEvent(c *gin.Context) {
	if utils.SessionManager.CheckSession(c) {
		// get the events token from the url
		eventToken := c.Param("eventToken")
		// get the event from the database
		event, err := db.GetEventByToken(eventToken)
		if err != nil {
			c.HTML(400, "event.html", gin.H{"error": "could not get event"})
			return
		}
		// get the session user
		user := utils.SessionManager.GetSessionUser(c)
		pictures, err := db.GetPicturesByEventToken(eventToken)
		if err != nil {
			c.HTML(400, "event.html", gin.H{"event": event, "user": user, "error": "could not get pictures"})
			return
		}

		c.HTML(200, "event.html", gin.H{"event": event, "user": user, "pictures": pictures})
	} else {
		c.Redirect(302, "/")
	}
}

// view all photos of an event
// upload a photo to an event
