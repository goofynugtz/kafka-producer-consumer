package dao

import (
	"context"

	"github.com/goofynugtz/kafka-producer-consumer/pkg/config"
	"github.com/goofynugtz/kafka-producer-consumer/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetProductById(pid string) (*models.Product, error) {
	var product models.Product
	if err := config.ProductCollection.FindOne(context.Background(), bson.M{"_id": pid}).Decode(&product); err != nil {
		return nil, err
	}
	return &product, nil
}

func AddProduct(ctx *context.Context, p *models.Product) error {
	_, err := config.ProductCollection.InsertOne(*ctx, p)
	if err != nil {
		return err
	}
	return nil
}
