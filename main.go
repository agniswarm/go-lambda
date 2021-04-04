package main

import (
	"golambda/controller"
	"golambda/response"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Router)
}
func Router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return controller.Show(req)
	case "POST":
		return controller.CreateBook(req)
	default:
		return response.ClientError(http.StatusMethodNotAllowed)
	}
}
