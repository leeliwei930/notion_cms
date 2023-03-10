package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leeliwei930/notion_cms/internal/api/resource"
)

func GetMilestones(ctx *fiber.Ctx) error {
	data, err := resource.GetMilestonesExperiencesResource()
	if err != nil {
		return err
	}
	return ctx.JSON(map[string]interface{}{
		"data": data,
	})
}
