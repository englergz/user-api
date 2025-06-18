package database

import (
	"fmt"
	"log"
	

	"github.com/englergz/user-api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=secret dbname=users_db port=5432"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error al conectar a la base de datos: ", err)
	}

	DB.AutoMigrate(&models.User{})
	fmt.Println("Base de datos conectada")
}
