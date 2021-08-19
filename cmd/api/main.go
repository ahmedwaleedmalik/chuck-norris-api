package main

import (
	"log"
	"net/http"

	"github.com/ahmedwaleedmalik/chuck-norris-api/internal/services/banter"
	"github.com/ahmedwaleedmalik/chuck-norris-api/internal/services/database"
)

const defaultServingPort = "8080"

func main() {

	// Initialize the database
	db, err := database.InitializeDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// Create new Banter Service Instance
	configService := banter.NewBanterService(db)

	// Register Endpoints for Banter Service
	configService.RegisterBanterServiceEndpoints()

	// Start serving the http server
	log.Printf("HTTP server listening on %v", defaultServingPort)
	log.Fatal(http.ListenAndServe(":"+defaultServingPort, nil))
}
