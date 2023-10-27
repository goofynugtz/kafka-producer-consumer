package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id"`
	Name             string             `json:"name"`
	Description      string             `json:"description"`
	Images           []string           `json:"images"`
	Price            int                `json:"price"`
	CompressedImages []string           `json:"compresses_images"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
}
