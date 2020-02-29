package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.S3Event) (string, error) {
	// Do some stuff, then return the file name to the success function
	return request.Records[0].S3.Object.Key, nil
}

func main() {
	lambda.Start(handler)
}
