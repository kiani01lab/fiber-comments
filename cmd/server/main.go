package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kiani01lab/fiber-comments/db"
	"github.com/kiani01lab/fiber-comments/routes"
)

// @title Fiber-Comments API
// @version 1.0
// @description API for fiber-comments
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /api
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
