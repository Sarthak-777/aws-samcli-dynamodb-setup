package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func dbConnect() *dynamodb.DynamoDB {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-west-2"),
		Endpoint: aws.String("http://192.168.1.86:8000"),
	})

	if err != nil {
		fmt.Println("failed to create session,", err)
		return nil
	}

	db := dynamodb.New(sess)
	return db
	// Use the db object for operations like put, get, query, etc.
}
