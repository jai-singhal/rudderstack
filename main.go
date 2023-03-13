package main

import (
	"fmt"
	"log"
	"net/http"

	// "github.com/jai-singhal/rudderstack/api"
	// "github.com/jai-singhal/rudderstack/config"
	// "github.com/jai-singhal/rudderstack/db"

	"rudderstack/internal/api/v1"
	"rudderstack/internal/config"
	"rudderstack/internal/db"
)

func main() {
	// Load app configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Connect to the database
	dbConn, err := db.NewDatabaseConnection(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer dbConn.Close()

	// Load API routes
	router := api.IntializeRouter(dbConn.Connection)

	// // Start server
	port := cfg.Server.Port
	fmt.Printf("Server listening on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
