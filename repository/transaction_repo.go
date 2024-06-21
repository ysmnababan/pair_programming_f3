package repository

import (
	"context"
	"log"
	"pair_programming/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	DB *mongo.Database
}

type TransactionRepo interface {
	GetTransactionID(id string) (models.Transaction, error)
	GetAllTransaction() ([]models.Transaction, error)
	CreateTransaction(t *models.Transaction) (interface{}, error)
}

func (r *Repo) GetTransactionID(id string) (models.Transaction, error) {
	var data models.Transaction

	// get id and convert to primitive.ObjectID
	obj_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return models.Transaction{}, err
	}
	err = r.DB.Collection("transactions").FindOne(context.Background(), bson.M{"_id": obj_id}).Decode(&data)
	if err != nil {
		log.Println(err)
		return models.Transaction{}, err
	}
	return data, nil
}

func (r *Repo) GetAllTransaction() ([]models.Transaction, error) {
	var datas []models.Transaction
	cursor, err := r.DB.Collection("transactions").Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// iterate through cursor
	for cursor.Next(context.Background()) {
		var data models.Transaction
		if err := cursor.Decode(&data); err != nil {
			log.Println(err)
			return nil, err
		}

		// append the data
		datas = append(datas, data)
	}

	return datas, nil
}

func (r *Repo) CreateTransaction(t *models.Transaction) (interface{}, error) {
	// get price from product
	var p models.Product
	err := r.DB.Collection("products").FindOne(context.Background(), bson.M{"_id": t.ProductID}).Decode(&p)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	t.TotalPrice = t.Quantity * p.Price
	res, err := r.DB.Collection("transactions").InsertOne(context.Background(), *t)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	
	return res, nil
}
