package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)


func ConnectDB(uri string) *mongo.Client{

	// Connects to MongoDB
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(opts)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}
	// Pings the database to verify connection
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal("Could not ping MongoDB:", err)
	}
	return client
}

func DisconnectDB(DB *mongo.Client) {
	if err := DB.Disconnect(context.TODO()); err != nil {
		log.Fatal("Error disconnecting from MongoDB:", err)
	}
}
