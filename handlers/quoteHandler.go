package handlers

import (
	"encoding/json"
	"net/http"
	"veda-backend/api"
)

func HandleQuoteRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Encode the quote in JSON format and send it in the response
	json.NewEncoder(w).Encode(api.GetQuoteInJSON())
	
}
