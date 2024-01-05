package main

import (
	"github.com/RouteHub-Link/routehub-service-graphql/database"
)

func main() {
	database.Init()

	Serve()
}
