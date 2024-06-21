package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context, dbName string) (*mongo.Client, *mongo.Database) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Error: ", err)
	}

	// ping the primary to verify connectivity
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Println("Connected to MongoDB")
	db := client.Database(dbName)
	return client, db
}
