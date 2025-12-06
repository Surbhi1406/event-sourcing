package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"eventSourcing/internal/models"
)

func ConnectToDB(uri string, dbName string) (*mongo.Client, *mongo.Database, error) {

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(nil, clientOptions)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}
	err = client.Ping(nil, nil)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}
	log.Println("Connected to MongoDB!")

	db := client.Database(dbName)
	log.Printf("Connected to MongoDB: %s/%s", uri, dbName)
	return client, db, nil
}

func AddOrderInDB(ctx context.Context, db *mongo.Database, order models.Order) error {
	collection := db.Collection("orders")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, order)
	if err != nil {
		return err
	}

	log.Printf("Inserted order with ID: %v", res.InsertedID)

	return nil
}

func GetOrdersFromDB(ctx context.Context, db *mongo.Database) ([]models.Order, error) {
	collection := db.Collection("orders")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []models.Order
	if err = cursor.All(ctx, &orders); err != nil {
		return nil, err
	}

	log.Printf("Fetched orders: %v", orders)
	return orders, nil
}
