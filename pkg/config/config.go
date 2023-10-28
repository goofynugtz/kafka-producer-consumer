package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	database "github.com/goofynugtz/kafka-producer-consumer/pkg/db"

)

func GetTopic() string {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	topic := os.Getenv("KAFKA_TOPIC")
	return topic
}

var KafkaTopic string = GetTopic()
var ProductCollection *mongo.Collection = database.OpenCollection(database.Client, "products")
