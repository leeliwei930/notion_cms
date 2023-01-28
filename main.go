package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leeliwei930/notion_cms/api"
	errorsHandler "github.com/leeliwei930/notion_cms/api/controllers/errors"
)

func main() {
	api.Start(fiber.Config{
		ErrorHandler: errorsHandler.ApiErrorHandler,
	})
}
