package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Movie struct {
	Title string
	Year  int
}

func addMovie(db *dynamodb.DynamoDB, title string, year int) error {
	// fmt.Println("Adding movie:", title, year)
	input := &dynamodb.PutItemInput{
		TableName: aws.String("MyDDBLocal"),
		Item: map[string]*dynamodb.AttributeValue{
			"orderId": {
				S: aws.String(title),
			},
		},
	}

	// fmt.Println("Calling PutItem")
	_, err := db.PutItem(input)
	if err != nil {
		fmt.Println("Error calling PutItem:", err)
	} else {
		fmt.Println("Successfully called PutItem")
	}

	return err
}

func printItems(db *dynamodb.DynamoDB) {
	input := &dynamodb.ScanInput{
		TableName: aws.String("Movies"),
	}

	result, err := db.Scan(input)
	if err != nil {
		fmt.Println("Got error calling Scan:")
		fmt.Println(err.Error())
		return
	}

	for _, i := range result.Items {
		movie := Movie{}

		err = dynamodbattribute.UnmarshalMap(i, &movie)
		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			continue
		}

		fmt.Println("Title: ", movie.Title)
		fmt.Println("Year: ", movie.Year)
	}
}
