package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"veda-backend/api"
)

func AddLoginTimeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body to get the FromTime parameter
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fromTimeStr := r.Form.Get("fromTime")
	fromTime, err := strconv.Atoi(fromTimeStr)
	if err != nil {
		http.Error(w, "Invalid fromTime parameter", http.StatusBadRequest)
		return
	}

	// Call the AddLoginTime function
	err = api.AddLoginTime(fromTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a success response
	w.WriteHeader(http.StatusOK)
}

func AddLogoutTimeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body to get the ToTime and SessionID parameters
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	toTimeStr := r.Form.Get("toTime")
	toTime, err := strconv.Atoi(toTimeStr)
	if err != nil {
		http.Error(w, "Invalid toTime parameter", http.StatusBadRequest)
		return
	}

	sessionIDStr := r.Form.Get("sessionID")
	sessionID, err := strconv.Atoi(sessionIDStr)
	if err != nil {
		http.Error(w, "Invalid sessionID parameter", http.StatusBadRequest)
		return
	}

	// Call the AddLogoutTime function
	err = api.AddLogoutTime(toTime, sessionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a success response
	w.WriteHeader(http.StatusOK)
}

func ReadUsageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Call the ReadUsage function
	usages, err := api.ReadUsage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the usages into JSON format
	response, err := json.Marshal(usages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and send the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
