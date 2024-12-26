package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kiani01lab/fiber-comments/cmd/handlers"
	"github.com/kiani01lab/fiber-comments/db"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(app *fiber.App, client *mongo.Client) {
	api := app.Group("/api")

	userHandler := handlers.NewUserHandler(db.NewMongoUserStore(client))

	api.Post("/user/", userHandler.HandlePostUser)
	api.Get("/users", userHandler.HandleGetUsers)
	api.Get("/user/:id", userHandler.HandleGetUser)
	api.Delete("/user/:id", userHandler.HandleDeleteUser)

}
