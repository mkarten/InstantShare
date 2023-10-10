package controllers

import (
	"InstantShare/models"
	"InstantShare/utils"
	"InstantShare/utils/db"
	"github.com/gin-gonic/gin"
	"os"
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
			c.HTML(400, "event.html", gin.H{"error": "could not get pictures"})
			return
		}
		if len(pictures) == 0 {
			c.HTML(200, "event.html", gin.H{"event": event, "user": user, "link": c.Request.Host + "/event/" + eventToken})
			return
		}
		if len(pictures) == 1 {
			highlightedPicture := pictures[:1]
			pictures = pictures[1:]
			highlighted := struct {
				First  models.Picture
				Second bool
				Third  bool
			}{
				First:  highlightedPicture[0],
				Second: false,
				Third:  false,
			}
			c.HTML(200, "event.html", gin.H{"event": event, "user": user, "pictures": pictures, "highlightedPictures": highlighted, "link": c.Request.Host + "/event/" + eventToken})
			return
		}
		if len(pictures) == 2 {
			highlightedPicture := pictures[:2]
			pictures = pictures[2:]
			highlighted := struct {
				First  models.Picture
				Second models.Picture
				Third  bool
			}{
				First:  highlightedPicture[0],
				Second: highlightedPicture[1],
				Third:  false,
			}
			c.HTML(200, "event.html", gin.H{"event": event, "user": user, "pictures": pictures, "highlightedPictures": highlighted, "link": c.Request.Host + "/event/" + eventToken})
			return
		}
		if len(pictures) >= 3 {
			highlightedPicture := pictures[:3]
			pictures = pictures[3:]
			highlighted := struct {
				First  models.Picture
				Second models.Picture
				Third  models.Picture
			}{
				First:  highlightedPicture[0],
				Second: highlightedPicture[1],
				Third:  highlightedPicture[2],
			}
			c.HTML(200, "event.html", gin.H{"event": event, "user": user, "pictures": pictures, "highlightedPictures": highlighted, "link": c.Request.Host + "/event/" + eventToken})
			return
		}
	} else {
		c.Redirect(302, "/")
	}
}

func DeleteEvent(c *gin.Context) {
	if utils.SessionManager.CheckSession(c) {
		// get the events token from the form
		eventToken := c.Param("eventToken")
		// check if the user is the owner of the event
		event, err := db.GetEventByToken(eventToken)
		if err != nil {
			c.Redirect(302, "/profile")
			return
		}
		if event.UserUUID != utils.SessionManager.GetSessionUser(c).UUID {
			c.Redirect(302, "/profile")
			return
		}
		// delete the event from the database
		err = db.DeleteEventByToken(eventToken)
		if err != nil {
			c.Redirect(302, "/profile")
			return
		}
		//delete the pictures from the database
		err = db.DeleteEventPictures(eventToken)
		if err != nil {
			c.Redirect(302, "/profile")
			return
		}
		// delete the event folder
		os.RemoveAll("./src/uploads/" + eventToken)
		// redirect to the profile page
		c.Redirect(302, "/profile")
	} else {
		c.Redirect(302, "/")
	}
}

// view all photos of an event
// upload a photo to an event
