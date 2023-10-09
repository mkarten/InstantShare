package controllers

import (
	"InstantShare/utils"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	// destroy the session
	utils.SessionManager.DeleteSession(c)
	// redirect to the home page
	c.Redirect(302, "/")
}
