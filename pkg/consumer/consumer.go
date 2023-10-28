package consumer

import (
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/goofynugtz/kafka-producer-consumer/pkg/config"
)

func ConsumerInit() *kafka.Consumer {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "foo",
	})
	if err != nil {
		log.Fatal(err)
	}
	if err := consumer.SubscribeTopics([]string{config.KafkaTopic}, nil); err != nil {
		log.Fatal(err)
	}
	run := true
	for run {
		ev := consumer.Poll(10)
		switch e := ev.(type) {
		case *kafka.Message:
			// application-specific processing
			
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
			fmt.Printf("Ignored %v\n", e)
		}
	}
	return consumer
}

var KafkaConsumer *kafka.Consumer = ConsumerInit()
