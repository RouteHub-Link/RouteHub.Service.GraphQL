package services

import (
	services_domain "github.com/RouteHub-Link/routehub-service-graphql/services/domain"
	services_link "github.com/RouteHub-Link/routehub-service-graphql/services/link"
	services_organization "github.com/RouteHub-Link/routehub-service-graphql/services/organization"
	services_platform "github.com/RouteHub-Link/routehub-service-graphql/services/platform"
	services_user "github.com/RouteHub-Link/routehub-service-graphql/services/user"
	"github.com/RouteHub-Link/routehub-service-graphql/worker"
	"gorm.io/gorm"
)

type ServiceContainer struct {
	UserService                   *services_user.UserService
	DomainService                 *services_domain.DomainService
	PlatformService               *services_platform.PlatformService
	WorkerService                 *worker.WrokerService
	LinkService                   *services_link.LinkService
	OrganizationService           *services_organization.OrganizationService
	PlatformPermissionService     *services_platform.PlatformPermissionService
	OrganizationPermissionService *services_organization.OrganizationPermissionService
}

func NewServiceContainer(db *gorm.DB) *ServiceContainer {
	return &ServiceContainer{
		UserService:                   &services_user.UserService{DB: db},
		DomainService:                 &services_domain.DomainService{DB: db},
		PlatformService:               &services_platform.PlatformService{DB: db},
		WorkerService:                 &worker.WrokerService{},
		LinkService:                   &services_link.LinkService{DB: db},
		OrganizationService:           &services_organization.OrganizationService{DB: db},
		PlatformPermissionService:     &services_platform.PlatformPermissionService{DB: db},
		OrganizationPermissionService: &services_organization.OrganizationPermissionService{DB: db},
	}
}
