package controllers

import (
	"book-store/api/config"
	"book-store/api/models"
	"book-store/api/types"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {

	return c.SendString("All books")

}

func CreateBook(c *fiber.Ctx) error {

	book := new(types.Book)

	err := c.BodyParser(&book)
	if err != nil {
		log.Fatal("binding problem")
	}

	booker := &models.Book{
		Name:        book.Name,
		Description: book.Description,
		Author:      book.Author,
		Tags:        book.Tags,
	}
	config.DB.Create(booker)

	return c.JSON(fiber.Map{
		"book": booker,
	})
}
