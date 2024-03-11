package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leeroyakbar/restapi-fiber/controller"
)

func Router(app *fiber.App) {
	base := app.Group("/api")

	book := base.Group("/books")
	book.Get("/", controller.Hello)
}
