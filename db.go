package main

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Project struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Src         string   `json:"src"`
	Stack       []string `json:"stack"`
}

func getClient() *mongo.Client {
	uri := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return client
}

func GetRecipes() []primitive.M {
	client := getClient()
	cur, err := client.Database("recipes-api").Collection("recipe").Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}

	var results []bson.M
	err = cur.All(context.TODO(), &results)
	return results
}
