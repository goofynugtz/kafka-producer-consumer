package dao

import (
	"context"

	database "github.com/goofynugtz/kafka-producer-consumer/pkg/db"
	"github.com/goofynugtz/kafka-producer-consumer/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetProductById(pid string) (*models.Product, error) {
	var product models.Product
	if err := database.ProductCollection.FindOne(context.Background(), bson.M{"_id": pid}).Decode(&product); err != nil {
		return nil, err
	}
	return &product, nil
}

func AddProduct(ctx *context.Context, p *models.Product) error {
	_, err := database.ProductCollection.InsertOne(*ctx, p)
	if err != nil {
		return err
	}
	return nil
}

func UpdateProductCompressedImages(ctx *context.Context, p *models.Product) error {
	filter := bson.M{"_id": p.ID}
	update := bson.M{"$set": bson.M{"compressed_images": p.CompressedImages}}
	_, err := database.ProductCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
