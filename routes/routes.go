package routes

import (
	"github.com/gin-gonic/gin"
)

func Home(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the User API",
		})
	})
}

