package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/goofynugtz/kafka-producer-consumer/pkg/config"
	"github.com/goofynugtz/kafka-producer-consumer/pkg/models"
	// "github.com/goofynugtz/kafka-producer-consumer/pkg/dao"
)

func main() {
	
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "foo",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(">> consumers", config.KafkaTopic)
	if err := consumer.SubscribeTopics([]string{config.KafkaTopic}, nil); err != nil {
		log.Fatal(err)
	}
	run := true
	for run {
		ev := consumer.Poll(10)
		switch e := ev.(type) {
		case *kafka.Message:

			fmt.Println(e)
			var product models.Product
			if err := json.Unmarshal(e.Value, &product); err != nil {
				fmt.Println(err)
			}
			fmt.Println(">> >> ", product)
			
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		}
	}
	consumer.Close()
}
