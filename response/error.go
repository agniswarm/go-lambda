package response

import (
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
	fmt.Fprintf(w, "{\"message\": \"Internal Server Error\"}")
}

func ClientError(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"message\": %s", http.StatusText(statusCode))

}
