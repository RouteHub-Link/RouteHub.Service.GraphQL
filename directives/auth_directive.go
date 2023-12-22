package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/RouteHub-Link/routehub-service-graphql/auth"
	"github.com/vektah/gqlparser/gqlerror"
)

func AuthDirectiveHandler(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	tokenData := auth.ForContext(ctx)
	if tokenData == nil {
		return nil, &gqlerror.Error{
			Message: "Access Denied",
		}
	}

	return next(ctx)
}
