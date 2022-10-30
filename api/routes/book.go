package routes

import (
	"github.com/tk42/gofiber-surrealdb-template/api/handlers"
	"github.com/tk42/gofiber-surrealdb-template/pkg/book"

	"github.com/gofiber/fiber/v2"
)

// BookRouter is the Router for GoFiber App
func BookRouter(app fiber.Router, service book.Service) {
	app.Get("/books", handlers.GetBooks(service))
	app.Post("/books", handlers.AddBook(service))
	app.Put("/books", handlers.UpdateBook(service))
	app.Delete("/books", handlers.RemoveBook(service))
}
