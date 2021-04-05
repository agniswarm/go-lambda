package response

import (
	"fmt"
	"net/http"
)

func SendResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, message)
}
