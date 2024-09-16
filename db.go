package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Project struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Src         string   `json:"src"`
	Stack       []string `json:"stack"`
}

func getClient() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	dynamoClient := dynamodb.NewFromConfig(cfg)
	return dynamoClient
}

func GetRecipes(dynamoClient *dynamodb.Client) []Project {
	response, err := dynamoClient.Scan(context.TODO(), &dynamodb.ScanInput{TableName: aws.String(("projects"))})
	println(response)
	projects := []Project{}
	err = attributevalue.UnmarshalListOfMaps(response.Items, &projects)
	if err != nil {
		log.Fatal(err)
	}
	return projects
}
