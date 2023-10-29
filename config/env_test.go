package config_test

import (
	// "fmt"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/goofynugtz/kafka-producer-consumer/config"
	// "github.com/joho/godotenv"
)

func TestConfig(t *testing.T) {
	e := config.Config()

	assert.NotEqual(t, e.DB_NAME, "");
	assert.NotEqual(t, e.KAFKA_TOPIC, "");
	assert.NotEqual(t, e.MONGODB_URI, "");
}
