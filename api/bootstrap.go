package api

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/leeliwei930/notion_cms/api/routes"
	"github.com/leeliwei930/notion_sdk/client"
	notionConfig "github.com/leeliwei930/notion_sdk/config"
)

func Start(config ...fiber.Config) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	client.InitializeNotionConfig(&notionConfig.NotionConfig{
		AccessToken:   os.Getenv("NOTION_ACCESS_TOKEN"),
		NotionVersion: "2022-06-28",
	})
	addr := fmt.Sprintf(":%s", os.Getenv("NOTION_CMS_PORT"))
	server := fiber.New(config...)
	server.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	routes.RegisterRoutes(server)
	server.Listen(addr)

}
