package router

import (
	"github.com/agniswarm/go-lambda/controller"
	"github.com/agniswarm/go-lambda/middleware"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	bookRoute := app.Group("/book")
	bookRoute.Get("/read", controller.Show)
	bookRoute.Post("/create", middleware.BookValidator(), controller.CreateBook)
}
