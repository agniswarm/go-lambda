package response

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, statusCode int, message interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	// js, _ := json.Marshal(message)
	json.NewEncoder(w).Encode(message)
}
