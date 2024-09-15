package main

import (
	auth_casbin "github.com/RouteHub-Link/routehub-service-graphql/auth/casbin"
	Configuration "github.com/RouteHub-Link/routehub-service-graphql/config"
	"github.com/RouteHub-Link/routehub-service-graphql/database"
)

var applicationConfig = Configuration.ConfigurationService{}.Get()

func main() {
	database.Init()
	database.Migration()

	auth_casbin.NewCasbinConfigurer(applicationConfig.CasbinConfig, applicationConfig.Database)

	database.Seed()

	Serve()
}
