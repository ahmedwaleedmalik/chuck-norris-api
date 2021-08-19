package main

import (
	"log"
	"net/http"

	"github.com/ahmedwaleedmalik/chuck-norris-api/internal/services/banter"
)

const defaultServingPort = "8080"

func main() {

	// TODO: Initialize DB here first

	// Create new Banter Service Instance
	configService := banter.NewBanterService()

	// Register Endpoints for Config Service
	configService.RegisterConfigServiceEndpoints()

	// Start serving the http server
	log.Printf("HTTP server listening on %v", defaultServingPort)
	log.Fatal(http.ListenAndServe(":"+defaultServingPort, nil))
}
