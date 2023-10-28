package producer

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ProducerInit() *kafka.Producer {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"acks":              "all"},
	)
	if err != nil {
		log.Fatal(err)
	}

	return producer
}

var KafkaProducer *kafka.Producer = ProducerInit()
var DeliveryChan chan kafka.Event = make(chan kafka.Event, 10000)
