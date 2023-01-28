package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leeliwei930/notion_cms/api/controllers"
)

// Register your routes overhere
func RegisterRoutes(app *fiber.App) {
	baseApi := app.Group("/api")
	v1 := baseApi.Group("/v1")

	educationResource := v1.Group("/education")
	educationResource.Get("/", controllers.GetEducationPathwayData)

	pageResource := v1.Group("/page")
	pageResource.Get("/", controllers.GetDefaultPageConfig)

}
