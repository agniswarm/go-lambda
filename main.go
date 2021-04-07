package main

import (
	"context"
	"fmt"

	"github.com/agniswarm/go-lambda/router"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
)

func main() {
	// if os.Getenv("LOCAL") != "true" {
	lambda.Start(handler)
	// } else {
	// local()
	// }
}

func handler(context context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// js, _ := json.Marshal(event)
	// fmt.Println(string(js))
	adapter := gorillamux.New(router.Router())
	res, err := adapter.ProxyWithContext(context, event)
	fmt.Println(res)
	return res, err
}

// func base64Decode(str string) (string, bool) {
// 	data, err := base64.StdEncoding.DecodeString(str)
// 	if err != nil {
// 		return "", true
// 	}
// 	return string(data), false
// }

// func local() {
// 	log.Fatal(http.ListenAndServe(":3000", router.Router()))
// }
