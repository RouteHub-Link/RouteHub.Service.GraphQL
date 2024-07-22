package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.46

import (
	"context"
	"fmt"

	"github.com/RouteHub-Link/routehub-service-graphql/auth"
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	"github.com/RouteHub-Link/routehub-service-graphql/graph"
	"github.com/RouteHub-Link/routehub-service-graphql/graph/model"
	graph_inputs "github.com/RouteHub-Link/routehub-service-graphql/graph/model/inputs"
)

// CreatePlatform is the resolver for the createPlatform field.
func (r *mutationResolver) CreatePlatform(ctx context.Context, input graph_inputs.PlatformCreateInput) (*database_models.Platform, error) {
	platformService := r.ServiceContainer.PlatformService
	userSession := auth.ForContext(ctx)

	platform, err := platformService.CreatePlatform(input, userSession.ID)
	return &platform, err
}

// Organization is the resolver for the organization field.
func (r *platformResolver) Organization(ctx context.Context, obj *database_models.Platform) (*database_models.Organization, error) {
	return r.ServiceContainer.PlatformService.GetPlatformOrganization(obj.ID)
}

// Domain is the resolver for the domain field.
func (r *platformResolver) Domain(ctx context.Context, obj *database_models.Platform) (*database_models.Domain, error) {
	domainService := r.ServiceContainer.DomainService
	return domainService.GetDomain(obj.DomainId)
}

// Permissions is the resolver for the permissions field.
func (r *platformResolver) Permissions(ctx context.Context, obj *database_models.Platform) ([]database_enums.PlatformPermission, error) {
	userSession := auth.ForContext(ctx)
	return r.ServiceContainer.PlatformPermissionService.GetPlatformPermissions(userSession.ID, obj.ID)
}

// Deployments is the resolver for the deployments field.
func (r *platformResolver) Deployments(ctx context.Context, obj *database_models.Platform) ([]*model.PlatformDeployment, error) {
	panic(fmt.Errorf("not implemented: Deployments - deployments"))
}

// Links is the resolver for the links field.
func (r *platformResolver) Links(ctx context.Context, obj *database_models.Platform) ([]*database_models.Link, error) {
	return r.ServiceContainer.LinkService.GetLinksByPlatformId(obj.ID)
}

// Analytics is the resolver for the analytics field.
func (r *platformResolver) Analytics(ctx context.Context, obj *database_models.Platform) ([]*model.AnalyticReport, error) {
	panic(fmt.Errorf("not implemented: Analytics - analytics"))
}

// AnalyticReports is the resolver for the analyticReports field.
func (r *platformResolver) AnalyticReports(ctx context.Context, obj *database_models.Platform) (*model.AnalyticReports, error) {
	panic(fmt.Errorf("not implemented: AnalyticReports - analyticReports"))
}

// Templates is the resolver for the templates field.
func (r *platformResolver) Templates(ctx context.Context, obj *database_models.Platform) ([]*model.Template, error) {
	panic(fmt.Errorf("not implemented: Templates - templates"))
}

// PinnedLinks is the resolver for the pinnedLinks field.
func (r *platformResolver) PinnedLinks(ctx context.Context, obj *database_models.Platform) ([]*database_models.Link, error) {
	pinnedLinks := obj.GetPinnedLinksAsIds()

	return r.ServiceContainer.LinkService.GetLinksByIds(pinnedLinks)
}

// Platforms is the resolver for the platforms field.
func (r *queryResolver) Platforms(ctx context.Context) ([]*database_models.Platform, error) {
	userSession := auth.ForContext(ctx)

	return r.ServiceContainer.PlatformService.GetPlatformsByUser(userSession.ID)
}

// Platform returns graph.PlatformResolver implementation.
func (r *Resolver) Platform() graph.PlatformResolver { return &platformResolver{r} }

type platformResolver struct{ *Resolver }
