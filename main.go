package main

import (
	"github.com/RouteHub-Link/routehub-service-graphql/database"
	"github.com/RouteHub-Link/routehub-service-graphql/worker"
)

func main() {
	database.Init()
	go worker.Init()

	Serve()
}
