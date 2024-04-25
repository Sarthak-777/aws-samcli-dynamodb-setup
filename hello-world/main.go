package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Response struct {
	Message string `json:"message"`
}

type OrderD struct {
	OrderId string `json:"id" dynamodbav:"id"`
}

func handler(request events.APIGatewayProxyRequest) ([]OrderD, error) {
	order := make([]OrderD, 0)
	db := dbConnect()

	result, err := db.Scan(&dynamodb.ScanInput{
		TableName: aws.String("orders_table"),
	})
	if err != nil {
		return nil, err
	}
	for _, item := range result.Items {
		// Access attributes of each item
		id := item["orderId"].(*dynamodb.AttributeValue).S
		// Do something with the attributes...

		newOrder := OrderD{*id}

		order = append(order, newOrder)
	}
	return order, nil
	// body, err := json.Marshal(result)
	// if err != nil {
	// 	return events.APIGatewayProxyResponse{}, fmt.Errorf("failed to marshal result: %v", err)
	// }

	// addOrders(db, "123")

	// return events.APIGatewayProxyResponse{
	// 	StatusCode: 200,
	// 	Body:       string(body),
	// }, nil
	// return events.APIGatewayProxyResponse{
	// 	StatusCode: 200,
	// 	Body:       "success",
	// }, nil

}

func addOrders(db *dynamodb.DynamoDB, orderId string) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("orders_table"),
		Item: map[string]*dynamodb.AttributeValue{
			"orderId": {
				S: aws.String(orderId),
			},
		},
	}

	_, err := db.PutItem(input)
	if err != nil {
		fmt.Println("Error calling PutItem:", err)
	} else {
		fmt.Println("Successfully called PutItem")
	}

	return err
}

func main() {
	lambda.Start(handler)
}
