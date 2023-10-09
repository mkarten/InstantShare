package controllers

import (
	"github.com/gin-gonic/gin"
	"os"
)

func UploadToCDN(c *gin.Context) {
	// get the multipart form in the request and save the image locally
	form, _ := c.MultipartForm()
	files := form.File["file"]
	// get the event_id
	eventID := c.PostForm("event_id")
	for _, file := range files {
		// check the file extension
		if file.Header["Content-Type"][0] != "image/jpeg" && file.Header["Content-Type"][0] != "image/png" {
			c.String(400, "Bad request")
			return
		}
		// save the file in the proper directory
		err := c.SaveUploadedFile(file, "./src/uploads/"+eventID+"/"+file.Filename)
		if err != nil {
			c.String(500, "Internal server error")
			return
		}
	}
}

func GetFromCDN(c *gin.Context) {
	// get the file name
	fileName := c.Param("fileName")
	// check if the file exists
	if _, err := os.Stat("./src/uploads/" + fileName); os.IsNotExist(err) {
		c.String(404, "File not found")
		return
	}
	// return the file
	c.File("./src/uploads/" + fileName)
}
