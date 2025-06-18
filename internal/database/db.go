package database

import (
	"fmt"
	"log"
	"os"
	"time"


	"github.com/englergz/user-api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = "localhost" // valor por defecto en desarrollo local
	}

	dsn := fmt.Sprintf("host=%s user=postgres password=secret dbname=users_db port=5432 sslmode=disable", host)

	var err error
	maxAttempts := 10

	for i := 1; i <= maxAttempts; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Intento %d: No se pudo conectar a la base de datos, reintentando en 3s...", i)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatalf("Error al conectar a la base de datos después de %d intentos: %v", maxAttempts, err)
	}

	DB.AutoMigrate(&models.User{})
	fmt.Println("Base de datos conectada")
}
