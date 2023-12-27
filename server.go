package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/RouteHub-Link/routehub-service-graphql/auth"
	"github.com/RouteHub-Link/routehub-service-graphql/database"
	"github.com/RouteHub-Link/routehub-service-graphql/directives"
	"github.com/RouteHub-Link/routehub-service-graphql/graph"
	Resolvers "github.com/RouteHub-Link/routehub-service-graphql/graph/resolvers"
	"github.com/RouteHub-Link/routehub-service-graphql/loaders"
	"github.com/RouteHub-Link/routehub-service-graphql/services"
)

const defaultPort = "8080"

func Serve() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	resolver := &Resolvers.Resolver{
		DB:               database.DB,
		ServiceContainer: services.NewServiceContainer(database.DB),
		LoaderContainer:  loaders.NewLoaderContainer(),
	}

	config := graph.Config{Resolvers: resolver}
	config.Directives.Auth = directives.AuthDirectiveHandler
	config.Directives.OrganizationPermission = directives.OrganizationPermissionDirectiveHandler
	config.Directives.PlatformPermission = directives.PlatformPermissionDirectiveHandler
	config.Directives.DomainURLCheck = directives.DomainURLCheckDirectiveHandler

	var srv http.Handler = handler.NewDefaultServer(graph.NewExecutableSchema(config))

	srv = auth.Middleware(srv)
	srv = loaders.Middleware(srv)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
