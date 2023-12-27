package graph

import (
	"github.com/RouteHub-Link/routehub-service-graphql/loaders"
	"github.com/RouteHub-Link/routehub-service-graphql/services"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB               *gorm.DB
	ServiceContainer *services.ServiceContainer
	LoaderContainer  *loaders.LoaderContainer
}
