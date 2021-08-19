package health

import (
	"net/http"

	"github.com/ahmedwaleedmalik/chuck-norris-api/pkg/util"
)

// genericHandlers directs REST request to appropriate method
func (c *healthService) readinessHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.readiness(w, r)
		return
	default:
		util.MethodNotAllowed(w)
		return
	}
}

// livenessHandler directs REST request to appropriate method
func (c *healthService) livenessHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.liveness(w, r)
		return
	default:
		util.MethodNotAllowed(w)
		return
	}
}
