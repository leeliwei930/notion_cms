package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leeliwei930/notion_cms/api/resource"
)

func GetEducationPathwayData(ctx *fiber.Ctx) error {

	data, err := resource.GetEducationPathwayResource()
	if err != nil {
		return ctx.JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}
	return ctx.JSON(map[string]interface{}{
		"data": data,
	})
}
