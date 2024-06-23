package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnecToDB() {
	// MongoDB connection URI
//	mongoURI := "mongodb://admin:password@localhost:27017" //when running locally

	mongoURI := "mongodb://admin:password@mongodb:27017" //when running inside docker container


	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	var err error

	// Connect to MongoDB
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the MongoDB server to check if connection was successful
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}
	fmt.Println("Connected to MongoDB!")

}

func Insert() {
	database := client.Database("mydb")
	sampleCollection := database.Collection("user")

	insertedDocument := bson.M{
		"name":       "michael",
		"content":    "test content",
		"bank_money": 1000,
		"create_at":  time.Now(),
	}
	insertedResult, err := sampleCollection.InsertOne(context.Background(), insertedDocument)

	if err != nil {
		log.Fatalf("inserted error : %v", err)
		return
	}
	fmt.Println("======= inserted id ================")
	log.Printf("inserted ID is : %v", insertedResult.InsertedID)

}
