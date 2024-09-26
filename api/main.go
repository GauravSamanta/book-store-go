package main

import (
	"book-store/api/config"
	"book-store/api/routes"

	"github.com/gofiber/fiber/v2"
)

func init() {

	config.LoadEnv()
	config.ConnectDB()
	config.MigrateDB()

}

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	routes.BookRoutes(app)

	app.Listen(":5000")

}
