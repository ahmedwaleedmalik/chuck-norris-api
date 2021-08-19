package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

const (
	// Environment Variables for database configuration
	serverURLEnv = "SERVER_URL"
	usernameEnv  = "USERNAME"
	passwordEnv  = "PASSWORD"
	databaseEnv  = "DATABASE"

	// Default values for database configuration
	defaultServerURL = "localhost:3306"
	defaultUsername  = "root"
	defaultDatabase  = "banter"
)

type config struct {
	serverURL string
	username  string
	password  string
	database  string
}

// InitializeDatabase initializes the database connection
func InitializeDatabase() (*sql.DB, error) {

	// Load database configuration
	config, err := loadDatabaseConfig()
	if err != nil {
		return nil, err
	}

	// Create connection string
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", config.username, config.password, config.serverURL, config.database)

	// Connect to database
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connection to database %s was successful!", config.database)

	return db, err
}

// loadDatabaseConfig loads the database configuration from environment variables
func loadDatabaseConfig() (config, error) {
	config := config{}

	// Retrieve serverURL from environment variables
	config.serverURL = os.Getenv(serverURLEnv)
	if len(config.serverURL) == 0 {
		config.serverURL = defaultServerURL
	}

	// Retrieve username from environment variables
	config.username = os.Getenv(usernameEnv)
	if len(config.username) == 0 {
		config.username = defaultUsername
	}

	// Retrieve passwordEnv from environment variables
	config.password = os.Getenv(passwordEnv)
	if len(config.password) == 0 {
		panic(passwordEnv + " environment variable not set")
	}

	// Retrieve serverURL from environment variables
	config.database = os.Getenv(databaseEnv)
	if len(config.database) == 0 {
		config.database = defaultDatabase
	}
	return config, nil
}
