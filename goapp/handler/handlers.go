package handler

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(c *gin.Engine) {
	//Default Route
	c.NoRoute(func(c *gin.Context) {
		c.JSON(405, gin.H{
			"errorCode": 405,
			"Message":   "Only POST HTTP methods are supported.",
		})
		c.Abort()
	})
	c.POST("/finnomena/suggest-funds", suggestFund())
}
