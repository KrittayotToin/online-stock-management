package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KrittayotToin/online-stock-management/internal/models"
	"gorm.io/gorm"
)

// StockHandler handles stock-related HTTP requests
func StockHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Log the current database connection
	var dbName string
	err := db.Raw("SELECT current_database();").Scan(&dbName).Error
	if err != nil {
		log.Printf("Failed to retrieve current database: %v", err)
	} else {
		log.Printf("Currently connected to database: %s", dbName)
	}

	switch r.Method {
	case http.MethodGet:
		// Fetch stocks from the database
		var stocks []models.Stock
		if err := db.Find(&stocks).Error; err != nil {
			http.Error(w, "Failed to fetch stocks", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(stocks)
	case http.MethodPost:
		// Create a new stock
		var stock models.Stock
		if err := json.NewDecoder(r.Body).Decode(&stock); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		if err := db.Create(&stock).Error; err != nil {
			http.Error(w, "Failed to create stock", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(stock)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
