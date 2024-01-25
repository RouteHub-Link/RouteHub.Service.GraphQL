package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/RouteHub-Link/routehub-service-graphql/auth"
	Configuration "github.com/RouteHub-Link/routehub-service-graphql/config"
	"github.com/RouteHub-Link/routehub-service-graphql/database"
	"github.com/RouteHub-Link/routehub-service-graphql/directives"
	"github.com/RouteHub-Link/routehub-service-graphql/graph"
	Resolvers "github.com/RouteHub-Link/routehub-service-graphql/graph/resolvers"
	"github.com/RouteHub-Link/routehub-service-graphql/loaders"
	"github.com/RouteHub-Link/routehub-service-graphql/services"
)

var applicationConfig = Configuration.ConfigurationService{}.Get()

func Serve() {
	resolver := &Resolvers.Resolver{
		DB:               database.DB,
		ServiceContainer: services.NewServiceContainer(database.DB),
		LoaderContainer:  loaders.NewLoaderContainer(),
	}

	config := graph.Config{Resolvers: resolver}
	directives.Assign(&config)

	var srv http.Handler = handler.NewDefaultServer(graph.NewExecutableSchema(config))

	srv = auth.Middleware(srv)
	srv = loaders.Middleware(srv)

	if applicationConfig.GraphQL.Playground {
		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
		log.Printf("GraphQL playground enabled Connect to http://localhost:%s/ for GraphQL playground", applicationConfig.GraphQL.PortAsString)
	}

	http.Handle("/query", srv)

	log.Printf("You can interact with the GraphQL API using http://localhost:%s/query", applicationConfig.GraphQL.PortAsString)
	log.Fatal(http.ListenAndServe(":"+applicationConfig.GraphQL.PortAsString, nil))
}
