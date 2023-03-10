package routes

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	// "github.com/leeliwei930/notion_cms/graph"
	"github.com/leeliwei930/notion_cms/generated/portfolio_graphql"
	"github.com/leeliwei930/notion_cms/generated/portfolio_graphql/resolvers"
	"github.com/leeliwei930/notion_cms/internal/api/controllers"
)

// Register your routes overhere
func RegisterRoutes(app *fiber.App) {
	app.Get("/metrics", monitor.New())
	graphQL := app.Group("/graphql")
	baseApi := app.Group("/api")
	v1 := baseApi.Group("/v1")

	educationResource := v1.Group("/education")
	educationResource.Get("/", controllers.GetEducationPathwayData)

	pageResource := v1.Group("/page")
	pageResource.Get("/", controllers.GetDefaultPageConfig)

	milestoneResource := v1.Group("/milestone")
	milestoneResource.Get("/", controllers.GetMilestones)
	graphQLSrv := handler.NewDefaultServer(portfolio_graphql.NewExecutableSchema(
		portfolio_graphql.Config{
			Resolvers: &resolvers.Resolver{},
		},
	))

	graphQL.Get("/playground", adaptor.HTTPHandlerFunc(playground.Handler("NotionCMS GraphQL playground", "/graphql/query")))
	graphQL.Post("/query", adaptor.HTTPHandler(graphQLSrv))
}
