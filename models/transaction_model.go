package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserID     primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	ProductID  primitive.ObjectID `json:"product_id" bson:"product_id,omitempty"`
	TotalPrice int                `json:"total_price" bson:"total_price"`
	Quantity   int                `json:"quantity" bson:"quantity"`
}
