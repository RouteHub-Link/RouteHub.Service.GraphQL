package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.46

import (
	"context"
	"fmt"
	"time"

	"github.com/RouteHub-Link/routehub-service-graphql/auth"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	"github.com/RouteHub-Link/routehub-service-graphql/graph"
	"github.com/RouteHub-Link/routehub-service-graphql/graph/model"
	"github.com/google/uuid"
)

// CreatedBy is the resolver for the createdBy field.
func (r *dNSVerificationResolver) CreatedBy(ctx context.Context, obj *database_models.DNSVerification) (*database_models.User, error) {
	return r.LoaderContainer.User.Get(ctx, obj.CreatedBy)
}

// Organization is the resolver for the organization field.
func (r *domainResolver) Organization(ctx context.Context, obj *database_models.Domain) (*database_models.Organization, error) {
	return r.ServiceContainer.OrganizationService.GetOrganization(obj.OrganizationId)
}

// Platform is the resolver for the platform field.
func (r *domainResolver) Platform(ctx context.Context, obj *database_models.Domain) (*database_models.Platform, error) {
	return r.ServiceContainer.PlatformService.GetPlatformByDomainId(obj.ID)
}

// Verifications is the resolver for the verifications field.
func (r *domainResolver) Verifications(ctx context.Context, obj *database_models.Domain) ([]*database_models.DNSVerification, error) {
	return r.ServiceContainer.DNSVerificationService.GetDNSVerificationsByDomainId(obj.ID)
}

// LastVerification is the resolver for the lastVerification field.
func (r *domainResolver) LastVerification(ctx context.Context, obj *database_models.Domain) (*database_models.DNSVerification, error) {
	return r.ServiceContainer.DNSVerificationService.GetDNSVerificationByDomainId(obj.ID)
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
	lastVerification, err := r.ServiceContainer.DNSVerificationService.GetDNSVerificationByDomainId(obj.ID)

	if err != nil {
		return nil, err
	}

	if lastVerification == nil {
		return nil, nil
	}

	return lastVerification.CompletedAt, nil
}

// CreateDomain is the resolver for the createDomain field.
func (r *mutationResolver) CreateDomain(ctx context.Context, input model.DomainCreateInput) (*database_models.Domain, error) {
	userSession := auth.ForContext(ctx)

	domainService := r.ServiceContainer.DomainService
	domain, err := domainService.CreateDomain(input)

	if err != nil {
		return nil, err
	}

	r.ServiceContainer.DNSVerificationService.Validate(userSession.ID, &domain, true)

	return &domain, err
}

// NewDomainVerification is the resolver for the newDomainVerification field.
func (r *mutationResolver) NewDomainVerification(ctx context.Context, domainID uuid.UUID, forced *bool) (*database_models.DNSVerification, error) {
	userSession := auth.ForContext(ctx)

	domainService := r.ServiceContainer.DomainService
	domain, err := domainService.GetDomain(domainID)

	if err != nil {
		return nil, err
	}

	_forced := false
	if forced != nil {
		_forced = *forced
	}

	dnsVerification, err := r.ServiceContainer.DNSVerificationService.Validate(userSession.ID, domain, _forced)

	return dnsVerification, err
}

// DNSVerification returns graph.DNSVerificationResolver implementation.
func (r *Resolver) DNSVerification() graph.DNSVerificationResolver {
	return &dNSVerificationResolver{r}
}

// Domain returns graph.DomainResolver implementation.
func (r *Resolver) Domain() graph.DomainResolver { return &domainResolver{r} }

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

type dNSVerificationResolver struct{ *Resolver }
type domainResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
