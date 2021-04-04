package controller

import (
	"encoding/json"
	"golambda/response"
	"golambda/services"
	"net/http"
	"regexp"

	"github.com/aws/aws-lambda-go/events"
)

var IsbnRegexp = regexp.MustCompile(`[0-9]{3}\-[0-9]{10}`)

func Show(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	isbn := req.QueryStringParameters["isbn"]
	if !IsbnRegexp.MatchString(isbn) {
		return response.ClientError(http.StatusBadRequest)
	}
	bk, err := services.GetBook(isbn)
	if err != nil {
		return response.ServerError(err)
	}
	if bk == nil {
		return response.ClientError(http.StatusNotFound)
	}
	js, err := json.Marshal(bk)
	if err != nil {
		return response.ServerError(err)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
	}, nil
}
