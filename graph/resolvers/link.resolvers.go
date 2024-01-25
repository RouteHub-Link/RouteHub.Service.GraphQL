package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"time"

	"github.com/RouteHub-Link/routehub-service-graphql/auth"
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	database_types "github.com/RouteHub-Link/routehub-service-graphql/database/types"
	"github.com/RouteHub-Link/routehub-service-graphql/graph"
	"github.com/RouteHub-Link/routehub-service-graphql/graph/model"
	"github.com/cloudmatelabs/gorm-gqlgen-relay/relay"
	"github.com/vektah/gqlparser/gqlerror"
)

// Creator is the resolver for the creator field.
func (r *linkResolver) Creator(ctx context.Context, obj *database_models.Link) (*database_models.User, error) {
	return r.LoaderContainer.User.Get(ctx, obj.CreatedBy)
}

// Domain is the resolver for the domain field.
func (r *linkResolver) Domain(ctx context.Context, obj *database_models.Link) (*database_models.Domain, error) {
	return r.ServiceContainer.DomainService.GetDomainByPlatformId(obj.PlatformID)
}

// Analytics is the resolver for the analytics field.
func (r *linkResolver) Analytics(ctx context.Context, obj *database_models.Link) ([]*model.MetricAnalytics, error) {
	mock := []*model.MetricAnalytics{
		{
			Feeder:       "Not Implemented (Mock)",
			TotalHits:    300,
			TotalSuccess: 123,
			TotalFailed:  177,
			StartAt:      time.Now().Add(-time.Hour * (24 * 3)),
			EndAt:        time.Now(),
		},
		{
			Feeder:       "Facebook (Mock)",
			TotalHits:    100,
			TotalSuccess: 50,
			TotalFailed:  50,
			StartAt:      time.Now().Add(-time.Hour * (24 * 7)),
			EndAt:        time.Now(),
		},
		{
			Feeder:       "X (Mock)",
			TotalHits:    300,
			TotalSuccess: 200,
			TotalFailed:  100,
			StartAt:      time.Now().Add(-time.Hour * (12 * 7)),
			EndAt:        time.Now(),
		}}
	return mock, nil
}

// OpenGraph is the resolver for the openGraph field.
func (r *linkResolver) OpenGraph(ctx context.Context, obj *database_models.Link) ([]*database_types.OpenGraph, error) {
	return []*database_types.OpenGraph{obj.OpenGraph}, nil
}

// RedirectionOptions is the resolver for the redirectionOptions field.
func (r *linkResolver) RedirectionOptions(ctx context.Context, obj *database_models.Link) (database_enums.RedirectionOptions, error) {
	return obj.RedirectionChoice, nil
}

// Crawls is the resolver for the crawls field.
func (r *linkResolver) Crawls(ctx context.Context, obj *database_models.Link) ([]*database_models.LinkCrawl, error) {
	return r.ServiceContainer.LinkService.GetCrawls(obj.ID)
}

// Link is the resolver for the link field.
func (r *linkCrawlResolver) Link(ctx context.Context, obj *database_models.LinkCrawl) (*database_models.Link, error) {
	return r.ServiceContainer.LinkService.GetLinkById(obj.LinkId)
}

// CrawledBy is the resolver for the crawledBy field.
func (r *linkCrawlResolver) CrawledBy(ctx context.Context, obj *database_models.LinkCrawl) (*database_models.User, error) {
	return r.LoaderContainer.User.Get(ctx, obj.CreatedBy)
}

// CreateLink is the resolver for the createLink field.
func (r *mutationResolver) CreateLink(ctx context.Context, input model.LinkCreateInput) (*database_models.Link, error) {
	userSession := auth.ForContext(ctx)
	if userSession == nil {
		return nil, gqlerror.Errorf("Access Denied")
	}

	return r.ServiceContainer.LinkService.CreateLink(input, userSession.ID)
}

// RequestCrawl is the resolver for the requestCrawl field.
func (r *mutationResolver) RequestCrawl(ctx context.Context, input model.CrawlRequestInput) (*database_models.LinkCrawl, error) {
	userSession := auth.ForContext(ctx)
	if userSession == nil {
		return nil, gqlerror.Errorf("Access Denied")
	}

	link, err := r.ServiceContainer.LinkService.RequestCrawl(input.LinkID, userSession.ID)
	if err != nil {
		return nil, gqlerror.Errorf("Process Failed %s", err.Error())
	}

	crawls, err := r.ServiceContainer.LinkService.GetCrawls(link.ID)

	return crawls[len(crawls)-1], err
}

// Links is the resolver for the links field.
func (r *queryResolver) Links(ctx context.Context, first *int, after *string, last *int, before *string, orderBy map[string]interface{}, where *model.LinkFilter) (*relay.Connection[database_models.Link], error) {
	linkName := database_models.Link{}.TableName()

	// If the cursor is "<uuid.UUID Value>" Check this link ; https://github.com/cloudmatelabs/gorm-gqlgen-relay/pull/2
	return relay.Paginate[database_models.Link](r.DB, where, orderBy, relay.PaginateOption{
		First:      first,
		After:      after,
		Last:       last,
		Before:     before,
		Table:      linkName,
		PrimaryKey: "id",
	})
}

// Link returns graph.LinkResolver implementation.
func (r *Resolver) Link() graph.LinkResolver { return &linkResolver{r} }

// LinkCrawl returns graph.LinkCrawlResolver implementation.
func (r *Resolver) LinkCrawl() graph.LinkCrawlResolver { return &linkCrawlResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type linkResolver struct{ *Resolver }
type linkCrawlResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
