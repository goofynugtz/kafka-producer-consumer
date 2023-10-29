package dao

import (
	"context"
	"fmt"

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

func UpdateProductCompressedImages(ctx *context.Context, p *models.Product) error {
	filter := bson.M{"_id": p.ID}
	fmt.Println("update ", p.CompressedImages)
	update := bson.M{"$set": bson.M{"compressed_images": p.CompressedImages}}
	_, err := config.ProductCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
