package controller

import (
	"encoding/json"
	"fmt"
	"golambda/models"
	"golambda/response"
	"golambda/services"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func CreateBook(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if req.Headers["content-type"] != "application/json" && req.Headers["Content-Type"] != "application/json" {
		return response.ClientError(http.StatusNotAcceptable)
	}
	bk := new(models.Book)
	err := json.Unmarshal([]byte(req.Body), bk)
	if err != nil {
		return response.ClientError(http.StatusUnprocessableEntity)
	}
	if !IsbnRegexp.MatchString(bk.ISBN) {
		return response.ClientError(http.StatusBadRequest)
	}
	// fmt.Println("position 2")
	if bk.Author == "" || bk.Title == "" {
		return response.ClientError(http.StatusBadRequest)
	}

	// fmt.Println("position 3")
	// fmt.Println(bk.ISBN)

	err = services.CreateBook(bk)
	if err != nil {
		return response.ServerError(err)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Headers: map[string]string{
			"Location": fmt.Sprintf("/books?isbn=%s", bk.ISBN),
		},
	}, nil
}
