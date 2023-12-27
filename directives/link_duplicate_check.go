package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func LinkDuplicateCheckDirectiveHandler(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {

	return next(ctx)
}
