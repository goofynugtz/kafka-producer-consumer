package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/goofynugtz/kafka-producer-consumer/pkg/config"
	"github.com/goofynugtz/kafka-producer-consumer/pkg/dao"
	"github.com/goofynugtz/kafka-producer-consumer/pkg/helpers"
	"github.com/goofynugtz/kafka-producer-consumer/pkg/models"
)

type KeyValue struct {
	Key int
	Value string
}

func main() {

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
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			var product models.Product
			if err := json.Unmarshal(e.Value, &product); err != nil {
				fmt.Println(err)
			}
			compressedChan := make(chan KeyValue, len(product.Images))
			var wg sync.WaitGroup

			for key, url := range product.Images {
				wg.Add(1)
				go func(url string, index int, wg *sync.WaitGroup) {
					defer wg.Done()
					
					img, err := helpers.DownloadImage(url)
					if err != nil {
						fmt.Println("Could not download image >> ", url)
					}
					directory := "images/"
					path, err := helpers.SaveAsJPEG(img, directory, 10)
					if err != nil {
						fmt.Println("Could not save image.", err)
					}
					compressedChan <- KeyValue{Key: index, Value: *path}
				}(url, key, &wg)
			}
			wg.Wait()
			close(compressedChan)

			compressedArr := make([]string, len(compressedChan))
			for t := range compressedChan {
				compressedArr[t.Key] = t.Value
				fmt.Println(t.Value)
			}
			product.CompressedImages = compressedArr
			
			if err := dao.UpdateProductCompressedImages(&ctx, &product); err != nil {
				fmt.Println(err)
			}

		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		}
	}
	consumer.Close()
}
