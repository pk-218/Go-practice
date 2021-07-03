package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pk-218/api/database"
	"github.com/pk-218/api/models"
)

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []models.Book
	db.Find(&books)
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book models.Book
	db.Find(&book, id)
	return c.JSON(book)
}

func NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	book := new(models.Book)
	// if err := c.BodyParser(book); err != nil {
	// 	c.Status(503).SendString(err)
	// 	return 
	// }
	db.Create(&book)
	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book models.Book
	db.First(&book, id)
	// if book.Title == "" {
	// 	c.Status(500).SendString("No book found with given ID")
	// 	return
	// }
	db.Delete(&book)
	return c.SendString("Book successfully deleted")
}
