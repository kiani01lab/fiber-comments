package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DBURI = "mongodb://db:27017"
const DBNAME = "comments"

func ConnectToMongo() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(DBURI)
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	clientOptions.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to mongo...")

	return client, nil
}
