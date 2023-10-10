package controllers

import (
	"InstantShare/utils"
	"InstantShare/utils/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"os"
	"path/filepath"
)

func UploadToCDN(c *gin.Context) {
	// get the multipart form in the request and save the image locally
	form, _ := c.MultipartForm()
	files := form.File["file"]
	// get the event_id
	eventID := c.PostForm("eventToken")
	for _, file := range files {
		// check the file extension
		if file.Header["Content-Type"][0] != "image/jpeg" && file.Header["Content-Type"][0] != "image/png" {
			c.String(400, "Invalid file type")
			return
		}
		//save the file extension
		fileExtension := filepath.Ext(file.Filename)
		// save the file in the proper directory
		fileName := uuid.New().String()
		// loop until we get a unique file name
		for {
			if _, err := os.Stat("./src/uploads/" + eventID + "/" + fileName + fileExtension); os.IsNotExist(err) {
				break
			}
			fileName = uuid.New().String()
		}
		err := c.SaveUploadedFile(file, "./src/uploads/"+eventID+"/"+fileName+fileExtension)
		if err != nil {
			c.String(500, "Internal server error")
			return
		}
		err = db.AddPictureToEvent("/"+eventID+"/"+fileName+fileExtension, utils.SessionManager.GetSessionUser(c).UUID, eventID)
		if err != nil {
			fmt.Println(err)
			// delete the file
			os.Remove("./src/uploads/" + eventID + "/" + fileName + fileExtension)
			c.String(500, "Internal server error")
			return
		}
	}
	c.Redirect(302, "/event/"+eventID)
}

func GetFromCDN(c *gin.Context) {
	// get the pictureUUID
	pictureUUID := c.Param("pictureUUID")
	// get the picture from the database
	picture, err := db.GetPictureByUUID(pictureUUID)
	if err != nil {
		c.String(404, "File not found")
		return
	}
	// check if the file exists
	if _, err := os.Stat("./src/uploads" + picture.FilePath); os.IsNotExist(err) {
		c.String(404, "File not found")
		return
	}
	// serve the file
	c.File("./src/uploads" + picture.FilePath)
}

func DownloadFromCDN(c *gin.Context) {
	// get the pictureUUID
	pictureUUID := c.Param("pictureUUID")
	// get the picture from the database
	picture, err := db.GetPictureByUUID(pictureUUID)
	if err != nil {
		c.String(404, "File not found")
		return
	}
	// check if the file exists
	if _, err := os.Stat("./src/uploads" + picture.FilePath); os.IsNotExist(err) {
		c.String(404, "File not found")
		return
	}
	// serve the file
	c.FileAttachment("./src/uploads"+picture.FilePath, picture.UUID+".jpg")
}
