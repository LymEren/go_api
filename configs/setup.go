package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	// Created new client
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))

	if err != nil {
		log.Fatal("EnvMongoURI Error:", err)
	}

	//System is try to connect in 20 seconds
	ctx, baaa := context.WithTimeout(context.Background(), 20*time.Second)
	err = client.Connect(ctx)
	fmt.Println(baaa, err)

	// Added ping for the test
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalln("Client Ping (setup.go)", err)
	}
	// If there are no errors program will run
	return client
}

// We have defined global variable DB *mongo.Client data type
var DB *mongo.Client = ConnectDB()

// And we created Collections function. Its same type to Tables in sql databases
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	// Created client database what name is "golangAPI"
	return client.Database("golangAPI").Collection(collectionName)
}
