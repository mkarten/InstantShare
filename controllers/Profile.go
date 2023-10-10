package controllers

import (
	"InstantShare/utils"
	"InstantShare/utils/db"
	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	if utils.SessionManager.CheckSession(c) {
		// get the events for the user
		user := utils.SessionManager.GetSessionUser(c)
		events, err := db.GetEventsByUserUUID(utils.SessionManager.GetSessionUser(c).UUID)
		if err != nil {
			c.HTML(400, "profile.html", gin.H{"error": "could not get events"})
			return
		}
		contribution, err := db.GetEventContributionByUserUUID(utils.SessionManager.GetSessionUser(c).UUID)

		if err != nil {
			c.HTML(400, "profile.html", gin.H{"error": "could not get events"})
			return
		}
		if len(contribution) == 0 {
			contribution = nil
		}
		if len(events) == 0 {
			events = nil
		}
		c.HTML(200, "profile.html", gin.H{"user": user, "events": events, "contributions": contribution})
	} else {
		c.Redirect(302, "/")
	}
}
