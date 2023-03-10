package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leeliwei930/notion_cms/internal/api/resource"
)

func GetEducationPathwayData(ctx *fiber.Ctx) error {

	data, err := resource.GetEducationPathwayResource()
	if err != nil {
		return err
	}
	return ctx.JSON(map[string]interface{}{
		"data": data,
	})
}
