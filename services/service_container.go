package services

import (
	services_domain "github.com/RouteHub-Link/routehub-service-graphql/services/domain"
	services_link "github.com/RouteHub-Link/routehub-service-graphql/services/link"
	services_organization "github.com/RouteHub-Link/routehub-service-graphql/services/organization"
	services_platform "github.com/RouteHub-Link/routehub-service-graphql/services/platform"
	services_user "github.com/RouteHub-Link/routehub-service-graphql/services/user"
	"gorm.io/gorm"
)

type ServiceContainer struct {
	UserService         *services_user.UserService
	DomainService       *services_domain.DomainService
	PlatformService     *services_platform.PlatformService
	LinkService         *services_link.LinkService
	OrganizationService *services_organization.OrganizationService
}

func NewServiceContainer(db *gorm.DB) *ServiceContainer {
	return &ServiceContainer{
		UserService:         &services_user.UserService{DB: db},
		DomainService:       &services_domain.DomainService{DB: db},
		PlatformService:     &services_platform.PlatformService{DB: db},
		LinkService:         &services_link.LinkService{DB: db},
		OrganizationService: &services_organization.OrganizationService{DB: db},
	}
}
