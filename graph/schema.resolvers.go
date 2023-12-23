package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"fmt"
	"time"

	"github.com/RouteHub-Link/routehub-service-graphql/auth"
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	database_relations "github.com/RouteHub-Link/routehub-service-graphql/database/relations"
	database_types "github.com/RouteHub-Link/routehub-service-graphql/database/types"
	"github.com/RouteHub-Link/routehub-service-graphql/graph/model"
	graph_inputs "github.com/RouteHub-Link/routehub-service-graphql/graph/model/inputs"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Organization is the resolver for the organization field.
func (r *domainResolver) Organization(ctx context.Context, obj *database.Domain) (*database.Organization, error) {
	panic(fmt.Errorf("not implemented: Organization - organization"))
}

// Platform is the resolver for the platform field.
func (r *domainResolver) Platform(ctx context.Context, obj *database.Domain) (*database.Platform, error) {
	panic(fmt.Errorf("not implemented: Platform - platform"))
}

// Verification is the resolver for the verification field.
func (r *domainResolver) Verification(ctx context.Context, obj *database.Domain) ([]*model.DomainVerification, error) {
	panic(fmt.Errorf("not implemented: Verification - verification"))
}

// State is the resolver for the state field.
func (r *domainResolver) State(ctx context.Context, obj *database.Domain) (database_enums.StatusState, error) {
	panic(fmt.Errorf("not implemented: State - state"))
}

// Analytics is the resolver for the analytics field.
func (r *domainResolver) Analytics(ctx context.Context, obj *database.Domain) ([]*model.MetricAnalytics, error) {
	panic(fmt.Errorf("not implemented: Analytics - analytics"))
}

// AnalyticReports is the resolver for the analyticReports field.
func (r *domainResolver) AnalyticReports(ctx context.Context, obj *database.Domain) (*model.AnalyticReports, error) {
	panic(fmt.Errorf("not implemented: AnalyticReports - analyticReports"))
}

// LastDNSVerificationAt is the resolver for the lastDNSVerificationAt field.
func (r *domainResolver) LastDNSVerificationAt(ctx context.Context, obj *database.Domain) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: LastDNSVerificationAt - lastDNSVerificationAt"))
}

// Organizations is the resolver for the Organizations field.
func (r *industryResolver) Organizations(ctx context.Context, obj *database_types.Industry) ([]*database.Organization, error) {
	panic(fmt.Errorf("not implemented: Organizations - Organizations"))
}

// Creator is the resolver for the creator field.
func (r *linkResolver) Creator(ctx context.Context, obj *database.Link) (*database.User, error) {
	return r.ServiceContainer.UserService.User(obj.CreatedBy)
}

// Platform is the resolver for the platform field.
func (r *linkResolver) Platform(ctx context.Context, obj *database.Link) (*database.Platform, error) {
	return r.ServiceContainer.PlatformService.GetPlatform(obj.PlatformID)
}

// Domain is the resolver for the domain field.
func (r *linkResolver) Domain(ctx context.Context, obj *database.Link) (*database.Domain, error) {
	return r.ServiceContainer.DomainService.GetDomainByPlatformId(obj.PlatformID)
}

// Analytics is the resolver for the analytics field.
func (r *linkResolver) Analytics(ctx context.Context, obj *database.Link) ([]*model.MetricAnalytics, error) {
	panic(fmt.Errorf("not implemented: Analytics - analytics"))
}

// OpenGraph is the resolver for the openGraph field.
func (r *linkResolver) OpenGraph(ctx context.Context, obj *database.Link) ([]*database_types.OpenGraph, error) {
	return []*database_types.OpenGraph{obj.OpenGraph}, nil
}

// RedirectionOptions is the resolver for the redirectionOptions field.
func (r *linkResolver) RedirectionOptions(ctx context.Context, obj *database.Link) (database_enums.RedirectionOptions, error) {
	return obj.RedirectionChoice, nil
}

// State is the resolver for the state field.
func (r *linkResolver) State(ctx context.Context, obj *database.Link) (database_enums.StatusState, error) {
	return obj.Status, nil
}

// Crawls is the resolver for the crawls field.
func (r *linkResolver) Crawls(ctx context.Context, obj *database.Link) ([]*database.LinkCrawl, error) {
	return r.ServiceContainer.LinkService.GetCrawls(obj.ID)
}

// Link is the resolver for the link field.
func (r *linkCrawlResolver) Link(ctx context.Context, obj *database.LinkCrawl) (*database.Link, error) {
	return r.ServiceContainer.LinkService.GetLinkById(obj.LinkId)
}

// CrawledBy is the resolver for the crawledBy field.
func (r *linkCrawlResolver) CrawledBy(ctx context.Context, obj *database.LinkCrawl) (*database.User, error) {
	return r.ServiceContainer.UserService.User(obj.CreatedBy)
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*database.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// LoginUser is the resolver for the loginUser field.
func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginInput) (*model.LoginPayload, error) {
	userService := r.ServiceContainer.UserService

	user, err := userService.Login(input)
	if err != nil {
		return nil, err
	}

	userSession := new(auth.UserSession)
	userSession.ID = user.ID
	userSession.Name = user.Fullname

	token, err := auth.GenerateToken(userSession.ToClaims())

	if err != nil {
		return nil, err
	}

	return &model.LoginPayload{
		Token: token,
	}, nil
}

// InviteUser is the resolver for the inviteUser field.
func (r *mutationResolver) InviteUser(ctx context.Context, input graph_inputs.UserInviteInput) (*database_relations.UserInvite, error) {
	userSession := auth.ForContext(ctx)
	userService := r.ServiceContainer.UserService

	invite, err := userService.InviteUser(input, userSession.ID)

	return invite, err
}

// UpdateUserInvitation is the resolver for the updateUserInvitation field.
func (r *mutationResolver) UpdateUserInvitation(ctx context.Context, input model.UpdateUserInviteInput) (database_enums.InvitationStatus, error) {
	userService := r.ServiceContainer.UserService
	invitation, err := userService.UpdateInvitation(input)

	if err != nil {
		return "", err
	}

	return invitation.Status, nil
}

// UpdateUserPassword is the resolver for the updateUserPassword field.
func (r *mutationResolver) UpdateUserPassword(ctx context.Context, userID string, input model.UserUpdatePasswordInput) (*database.User, error) {
	panic(fmt.Errorf("not implemented: UpdateUserPassword - updateUserPassword"))
}

// RequestPasswordReset is the resolver for the requestPasswordReset field.
func (r *mutationResolver) RequestPasswordReset(ctx context.Context, input model.PasswordResetCreateInput) (*model.PasswordReset, error) {
	panic(fmt.Errorf("not implemented: RequestPasswordReset - requestPasswordReset"))
}

// ResetPassword is the resolver for the resetPassword field.
func (r *mutationResolver) ResetPassword(ctx context.Context, input model.PasswordResetUpdateInput) (*database.User, error) {
	panic(fmt.Errorf("not implemented: ResetPassword - resetPassword"))
}

// CreateDomain is the resolver for the createDomain field.
func (r *mutationResolver) CreateDomain(ctx context.Context, input model.DomainCreateInput) (*database.Domain, error) {
	domainService := r.ServiceContainer.DomainService
	domain, err := domainService.CreateDomain(input)

	return &domain, err
}

// CreatePlatform is the resolver for the createPlatform field.
func (r *mutationResolver) CreatePlatform(ctx context.Context, input graph_inputs.PlatformCreateInput) (*database.Platform, error) {
	platformService := r.ServiceContainer.PlatformService
	userSession := auth.ForContext(ctx)

	platform, err := platformService.CreatePlatform(input, userSession.ID)
	return &platform, err
}

// CreateLink is the resolver for the createLink field.
func (r *mutationResolver) CreateLink(ctx context.Context, input model.LinkCreateInput) (*database.Link, error) {
	userSession := auth.ForContext(ctx)
	if userSession == nil {
		return nil, gqlerror.Errorf("Access Denied")
	}

	return r.ServiceContainer.LinkService.CreateLink(input, userSession.ID)
}

// RequestCrawl is the resolver for the requestCrawl field.
func (r *mutationResolver) RequestCrawl(ctx context.Context, input model.CrawlRequestInput) (*database.LinkCrawl, error) {
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

// Permissions is the resolver for the permissions field.
func (r *organizationResolver) Permissions(ctx context.Context, obj *database.Organization) ([]database_enums.OrganizationPermission, error) {
	panic(fmt.Errorf("not implemented: Permissions - permissions"))
}

// Platforms is the resolver for the platforms field.
func (r *organizationResolver) Platforms(ctx context.Context, obj *database.Organization) ([]*database.Platform, error) {
	return r.ServiceContainer.PlatformService.GetPlatformsByOrganization(obj.ID)
}

// Users is the resolver for the users field.
func (r *organizationResolver) Users(ctx context.Context, obj *database.Organization) ([]*database.User, error) {
	return r.ServiceContainer.UserService.UsersByOrganization(obj.ID)
}

// Domains is the resolver for the domains field.
func (r *organizationResolver) Domains(ctx context.Context, obj *database.Organization) ([]*database.Domain, error) {
	domainService := r.ServiceContainer.DomainService
	return domainService.GetDomainsByOrganization(obj.ID)
}

// PaymentPlan is the resolver for the paymentPlan field.
func (r *organizationResolver) PaymentPlan(ctx context.Context, obj *database.Organization) (model.PaymentPlan, error) {
	panic(fmt.Errorf("not implemented: PaymentPlan - paymentPlan"))
}

// Payments is the resolver for the payments field.
func (r *organizationResolver) Payments(ctx context.Context, obj *database.Organization) ([]*model.Payment, error) {
	panic(fmt.Errorf("not implemented: Payments - payments"))
}

// Organization is the resolver for the organization field.
func (r *platformResolver) Organization(ctx context.Context, obj *database.Platform) (*database.Organization, error) {
	panic(fmt.Errorf("not implemented: Organization - organization"))
}

// Domain is the resolver for the domain field.
func (r *platformResolver) Domain(ctx context.Context, obj *database.Platform) (*database.Domain, error) {
	domainService := r.ServiceContainer.DomainService
	return domainService.GetDomain(obj.DomainId)
}

// Permissions is the resolver for the permissions field.
func (r *platformResolver) Permissions(ctx context.Context, obj *database.Platform) ([]database_enums.PlatformPermission, error) {
	panic(fmt.Errorf("not implemented: Permissions - permissions"))
}

// Deployments is the resolver for the deployments field.
func (r *platformResolver) Deployments(ctx context.Context, obj *database.Platform) ([]*model.PlatformDeployment, error) {
	panic(fmt.Errorf("not implemented: Deployments - deployments"))
}

// Links is the resolver for the links field.
func (r *platformResolver) Links(ctx context.Context, obj *database.Platform) ([]*database.Link, error) {
	return r.ServiceContainer.LinkService.GetLinksByPlatformId(obj.ID)
}

// Analytics is the resolver for the analytics field.
func (r *platformResolver) Analytics(ctx context.Context, obj *database.Platform) ([]*model.AnalyticReport, error) {
	panic(fmt.Errorf("not implemented: Analytics - analytics"))
}

// AnalyticReports is the resolver for the analyticReports field.
func (r *platformResolver) AnalyticReports(ctx context.Context, obj *database.Platform) (*model.AnalyticReports, error) {
	panic(fmt.Errorf("not implemented: AnalyticReports - analyticReports"))
}

// Templates is the resolver for the templates field.
func (r *platformResolver) Templates(ctx context.Context, obj *database.Platform) ([]*model.Template, error) {
	panic(fmt.Errorf("not implemented: Templates - templates"))
}

// PinnedLinks is the resolver for the pinnedLinks field.
func (r *platformResolver) PinnedLinks(ctx context.Context, obj *database.Platform) ([]*database.Link, error) {
	panic(fmt.Errorf("not implemented: PinnedLinks - pinnedLinks"))
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*database.User, error) {
	userSession := auth.ForContext(ctx)
	if userSession == nil {
		return nil, gqlerror.Errorf("Access Denied")
	}

	userService := r.ServiceContainer.UserService

	return userService.User(userSession.ID)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*database.User, error) {
	return r.ServiceContainer.UserService.Users()
}

// Organizations is the resolver for the organizations field.
func (r *queryResolver) Organizations(ctx context.Context) ([]*database.Organization, error) {
	panic(fmt.Errorf("not implemented: Organizations - organizations"))
}

// Platforms is the resolver for the platforms field.
func (r *queryResolver) Platforms(ctx context.Context) ([]*database.Platform, error) {
	return r.ServiceContainer.PlatformService.GetPlatforms()
}

// RefreshToken is the resolver for the refreshToken field.
func (r *queryResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented: RefreshToken - refreshToken"))
}

// Organizations is the resolver for the Organizations field.
func (r *userResolver) Organizations(ctx context.Context, obj *database.User) ([]*database.Organization, error) {
	return r.ServiceContainer.UserService.OrganizationUser(obj.ID)
}

// Platforms is the resolver for the platforms field.
func (r *userResolver) Platforms(ctx context.Context, obj *database.User) ([]*database.Platform, error) {
	return r.ServiceContainer.PlatformService.GetPlatformsByUser(obj.ID)
}

// Permissions is the resolver for the permissions field.
func (r *userResolver) Permissions(ctx context.Context, obj *database.User) ([]*model.Permission, error) {
	panic(fmt.Errorf("not implemented: Permissions - permissions"))
}

// Payments is the resolver for the payments field.
func (r *userResolver) Payments(ctx context.Context, obj *database.User) ([]*model.Payment, error) {
	panic(fmt.Errorf("not implemented: Payments - payments"))
}

// UserInvites is the resolver for the userInvites field.
func (r *userResolver) UserInvites(ctx context.Context, obj *database.User) ([]*database_relations.UserInvite, error) {
	panic(fmt.Errorf("not implemented: UserInvites - userInvites"))
}

// Organization is the resolver for the organization field.
func (r *userInviteResolver) Organization(ctx context.Context, obj *database_relations.UserInvite) (*database.Organization, error) {
	panic(fmt.Errorf("not implemented: Organization - organization"))
}

// Platforms is the resolver for the platforms field.
func (r *userInviteResolver) Platforms(ctx context.Context, obj *database_relations.UserInvite) ([]*database.Platform, error) {
	panic(fmt.Errorf("not implemented: Platforms - platforms"))
}

// User is the resolver for the user field.
func (r *userInviteResolver) User(ctx context.Context, obj *database_relations.UserInvite) (*database.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// DeletedAt is the resolver for the deletedAt field.
func (r *userInviteResolver) DeletedAt(ctx context.Context, obj *database_relations.UserInvite) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: DeletedAt - deletedAt"))
}

// Permissions is the resolver for the permissions field.
func (r *organizationsWithPermissionsInputResolver) Permissions(ctx context.Context, obj *database_relations.OrganizationsWithPermissions, data []database_enums.OrganizationPermission) error {
	// TODO check the organizations i guess user has permisison or just remove this
	return nil
}

// Permissions is the resolver for the permissions field.
func (r *platformsWithPermissionsInputResolver) Permissions(ctx context.Context, obj *database_relations.PlatformsWithPermissions, data []database_enums.PlatformPermission) error {
	// TODO check the platform i guess user has permisison or just remove this
	return nil
}

// Domain returns DomainResolver implementation.
func (r *Resolver) Domain() DomainResolver { return &domainResolver{r} }

// Industry returns IndustryResolver implementation.
func (r *Resolver) Industry() IndustryResolver { return &industryResolver{r} }

// Link returns LinkResolver implementation.
func (r *Resolver) Link() LinkResolver { return &linkResolver{r} }

// LinkCrawl returns LinkCrawlResolver implementation.
func (r *Resolver) LinkCrawl() LinkCrawlResolver { return &linkCrawlResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Organization returns OrganizationResolver implementation.
func (r *Resolver) Organization() OrganizationResolver { return &organizationResolver{r} }

// Platform returns PlatformResolver implementation.
func (r *Resolver) Platform() PlatformResolver { return &platformResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

// UserInvite returns UserInviteResolver implementation.
func (r *Resolver) UserInvite() UserInviteResolver { return &userInviteResolver{r} }

// OrganizationsWithPermissionsInput returns OrganizationsWithPermissionsInputResolver implementation.
func (r *Resolver) OrganizationsWithPermissionsInput() OrganizationsWithPermissionsInputResolver {
	return &organizationsWithPermissionsInputResolver{r}
}

// PlatformsWithPermissionsInput returns PlatformsWithPermissionsInputResolver implementation.
func (r *Resolver) PlatformsWithPermissionsInput() PlatformsWithPermissionsInputResolver {
	return &platformsWithPermissionsInputResolver{r}
}

type domainResolver struct{ *Resolver }
type industryResolver struct{ *Resolver }
type linkResolver struct{ *Resolver }
type linkCrawlResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type organizationResolver struct{ *Resolver }
type platformResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type userInviteResolver struct{ *Resolver }
type organizationsWithPermissionsInputResolver struct{ *Resolver }
type platformsWithPermissionsInputResolver struct{ *Resolver }
