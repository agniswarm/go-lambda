package controller

import (
	"encoding/json"
	"net/http"

	"github.com/agniswarm/go-lambda/models"
	"github.com/agniswarm/go-lambda/response"
	"github.com/agniswarm/go-lambda/services"
)

func CreateBook(w http.ResponseWriter, req *http.Request) {

	if req.Header.Get("content-type") != "application/json" && req.Header.Get("Content-Type") != "application/json" {
		response.ClientError(w, http.StatusNotAcceptable)
		return
	}
	bk := new(models.Book)
	err := json.NewDecoder(req.Body).Decode(&bk)
	defer req.Body.Close()
	if err != nil {
		response.ClientError(w, http.StatusUnprocessableEntity)
		return
	}

	if !IsbnRegexp.MatchString(bk.ISBN) {
		response.ClientError(w, http.StatusBadRequest)
		return
	}

	if bk.Author == "" || bk.Title == "" {
		response.ClientError(w, http.StatusBadRequest)
		return
	}

	err = services.CreateBook(bk)
	if err != nil {
		response.ServerError(w, err)
		return
	}

	res, err := json.Marshal(bk)
	if err != nil {
		response.ServerError(w, err)
	}
	response.SendResponse(w, http.StatusCreated, string(res))
	return
	// return events.APIGatewayProxyResponse{
	// 	StatusCode: http.StatusCreated,
	// 	Headers: map[string]string{
	// 		"Location": fmt.Sprintf("/books?isbn=%s", bk.ISBN),
	// 	},
	// }, nil
}
