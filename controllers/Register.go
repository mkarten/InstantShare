package controllers

import (
	"InstantShare/utils"
	"InstantShare/utils/db"
	"github.com/gin-gonic/gin"
)

func GetRegisterPage(c *gin.Context) {
	if utils.SessionManager.CheckSession(c) {
		c.Redirect(302, "/")
		return
	}
	c.HTML(200, "register.html", gin.H{})
}

func PostRegisterPage(c *gin.Context) {
	// get the username, password, first name, and last name
	username := c.PostForm("username")
	password := c.PostForm("password")
	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")

	// check if the username is not empty
	if username == "" {
		c.HTML(400, "register.html", gin.H{
			"error": "username cannot be empty",
		})
		return
	}
	// create the user
	user, err := db.CreateUser(username, password, firstName, lastName)
	if err != nil {
		c.HTML(400, "register.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	// set the session
	utils.SessionManager.CreateSession(c, &user)
	// redirect to the home page
	c.Redirect(302, "/")
}
