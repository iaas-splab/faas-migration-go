package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/c-mueller/faas-migration-go/aws"
	"github.com/c-mueller/faas-migration-go/core"
)

func Handler(ctx context.Context, request aws.Request) (aws.Response, error) {
	var elem core.PutRequest
	err := json.Unmarshal([]byte(request.Body), &elem)
	if err != nil {
		fmt.Printf("Put Failed %q", err.Error())
		return aws.Response{
			StatusCode: 400,
		}, err
	}

	r := aws.NewDynamoRepo()

	res, err := core.Put(elem, r)
	if err != nil {
		fmt.Printf("Put Failed %q", err.Error())
		return aws.Response{
			StatusCode: 500,
		}, err
	}

	d, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("Response Marshal failed %q", err.Error())
		return aws.Response{
			StatusCode: 500,
		}, err
	}

	resp := aws.Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(d),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
