package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leeliwei930/notion_cms/api/controllers"
)

// Register your routes overhere
func RegisterRoutes(app *fiber.App) {
	baseApi := app.Group("/api")
	v1 := baseApi.Group("/v1")

	educationResourceApi := v1.Group("/education")
	educationResourceApi.Get("/", controllers.GetEducationPathwayData)
}
