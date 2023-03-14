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

	"github.com/gorilla/handlers"
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

	// Enable CORS using the handlers package: Local Development
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	// // Start server
	port := cfg.Server.Port
	fmt.Printf("Server listening on port %d...\n", port)

	err = http.ListenAndServe(":8080", handlers.CORS(allowedOrigins, allowedMethods)(router))
	if err != nil {
		panic(err)
	}

}
