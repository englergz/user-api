package main

import (
	"github.com/englergz/user-api/internal/controllers"
	"github.com/englergz/user-api/internal/database"
	"github.com/englergz/user-api/internal/models"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	// Auto-migrar la tabla User
	database.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	// Endpoints CRUD
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	// Iniciar servidor
	r.Run(":3000") // Escuchar en el puerto 3000
}
