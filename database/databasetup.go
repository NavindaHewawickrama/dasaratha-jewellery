package database

import (
	"context"
	"fnt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.CLient {
	client, err := mongo.NewClient(options.Client().ApplyURL("mongodb://localhost:3000"))

	if err != null {
		log.fatal(err)
	}

	ctx, cancel := context.WithTimeOut(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != null {
		log.Fatal(err)
	}

	client.Ping(context.TODO(), nil)
	if err != null {
		log.Println("failed to connect to mongodb")
		return nil
	}

	fnt.Println("Successfully connected to mongodb")
	return client
}

var Client *mongo.Client = DBSet()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("ECommerce").Collection(collectionName)
	return collection
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	var productCollection *mongo.Collection = client.Database("ECommerce").Collection(collectionName)
	return productCollection
}
