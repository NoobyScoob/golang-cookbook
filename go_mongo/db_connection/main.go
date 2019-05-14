package main

// db connection with mongodb sample
// dependecies mongo-driver (official)

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	// handle connection error
	if err != nil {
		log.Fatal(err)
	}

	// ping connection
	err = client.Ping(context.TODO(), nil)
	// handle err
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to mongodb")

}