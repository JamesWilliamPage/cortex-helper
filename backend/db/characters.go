package db

import (
	"github.com/JamesWilliamPage/cortex-helper-backend/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func CreateTable(svc *dynamodb.DynamoDB) error {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Name"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Name"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
		TableName: aws.String("Characters"),
	}

	_, err := svc.CreateTable(input)
	return err
}

func PutCharacter(svc *dynamodb.DynamoDB, character types.Character) error {
	av, err := dynamodbattribute.MarshalMap(character)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Characters"),
	}

	_, err = svc.PutItem(input)
	return err
}

func GetCharacter(svc *dynamodb.DynamoDB, name string) (*types.Character, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Name": {
				S: aws.String(name),
			},
		},
		TableName: aws.String("Characters"),
	}

	result, err := svc.GetItem(input)
	if err != nil {
		return nil, err
	}

	character := new(types.Character)
	err = dynamodbattribute.UnmarshalMap(result.Item, character)
	return character, err
}
