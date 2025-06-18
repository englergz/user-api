package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/englergz/user-api/internal/controllers"
)

func RegisterUserRoutes(router *gin.Engine) {

	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.GET("/:id", controllers.GetUser)
		userRoutes.POST("/", controllers.CreateUser)
		userRoutes.PUT("/:id", controllers.UpdateUser)
		userRoutes.DELETE("/:id", controllers.DeleteUser)
	}
}