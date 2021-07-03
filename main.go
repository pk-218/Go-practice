package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pk-218/api/database"
	"github.com/pk-218/api/routes"
)

func main() {
    database.Connect()
    app := fiber.New()

	routes.Setup(app)
    app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    log.Fatal(app.Listen(":3000"))
}
