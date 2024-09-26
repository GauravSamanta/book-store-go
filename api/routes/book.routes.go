package routes

import (
	"book-store/api/controllers"

	"github.com/gofiber/fiber/v2"
)

func BookRoutes(r *fiber.App) {
	route := r.Group("/api/v1/book")

	route.Get("/", controllers.GetBooks)
	route.Post("/create", controllers.CreateBook)

}
