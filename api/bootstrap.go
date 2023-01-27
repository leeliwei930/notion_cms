package api

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/leeliwei930/notion_cms/api/routes"
)

var server *fiber.App

func Start(config ...fiber.Config) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	addr := fmt.Sprintf(":%s", os.Getenv("NOTION_CMS_PORT"))
	server := fiber.New(config...)

	routes.RegisterRoutes(server)
	server.Listen(addr)

}
