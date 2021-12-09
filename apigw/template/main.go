package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	res := "this is template"
	return events.APIGatewayProxyResponse{
		Body:       res,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
