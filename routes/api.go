package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/kiani01lab/fiber-comments/cmd/handlers"
	"github.com/kiani01lab/fiber-comments/db"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(app *fiber.App, client *mongo.Client) {
	api := app.Group("/api")

	var (
		userStore = db.NewMongoUserStore(client)
		store     = &db.Store{
			User: userStore,
		}
		userHandler = handlers.NewUserHandler(store)
		authHandler = handlers.NewAuthHandler(store)
	)

	// Swagger
	api.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:          "http://localhost:8000/doc.json",
		DeepLinking:  false,
		DocExpansion: "none",
	}))

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", authHandler.Login)

	// User
	api.Post("/user/", userHandler.HandlePostUser)
	api.Get("/users", userHandler.HandleGetUsers)
	api.Get("/user/:id", userHandler.HandleGetUser)
	api.Delete("/user/:id", userHandler.HandleDeleteUser)
}
