package controllers

import "github.com/gin-gonic/gin"

func GetLoginPage(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func PostLoginPage(c *gin.Context) {

}