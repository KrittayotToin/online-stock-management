package main

import (
	"fmt"
	"log"
	"os"

	"github.com/KrittayotToin/online-stock-management/internal/app"
	"github.com/KrittayotToin/online-stock-management/internal/models"
	"github.com/KrittayotToin/online-stock-management/pkg/database"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get database credentials from environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Bangkok",
		host, user, password, dbName, port, sslmode,
	)

	// Initialize the database connection
	db, err := database.NewPostgresConnection(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connected successfully!")

	// Verify the connection with a simple query
	if err := db.Exec("SELECT 1").Error; err != nil {
		log.Fatalf("Failed to verify connection: %v", err)
	}

	// Auto-migrate models (Product and Stock)
	log.Println("Running database migration...")

	err = db.AutoMigrate(&models.Product{}, &models.Stock{})
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	log.Println("Database schema migrated successfully!")

	// Initialize and start the HTTP server
	server := app.NewServer(db) // Pass the database connection to the server
	if err := server.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
