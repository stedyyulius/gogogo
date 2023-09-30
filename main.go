package main

import (
	"context"
	"fmt"
	"iseng/configs"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
}

func main() {
	
	loadEnv()

	uri := os.Getenv("MONGO_URI")

	opts := options.Client().ApplyURI(uri)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	db := client.Database("dbpool")
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := db.RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}

	fmt.Println("Mongo DB connected to ", uri)

	configs.SetupRoutes(db)

}