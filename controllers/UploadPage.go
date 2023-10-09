package controllers

import (
	"github.com/gin-gonic/gin"
)

func UploadPage(c *gin.Context) {
	// render html file
	c.HTML(200, "upload.html", gin.H{})
}
