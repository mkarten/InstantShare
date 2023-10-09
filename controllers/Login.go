package controllers

import (
	"InstantShare/utils"
	"InstantShare/utils/db"
	"InstantShare/utils/hash"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetLoginPage(c *gin.Context) {
	if utils.SessionManager.CheckSession(c) {
		c.Redirect(302, "/")
		return
	}
	c.HTML(200, "login.html", gin.H{})
}

func PostLoginPage(c *gin.Context) {
	// get the username and password
	username := c.PostForm("username")
	password := c.PostForm("password")

	// check if the username and password are correct
	user, err := db.GetUserByUsername(username)
	if err != nil {
		fmt.Println(err)
		c.HTML(400, "login.html", gin.H{
			"error": "we could not find a user with that username",
		})
		return
	}
	if !hash.CheckPasswordHash(password, user.Password) {
		c.HTML(400, "login.html", gin.H{
			"error": "incorrect password",
		})
		return
	}

	// set the session
	utils.SessionManager.CreateSession(c, &user)
	// redirect to the home page
	c.Redirect(302, "/")
}
