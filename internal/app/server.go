package app

import (
	"log"
	"net/http"

	"github.com/KrittayotToin/online-stock-management/internal/routes"
	"gorm.io/gorm"
)

// Server holds the HTTP server and database connection
type Server struct {
	httpServer *http.Server
	db         *gorm.DB
}

// NewServer creates a new server instance with a database connection
func NewServer(db *gorm.DB) *Server {
	mux := http.NewServeMux()

	// Pass the database connection to routes
	routes.RegisterRoutes(mux, db)

	return &Server{
		httpServer: &http.Server{
			Addr:    ":8080", // Port for the application
			Handler: mux,
		},
		db: db, // Store the database connection
	}
}

// Run starts the HTTP server
func (s *Server) Run() error {
	log.Println("Server is running on port 8080")
	return s.httpServer.ListenAndServe()
}
