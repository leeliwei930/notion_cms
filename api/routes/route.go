package routes

import (
	"github.com/gofiber/fiber/v2"
)

// Register your routes overhere
func RegisterRoutes(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.
			JSON(map[string]interface{}{
				"message": "hello world",
			})
	})
}
