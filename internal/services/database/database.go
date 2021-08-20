package database

import (
	"fmt"
	"log"
	"os"

	"github.com/ahmedwaleedmalik/chuck-norris-api/internal/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	// Environment Variables for database configuration
	sqlHostEnv  = "SQL_HOST"
	sqlPortEnv  = "SQL_PORT"
	usernameEnv = "SQL_USERNAME"
	passwordEnv = "SQL_PASSWORD"
	databaseEnv = "SQL_DATABASE"

	// Default values for database configuration
	defaultSqlHost  = "localhost"
	defaultSqlPort  = "3306"
	defaultUsername = "root"
	defaultDatabase = "banter"
)

type config struct {
	serverHost string
	serverPort string
	username   string
	password   string
	database   string
}

// InitializeDatabase initializes the database connection
func InitializeDatabase() (*gorm.DB, error) {

	// Load database configuration
	config, err := loadDatabaseConfig()
	if err != nil {
		return nil, err
	}

	// Create connection string
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", config.username, config.password, config.serverHost, config.serverPort, config.database)

	// Connect to database
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connection to database %s was successful!", config.database)

	// Migrate the schema
	db = db.AutoMigrate(&models.Joke{})

	// Add dummy data in the database, if the database is empty
	var jokes []models.Joke
	db.Find(&jokes)

	if len(jokes) == 0 {
		log.Printf("Database is empty, adding dummy data")
		err = insertDummyData(db)
	}

	return db, err
}

// insertDummyData inserts dummy data into the database
func insertDummyData(db *gorm.DB) error {
	// Alternate to using transactions can be direct queries like:
	// db.Exec("INSERT INTO jokes (id, joke) VALUES (1, 'Chuck Norris does not dodge bullets, bullets dodge Chuck Norris')")

	// Begin transaction
	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&models.Joke{ID: 1, Joke: "Chuck Norris does not dodge bullets, bullets dodge Chuck Norris"}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(&models.Joke{ID: 2, Joke: "Once, Chuck Norris held a judge in contempt, while in court."}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&models.Joke{ID: 3, Joke: "Chuck Norris once found a piece of hay in a neelestack."}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&models.Joke{ID: 4, Joke: "Chuck Norris is better than Ash Ketchum because Ash didn't catch all the Pokemon yet."}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&models.Joke{ID: 5, Joke: "Documentaries about Chuck Norris' life can be found in video libraries in the horror section."}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// End transaction and commit changeset
	return tx.Commit().Error
}

// loadDatabaseConfig loads the database configuration from environment variables
func loadDatabaseConfig() (config, error) {
	config := config{}

	// Retrieve serverHost from environment variables
	config.serverHost = os.Getenv(sqlHostEnv)
	if len(config.serverHost) == 0 {
		config.serverHost = defaultSqlHost
	}

	// Retrieve serverPort from environment variables
	config.serverPort = os.Getenv(sqlPortEnv)
	if len(config.serverPort) == 0 {
		config.serverPort = defaultSqlPort
	}

	// Retrieve username from environment variables
	config.username = os.Getenv(usernameEnv)
	if len(config.username) == 0 {
		config.username = defaultUsername
	}

	// Retrieve passwordEnv from environment variables
	config.password = os.Getenv(passwordEnv)
	if len(config.password) == 0 {
		log.Fatal(passwordEnv + " environment variable not set")
	}

	// Retrieve serverURL from environment variables
	config.database = os.Getenv(databaseEnv)
	if len(config.database) == 0 {
		config.database = defaultDatabase
	}
	return config, nil
}
