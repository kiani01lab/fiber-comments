package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kiani01lab/fiber-comments/cmd/handlers"
	"github.com/kiani01lab/fiber-comments/db"
	"github.com/spf13/viper"
)

func main() {
	listenAddr := flag.String("listernAddress", ":8000", "the listen address of the servie")
	flag.Parse()

	viperEvaluated()

	app := fiber.New()
	api := app.Group("/api")

	client, err := db.ConnectToMongo()
	if err != nil {
		log.Panic(err)
	}

	userHandler := handlers.NewUserHandler(db.NewMongoUserStore(client))

	api.Post("/user/", userHandler.HandlePostUser)
	api.Get("/users", userHandler.HandleGetUsers)
	api.Get("/user/:id", userHandler.HandleGetUser)
	api.Delete("/user/:id", userHandler.HandleDeleteUser)
	app.Listen(*listenAddr)
}

func viperEvaluated() {
	viper.SetConfigFile("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("unable to evaluate Viper: %w", err))
	}
	log.Println("Viper config evaluated")
}
