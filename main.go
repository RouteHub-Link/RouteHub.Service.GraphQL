package main

import (
	"github.com/RouteHub-Link/routehub-service-graphql/database"
	database_mock "github.com/RouteHub-Link/routehub-service-graphql/database/mock"
)

func main() {

	database.RunEmbeddedPostgres()
	go database.InterruptEmbedded()
	database.Init()
	database_mock.Seed()

	Serve()
}
