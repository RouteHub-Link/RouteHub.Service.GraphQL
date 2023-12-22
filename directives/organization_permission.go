package directives

import (
	"context"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/RouteHub-Link/routehub-service-graphql/auth"
	"github.com/RouteHub-Link/routehub-service-graphql/database"
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_relations "github.com/RouteHub-Link/routehub-service-graphql/database/relations"
	"github.com/google/uuid"
	"github.com/vektah/gqlparser/gqlerror"
)

func OrganizationPermissionDirectiveHandler(ctx context.Context, obj interface{}, next graphql.Resolver, permission database_enums.OrganizationPermission) (res interface{}, err error) {
	userSession := auth.ForContext(ctx)
	db := database.DB
	if userSession == nil {
		return nil, gqlerror.Errorf("Access Denied")
	}

	//log.Printf("ctx: %+v\n obj %+v\n", ctx, obj)

	organizationId, ok := obj.(map[string]interface{})["organizationId"].(string)
	if !ok {
		return nil, gqlerror.Errorf("organizationId not found in obj")
	}

	organizationUUID := uuid.MustParse(organizationId)
	var organizationUser database_relations.OrganizationUser
	err = db.Where("user_id = ? AND organization_id = ?", userSession.ID, organizationUUID).First(&organizationUser).Error
	if err != nil {
		return nil, gqlerror.Errorf("Access Denied")
	}

	log.Printf("\n \norganizationUser: %+v\narg permission : %+v", organizationUser, permission)

	for _, organization_user_permission := range organizationUser.Permissions {
		if organization_user_permission == permission {
			return next(ctx)
		}
	}
	return nil, gqlerror.Errorf("Access Denied")
}
