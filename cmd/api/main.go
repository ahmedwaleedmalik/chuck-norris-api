package main

import (
	"log"
	"net/http"

	"github.com/ahmedwaleedmalik/chuck-norris-api/internal/services/banter"
	"github.com/ahmedwaleedmalik/chuck-norris-api/internal/services/database"
	"github.com/ahmedwaleedmalik/chuck-norris-api/internal/services/health"
)

const defaultServingPort = "8080"

func main() {

	// Initialize the database
	db, err := database.InitializeDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// Close connection to the database
	defer db.Close()

	// Create new Banter Service Instance
	banterService := banter.NewBanterService(db)

	// Register Endpoints for Banter Service
	banterService.RegisterBanterServiceEndpoints()

	// Create new Health Service Instance
	healthService := health.NewHealthService(db)

	// Register Endpoints for Health Service
	healthService.RegisterHealthServiceEndpoints()

	// Start serving the http server
	log.Printf("HTTP server listening on %v", defaultServingPort)
	log.Fatal(http.ListenAndServe(":"+defaultServingPort, nil))
}
