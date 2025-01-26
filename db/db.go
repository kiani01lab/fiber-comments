package db

import (
	"context"
	"log"

	"github.com/kiani01lab/fiber-comments/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	User UserStore
}

func ConnectToMongo() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(config.Config("DB_URI"))
	clientOptions.SetAuth(options.Credential{
		Username: config.Config("DB_USERNAME"),
		Password: config.Config("DB_PASSWORD"),
	})

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to mongo...")

	return client, nil
}
