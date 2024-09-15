package main

import (
	auth_casbin "github.com/RouteHub-Link/routehub-service-graphql/auth/casbin"
	Configuration "github.com/RouteHub-Link/routehub-service-graphql/config"
	"github.com/RouteHub-Link/routehub-service-graphql/database"
)

var applicationConfig = Configuration.ConfigurationService{}.Get()

func main() {
	database.Init()
	auth_casbin.NewCasbinConfigurer(applicationConfig.CasbinConfig, database.DB)

	database.MigrateAndSeed()

	Serve()
}
