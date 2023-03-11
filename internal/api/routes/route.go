package routes

import (
	"context"
	"errors"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/leeliwei930/notion_cms/generated/portfolio_graphql"
	"github.com/leeliwei930/notion_cms/generated/portfolio_graphql/resolvers"
	"github.com/leeliwei930/notion_cms/internal/api/controllers"
	"github.com/leeliwei930/notion_sdk/models"
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
	graphQLSrv.SetErrorPresenter(
		func(ctx context.Context, e error) *gqlerror.Error {
			err := graphql.DefaultErrorPresenter(ctx, e)

			var notionError *models.NotionError
			if errors.As(e, &notionError) {
				err.Message = notionError.Error()
				err.Extensions = map[string]interface{}{
					"notion": map[string]interface{}{
						"error_code": notionError.Code,
						"messages":   strings.Split(notionError.Message, "\n"),
					},
				}
			}
			return err
		})

	graphQL.Get("/playground", adaptor.HTTPHandlerFunc(playground.Handler("NotionCMS GraphQL playground", "/graphql/query")))
	graphQL.Post("/query", adaptor.HTTPHandler(graphQLSrv))
}
