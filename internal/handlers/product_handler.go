package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KrittayotToin/online-stock-management/internal/models"
	"gorm.io/gorm"
)

// ProductHandler handles product-related HTTP requests
func ProductHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Log the current database connection inside the handler
	var dbName string
	err := db.Raw("SELECT current_database();").Scan(&dbName).Error
	if err != nil {
		log.Printf("Failed to retrieve current database: %v", err)
	} else {
		log.Printf("Currently connected to database inside handler: %s", dbName)
	}

	switch r.Method {
	case http.MethodGet:
		// Fetch products from the database
		var products []models.Product
		if err := db.Find(&products).Error; err != nil {
			http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(products)
	case http.MethodPost:
		// Create a new product
		var product models.Product
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		if err := db.Create(&product).Error; err != nil {
			http.Error(w, "Failed to create product", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(product)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
