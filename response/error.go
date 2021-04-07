package response

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

func ServerError(w http.ResponseWriter, err error) {
	errorLogger.Println(err.Error())

	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fmt.Sprintf("{\"message\": %s", http.StatusText(http.StatusInternalServerError)))
}

func ClientError(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fmt.Sprintf("{\"message\": %s", http.StatusText(statusCode)))

}
