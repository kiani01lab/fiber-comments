package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kiani01lab/fiber-comments/cmd/handlers"
	"github.com/kiani01lab/fiber-comments/db"
)

func main() {
	listenAddr := flag.String("listernAddress", ":8000", "the listen address of the servie")
	flag.Parse()

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
	app.Listen(*listenAddr)
}
