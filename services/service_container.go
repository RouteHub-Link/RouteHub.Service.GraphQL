package services

import (
	services_domain "github.com/RouteHub-Link/routehub-service-graphql/services/domain"
	services_domain_utils "github.com/RouteHub-Link/routehub-service-graphql/services/domain_utils"
	services_link "github.com/RouteHub-Link/routehub-service-graphql/services/link"
	services_organization "github.com/RouteHub-Link/routehub-service-graphql/services/organization"
	services_platform "github.com/RouteHub-Link/routehub-service-graphql/services/platform"
	services_user "github.com/RouteHub-Link/routehub-service-graphql/services/user"
	"gorm.io/gorm"
)

type ServiceContainer struct {
	UserService                   *services_user.UserService
	DomainService                 *services_domain.DomainService
	DomainUtilsService            *services_domain_utils.DomainUtilsService
	PlatformService               *services_platform.PlatformService
	LinkService                   *services_link.LinkService
	LinkValidationService         *services_link.LinkValidationService
	DNSVerificationService        *services_domain.DNSVerificationService
	OrganizationService           *services_organization.OrganizationService
	PlatformPermissionService     *services_platform.PlatformPermissionService
	OrganizationPermissionService *services_organization.OrganizationPermissionService
}

func NewServiceContainer(db *gorm.DB) *ServiceContainer {
	domainUtilsService := *services_domain_utils.NewDomainUtilsService()

	return &ServiceContainer{
		UserService:                   &services_user.UserService{DB: db},
		DomainService:                 &services_domain.DomainService{DB: db},
		DomainUtilsService:            &domainUtilsService,
		PlatformService:               &services_platform.PlatformService{DB: db},
		LinkService:                   &services_link.LinkService{DB: db},
		LinkValidationService:         services_link.NewLinkValidationService(&domainUtilsService, db),
		DNSVerificationService:        services_domain.NewDNSVerificationService(&domainUtilsService, db),
		OrganizationService:           &services_organization.OrganizationService{DB: db},
		PlatformPermissionService:     &services_platform.PlatformPermissionService{DB: db},
		OrganizationPermissionService: &services_organization.OrganizationPermissionService{DB: db},
	}
}
