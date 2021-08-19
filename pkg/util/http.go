package util

import (
	"encoding/json"
	"net/http"
)

const methodNotAllowedError = "Method not allowed"

// JSONResponse makes the response with payload as json format
func JSONResponse(w http.ResponseWriter, status int, payload interface{}) {
	// Encode payload to JSON
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func MethodNotAllowed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte(methodNotAllowedError))
}
