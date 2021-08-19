package banter

import (
	"net/http"

	"github.com/ahmedwaleedmalik/chuck-norris-api/pkg/util"
)

// genericHandlers directs REST request to appropriate method
func (c *banterService) genericHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.list(w, r)
		return
	default:
		util.MethodNotAllowed(w)
		return
	}
}
