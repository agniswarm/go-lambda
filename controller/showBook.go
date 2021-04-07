package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/agniswarm/go-lambda/response"
	"github.com/agniswarm/go-lambda/services"
)

var IsbnRegexp = regexp.MustCompile(`[0-9]{3}\-[0-9]{10}`)

func Show(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	isbn := req.URL.Query().Get("isbn")
	if !IsbnRegexp.MatchString(isbn) {
		response.ClientError(w, http.StatusBadRequest)
		return
	}
	bk, err := services.GetBook(isbn)
	if err != nil {
		response.ClientError(w, http.StatusBadRequest)
		return
	}
	if bk == nil {
		response.ClientError(w, http.StatusNotFound)
		return
	}
	js, err := json.Marshal(bk)
	if err != nil {
		response.ServerError(w, err)
		return
	}
	response.SendResponse(w, http.StatusOK, js)
	return
}
