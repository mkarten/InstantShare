package controllers

import (
	"InstantShare/utils"
	"InstantShare/utils/db"
	"github.com/gin-gonic/gin"
)

func LikePicture(c *gin.Context) {
	// check if the user is logged in
	if utils.SessionManager.CheckSession(c) {
		// get the picture token from the url
		pictureUUID := c.Param("pictureUUID")
		// get the session user
		user := utils.SessionManager.GetSessionUser(c)
		// get the picture from the database
		picture, err := db.GetPictureByUUID(pictureUUID)
		if err != nil {
			c.HTML(400, "event.html", gin.H{"error": "could not get picture"})
			return
		}
		err = db.LikePicture(picture.UUID, user.UUID)
		if err != nil {
			c.HTML(400, "event.html", gin.H{"error": "could not like picture"})
			return
		}
		c.Redirect(302, "/event/"+picture.EventToken)
	} else {
		c.Redirect(302, "/login")
	}

}
