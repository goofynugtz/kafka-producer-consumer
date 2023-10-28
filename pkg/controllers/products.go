package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"github.com/goofynugtz/kafka-producer-consumer/pkg/config"
	database "github.com/goofynugtz/kafka-producer-consumer/pkg/db"
	models "github.com/goofynugtz/kafka-producer-consumer/pkg/models"
	p "github.com/goofynugtz/kafka-producer-consumer/pkg/producer"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = database.OpenCollection(database.Client, "products")

func RecieveProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		var request models.RecieveProductSchema
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var product models.Product
		product.ID = primitive.NewObjectID()
		product.Name = request.ProductName
		product.Description = request.ProductDescription
		product.Images = request.ProductImages
		product.Price = request.ProductPrice
		product.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		product.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		resultInsertionNumber, insertErr := productCollection.InsertOne(ctx, product)
		if insertErr != nil {
			fmt.Println(insertErr)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User item was not created"})
			return
		}
		if err := p.KafkaProducer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &config.KafkaTopic, Partition: kafka.PartitionAny},
			Value:          product.ID[:],
		},
			p.DeliveryChan,
		); err != nil {
			fmt.Printf("Could not pass product_id %v due to %v", product.ID, err)
		}

		// e := <-p.DeliveryChan
		// m := e.(*kafka.Message)

		// if m.TopicPartition.Error != nil {
		// 	fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		// } else {
		// 	fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
		// 		*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		// }

		c.JSON(http.StatusOK, resultInsertionNumber)

	}
}
