package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, payload map[string]interface{}) (string, error) {

	fmt.Println("Failure function called")

	return "Failed!", nil
}

func main() {
	lambda.Start(handler)
}
