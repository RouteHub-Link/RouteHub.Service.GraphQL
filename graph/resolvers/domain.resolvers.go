package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"fmt"
	"time"

	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	"github.com/RouteHub-Link/routehub-service-graphql/graph"
	"github.com/RouteHub-Link/routehub-service-graphql/graph/model"
)

// Organization is the resolver for the organization field.
func (r *domainResolver) Organization(ctx context.Context, obj *database_models.Domain) (*database_models.Organization, error) {
	return r.ServiceContainer.OrganizationService.GetOrganization(obj.OrganizationId)
}

// Platform is the resolver for the platform field.
func (r *domainResolver) Platform(ctx context.Context, obj *database_models.Domain) (*database_models.Platform, error) {
	return r.ServiceContainer.PlatformService.GetPlatformByDomainId(obj.ID)
}

// Verification is the resolver for the verification field.
func (r *domainResolver) Verification(ctx context.Context, obj *database_models.Domain) ([]*model.DomainVerification, error) {
	panic(fmt.Errorf("not implemented: Verification - verification"))
}

// Analytics is the resolver for the analytics field.
func (r *domainResolver) Analytics(ctx context.Context, obj *database_models.Domain) ([]*model.MetricAnalytics, error) {
	panic(fmt.Errorf("not implemented: Analytics - analytics"))
}

// AnalyticReports is the resolver for the analyticReports field.
func (r *domainResolver) AnalyticReports(ctx context.Context, obj *database_models.Domain) (*model.AnalyticReports, error) {
	panic(fmt.Errorf("not implemented: AnalyticReports - analyticReports"))
}

// LastDNSVerificationAt is the resolver for the lastDNSVerificationAt field.
func (r *domainResolver) LastDNSVerificationAt(ctx context.Context, obj *database_models.Domain) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: LastDNSVerificationAt - lastDNSVerificationAt"))
}

// CreateDomain is the resolver for the createDomain field.
func (r *mutationResolver) CreateDomain(ctx context.Context, input model.DomainCreateInput) (*database_models.Domain, error) {
	domainService := r.ServiceContainer.DomainService
	domain, err := domainService.CreateDomain(input)

	return &domain, err
}

// Domain returns graph.DomainResolver implementation.
func (r *Resolver) Domain() graph.DomainResolver { return &domainResolver{r} }

type domainResolver struct{ *Resolver }