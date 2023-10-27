package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	database "github.com/goofynugtz/kafka-producer-consumer/db"
	models "github.com/goofynugtz/kafka-producer-consumer/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = database.OpenCollection(database.Client, "products")
// var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

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
		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}
