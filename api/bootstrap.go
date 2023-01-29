package api

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
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
	server.Use(requestid.New())
	server.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))
	routes.RegisterRoutes(server)
	log.Fatal(server.Listen(addr))

}
