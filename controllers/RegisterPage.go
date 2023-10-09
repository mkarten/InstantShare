package controllers

import "github.com/gin-gonic/gin"

func GetRegisterPage(c *gin.Context) {
	c.HTML(200, "register.html", gin.H{})
}

func PostRegisterPage(c *gin.Context) {

}
