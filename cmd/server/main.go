package main

import (
	"log"

	"github.com/Anamika1608/re-memorise/internal/api"
	"github.com/Anamika1608/re-memorise/internal/config"
	"github.com/Anamika1608/re-memorise/internal/db"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Connect to database
	database, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate database models
	if err := db.Migrate(database); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Setup API server
	server := api.NewServer(cfg, database)

	// Start server
	log.Printf("Starting server on %s", cfg.ServerAddress)
	if err := server.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}