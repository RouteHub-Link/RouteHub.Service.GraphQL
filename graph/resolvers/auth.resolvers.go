package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"

	"github.com/RouteHub-Link/routehub-service-graphql/auth"
	"github.com/RouteHub-Link/routehub-service-graphql/graph"
	"github.com/RouteHub-Link/routehub-service-graphql/graph/model"
)

// LoginUser is the resolver for the loginUser field.
func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginInput) (*model.LoginPayload, error) {
	userService := r.ServiceContainer.UserService

	user, err := userService.Login(input)
	if err != nil {
		return nil, err
	}

	userSession := new(auth.UserSession)
	userSession.ID = user.ID
	userSession.Name = user.Fullname

	token, err := auth.GenerateToken(userSession.ToClaims())

	if err != nil {
		return nil, err
	}

	return &model.LoginPayload{
		Token: token,
	}, nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
