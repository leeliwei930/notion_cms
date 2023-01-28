package api

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/leeliwei930/notion_cms/api/routes"
	"github.com/leeliwei930/notion_sdk/client"
	notionConfig "github.com/leeliwei930/notion_sdk/config"
)

var server *fiber.App

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

	routes.RegisterRoutes(server)
	server.Listen(addr)

}
