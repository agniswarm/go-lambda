package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/agniswarm/go-lambda/router"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
)

func main() {
	if os.Getenv("LOCAL") != "true" {
		lambda.Start(handler)
	} else {
		local()
	}
}

func handler(context context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	adapter := gorillamux.New(router.Router())
	return adapter.ProxyWithContext(context, event)
}

func local() {
	log.Fatal(http.ListenAndServe(":3000", router.Router()))
}
