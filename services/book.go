package services

import (
	"golambda/config"
	"golambda/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func GetBook(isbn string) (*models.Book, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Books"),
		Key: map[string]*dynamodb.AttributeValue{
			"ISBN": {
				S: aws.String(isbn),
			},
		},
	}

	result, err := config.DynamoDB.GetItem(input)

	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}
	bk := new(models.Book)
	err = dynamodbattribute.UnmarshalMap(result.Item, bk)
	if err != nil {
		return nil, err
	}
	return bk, nil
}

func CreateBook(bk *models.Book) error {
	input := dynamodb.PutItemInput{
		TableName: aws.String("Books"),
		Item: map[string]*dynamodb.AttributeValue{
			"ISBN":   {S: aws.String(bk.ISBN)},
			"Title":  {S: aws.String(bk.Title)},
			"Author": {S: aws.String(bk.Author)},
		},
	}
	_, err := config.DynamoDB.PutItem(&input)
	if err != nil {
		return err
	}
	return nil
}
