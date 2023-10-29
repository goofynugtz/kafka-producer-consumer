package config

import (
	"fmt"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

type Environment struct {
	KAFKA_TOPIC string
	MONGODB_URI string
	DB_NAME     string
}

func Config() Environment {
	projectName := regexp.MustCompile(`^(.*` + "kafka-producer-consumer" + `)`)
	currentWorkDirectory, _ := os.Getwd()
  rootPath := projectName.Find([]byte(currentWorkDirectory))

	// Absolute path required for test cases to pass.
	if err := godotenv.Load(string(rootPath) + `/.env`); err != nil {
		fmt.Println("No .env file found")
	}
	var e Environment
	e.KAFKA_TOPIC = os.Getenv("KAFKA_TOPIC")
	e.MONGODB_URI = os.Getenv("MONGODB_URI")
	e.DB_NAME = os.Getenv("DB_NAME")
	return e
}

var Env Environment = Config()
