package routes

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/leeliwei930/notion_cms/api/controllers"
	"github.com/leeliwei930/notion_cms/graph"
)

// Register your routes overhere
func RegisterRoutes(app *fiber.App) {
	graphQL := app.Group("/graphql")
	baseApi := app.Group("/api")
	v1 := baseApi.Group("/v1")

	educationResource := v1.Group("/education")
	educationResource.Get("/", controllers.GetEducationPathwayData)

	pageResource := v1.Group("/page")
	pageResource.Get("/", controllers.GetDefaultPageConfig)

	graphQLSrv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	graphQL.Get("/playground", adaptor.HTTPHandlerFunc(playground.Handler("NotionCMS GraphQL playground", "/graphql/query")))
	graphQL.Post("/query", adaptor.HTTPHandler(graphQLSrv))
}
