package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, payload map[string]interface{}) (string, error) {

	// log the payload coming from s3 file processor function
	fmt.Println(payload["responsePayload"])

	return "Success", nil
}

func main() {
	lambda.Start(handler)
}
