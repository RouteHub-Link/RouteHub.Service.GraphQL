package main

import (
	"log"
	"net/http"

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

func Serve() {
	resolver := &Resolvers.Resolver{
		DB:               database.DB,
		ServiceContainer: services.NewServiceContainer(database.DB),
		LoaderContainer:  loaders.NewLoaderContainer(),
	}

	config := graph.Config{Resolvers: resolver}
	directives.Assign(&config)

	var srv http.Handler = handler.NewDefaultServer(graph.NewExecutableSchema(config))

	//srv = auth.JWTMiddleware(srv)
	srv = auth.PKCEMiddleware(srv)
	srv = loaders.Middleware(srv)

	if applicationConfig.GraphQL.Playground {
		http.Handle("/playground", playground.ApolloSandboxHandler("GraphQL playground", "/query"))
		log.Printf("GraphQL playground enabled Connect to %s/ for GraphQL playground", applicationConfig.Host)
	}

	http.Handle("/query", srv)
	http.Handle("/healthz", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	log.Printf("You can interact with the GraphQL API using %s/query", applicationConfig.Host)
	log.Printf("For logging in, you can use the link %s/oauth2/login", applicationConfig.Host)

	log.Fatal(http.ListenAndServe(":"+applicationConfig.GraphQL.PortAsString, nil))
}
