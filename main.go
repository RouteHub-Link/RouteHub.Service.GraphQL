package main

import (
	"github.com/RouteHub-Link/routehub-service-graphql/auth"
	Configuration "github.com/RouteHub-Link/routehub-service-graphql/config"
	"github.com/RouteHub-Link/routehub-service-graphql/database"
)

var applicationConfig = Configuration.ConfigurationService{}.Get()

func main() {
	auth.NewCasbinConfigurer(applicationConfig.CasbinConfig)

	database.Init()

	Serve()
}
