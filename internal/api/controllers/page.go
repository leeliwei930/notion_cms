package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leeliwei930/notion_cms/internal/api/resource"
)

func GetDefaultPageConfig(ctx *fiber.Ctx) error {
	data, err := resource.GetDefaultPageResource()
	if err != nil {
		return err
	}
	return ctx.JSON(map[string]interface{}{
		"data": data,
	})
}
