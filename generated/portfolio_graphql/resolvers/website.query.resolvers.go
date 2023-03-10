package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"

	generated "github.com/leeliwei930/notion_cms/generated/portfolio_graphql"
	"github.com/leeliwei930/notion_cms/generated/portfolio_graphql/models"
	"github.com/leeliwei930/notion_cms/internal/api/resource"
)

// Config is the resolver for the config field.
func (r *websiteQueryResolver) Config(ctx context.Context, obj *models.WebsiteQuery) (*models.PageConfigurationPayload, error) {
	return resource.GetDefaultPageResource()
}

// WebsiteQuery returns generated.WebsiteQueryResolver implementation.
func (r *Resolver) WebsiteQuery() generated.WebsiteQueryResolver { return &websiteQueryResolver{r} }

type websiteQueryResolver struct{ *Resolver }