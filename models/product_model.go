package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ProductName string             `json:"product_name"`
	Price       int                `json:"price"`
	Description string             `json:"description,omitempty"`
}
