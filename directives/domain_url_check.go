package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/RouteHub-Link/routehub-service-graphql/database"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	"github.com/vektah/gqlparser/gqlerror"
)

func DomainURLCheckDirectiveHandler(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	domain_url, ok := obj.(map[string]interface{})["url"].(string)
	if !ok {
		return nil, gqlerror.Errorf("url not found in obj")
	}

	db := database.DB
	var count int64

	// domain must be not found
	db.Model(&database_models.Domain{}).Where("url = ?", domain_url).Count(&count)

	if count != 0 {
		return nil, gqlerror.Errorf("domain url already exists")
	}

	return next(ctx)
}
