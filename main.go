package main

import (

	"github.com/englergz/user-api/internal/database"
	"github.com/englergz/user-api/internal/models"
	"github.com/englergz/user-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	// Auto-migrar la tabla User
	database.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	// Endpoints CRUD
	// Registrar rutas de usuario
	routes.RegisterUserRoutes(r)

	// Iniciar servidor
	r.Run(":3000") // Escuchar en el puerto 3000
}
