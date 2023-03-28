package router

import (
	"idkanymore/handler"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/greeting", handler.HiApaKhabar)

}
