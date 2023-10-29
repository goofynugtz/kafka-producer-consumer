package database

import (
	"context"
	"log"

	"github.com/goofynugtz/kafka-producer-consumer/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DatabaseInit() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Env.MONGODB_URI))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

var Client *mongo.Client = DatabaseInit()
var ProductCollection *mongo.Collection = OpenCollection(Client, "products")

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database(config.Env.DB_NAME).Collection(collectionName)
	return collection
}
