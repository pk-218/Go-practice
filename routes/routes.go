package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pk-218/api/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/api/book", controllers.GetBooks)
	app.Get("/api/book/:id", controllers.GetBook)
	app.Post("/api/book", controllers.NewBook)
	app.Post("/api/book/:id", controllers.DeleteBook)
}