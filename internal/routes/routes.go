package routes

import (
	"net/http"

	"github.com/KrittayotToin/online-stock-management/internal/handlers"
	"gorm.io/gorm"
)

// RegisterRoutes registers all application routes
func RegisterRoutes(mux *http.ServeMux, db *gorm.DB) {
	// Pass the database to the handlers
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		handlers.ProductHandler(w, r, db) // Pass database connection
	})
	mux.HandleFunc("/stocks", func(w http.ResponseWriter, r *http.Request) {
		handlers.StockHandler(w, r, db) // Pass database connection
	})
}
