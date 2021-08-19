package util

import (
	"encoding/json"
	"log"
	"net/http"
)

const methodNotAllowedError = "Method not allowed"

// JSONResponse makes the response with payload as json format
func JSONResponse(w http.ResponseWriter, status int, payload interface{}) {
	// Encode payload to JSON
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
		}
		return
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write([]byte(response))
	if err != nil {
		log.Println(err)
	}
}

func MethodNotAllowed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	_, err := w.Write([]byte(methodNotAllowedError))
	if err != nil {
		log.Println(err)
	}
}
