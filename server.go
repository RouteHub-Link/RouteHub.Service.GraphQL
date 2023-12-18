package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/RouteHub-Link/routehub-service-graphql/auth"
	"github.com/RouteHub-Link/routehub-service-graphql/database"
	"github.com/RouteHub-Link/routehub-service-graphql/graph"
	"github.com/RouteHub-Link/routehub-service-graphql/services"
)

const defaultPort = "8080"

func Serve() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	resolver := &graph.Resolver{
		DB:          database.DB,
		UserService: &services.UserService{DB: database.DB},
	}

	config := graph.Config{Resolvers: resolver}
	config.Directives.Auth = auth.AuthDirectiveHandler

	var srv http.Handler = handler.NewDefaultServer(graph.NewExecutableSchema(config))

	srv = auth.Middleware(srv)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
