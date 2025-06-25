package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var DB *mongo.Client

func ConnectDB() {
	uri := "mongodb+srv://schakali:s494@s494.j1uqejp.mongodb.net/?retryWrites=true&w=majority&appName=S494"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB ping failed:", err)
	}

	fmt.Println("Connected to MongoDB Atlas!")
	DB = client
}

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Database("gonotesdb").Collection(collectionName)
}
