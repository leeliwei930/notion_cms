package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leeliwei930/notion_cms/api/controllers"
)

// Register your routes overhere
func RegisterRoutes(app *fiber.App) {
	app.Get("/", controllers.GetEducationPathwayData)
}
