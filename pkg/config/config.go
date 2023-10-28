package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetTopic() string {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	topic := os.Getenv("KAFKA_TOPIC")
	return topic
}

var KafkaTopic string = GetTopic()