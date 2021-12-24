package routes

import (
	"github.com/Jeanpigi/go-api/src/models"
	"github.com/gofiber/fiber/v2"
)

// SetRoutes allow set up all routes you need it
func SetRoutes(app *fiber.App) {
	app.Get("/api/v1/book", models.GetAllBooks)
	app.Post("/api/v1/book", models.InsertBook)
	app.Put("/api/v1/book/:id", models.UpdateBook)
	app.Delete("/api/v1/book/:id", models.DeleteBook)
}