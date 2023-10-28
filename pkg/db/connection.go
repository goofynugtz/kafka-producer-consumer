package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DatabaseInit() *mongo.Client {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	mongoURI := os.Getenv("MONGODB_URI")
	
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

var Client *mongo.Client = DatabaseInit()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("pc-kafka").Collection(collectionName)
	return collection
}
