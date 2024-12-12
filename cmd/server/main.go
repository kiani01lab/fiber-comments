package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listernAddress", ":8000", "the listen address of the servie")
	flag.Parse()

	app := fiber.New()
	app.Listen(*listenAddr)
}
