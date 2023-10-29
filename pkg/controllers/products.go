package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"github.com/goofynugtz/kafka-producer-consumer/config"
	"github.com/goofynugtz/kafka-producer-consumer/pkg/dao"
	models "github.com/goofynugtz/kafka-producer-consumer/pkg/models"
	p "github.com/goofynugtz/kafka-producer-consumer/pkg/producer"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

		if err := dao.AddProduct(&ctx, &product); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Product item was not created"})
		}
		productJSON, _ := json.Marshal(product)
		if err := p.KafkaProducer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &config.Env.KAFKA_TOPIC, Partition: kafka.PartitionAny},
			Value:          productJSON,
		},
			p.DeliveryChan,
		); err != nil {
			fmt.Printf("Could not pass product_id %v due to %v", product.ID, err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "Product item created"})
	}
}
