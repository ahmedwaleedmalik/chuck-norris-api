package banter

import (
	"log"
	"net/http"
	"sync"

	"github.com/ahmedwaleedmalik/chuck-norris-api/internal/models"
	"github.com/ahmedwaleedmalik/chuck-norris-api/pkg/util"
	"github.com/jinzhu/gorm"
)

const banterEndpoint = "/banter"

type banterService struct {
	//The lock can be held by an arbitrary number of readers or a single writer
	sync.RWMutex

	// db is a client that holds the connection to database
	db *gorm.DB
}

// NewBanterService Returns new banterService
func NewBanterService(db *gorm.DB) *banterService {
	return &banterService{
		db: db,
	}
}

// RegisterBanterServiceEndpoints registers endpoint and handlers for banter service
func (c *banterService) RegisterBanterServiceEndpoints() {
	http.HandleFunc(banterEndpoint, c.genericHandlers)
}

// list performs LIST action and returns all stored configs
func (c *banterService) list(w http.ResponseWriter, r *http.Request) {
	log.Printf("REST request to list all jokes\n")

	c.Lock()

	var jokes []models.Joke
	c.db.Find(&jokes)

	c.Unlock()

	// Respond with appropriate status code and payload as JSON
	util.JSONResponse(w, http.StatusOK, jokes)
}
