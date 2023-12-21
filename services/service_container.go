package services

import (
	services_domain "github.com/RouteHub-Link/routehub-service-graphql/services/domain"
	services_user "github.com/RouteHub-Link/routehub-service-graphql/services/user"
	"gorm.io/gorm"
)

type ServiceContainer struct {
	UserService   *services_user.UserService
	DomainService *services_domain.DomainService
}

func NewServiceContainer(db *gorm.DB) *ServiceContainer {
	return &ServiceContainer{
		UserService:   &services_user.UserService{DB: db},
		DomainService: &services_domain.DomainService{DB: db},
	}
}
