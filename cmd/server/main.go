package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kiani01lab/fiber-comments/db"
	"github.com/kiani01lab/fiber-comments/routes"
)

func main() {
	listenAddr := flag.String("listernAddress", ":8000", "the listen address of the service")
	flag.Parse()

	app := fiber.New()

	client, err := db.ConnectToMongo()
	if err != nil {
		log.Panic(err)
	}

	routes.SetupRoutes(app, client)

	app.Listen(*listenAddr)
}
