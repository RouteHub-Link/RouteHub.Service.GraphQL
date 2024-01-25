package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/RouteHub-Link/routehub-service-graphql/auth"
	"github.com/RouteHub-Link/routehub-service-graphql/database"
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	services_organization "github.com/RouteHub-Link/routehub-service-graphql/services/organization"
	"github.com/google/uuid"
	"github.com/vektah/gqlparser/gqlerror"
)

func OrganizationPermissionDirectiveHandler(ctx context.Context, obj interface{}, next graphql.Resolver, permission database_enums.OrganizationPermission) (res interface{}, err error) {
	userSession := auth.ForContext(ctx)
	db := database.DB
	if userSession == nil {
		return nil, gqlerror.Errorf("Access Denied")
	}

	organizationId, ok := obj.(map[string]interface{})["organizationId"].(string)
	if !ok {
		return nil, gqlerror.Errorf("organizationId not found in obj")
	}

	organizationUUID := uuid.MustParse(organizationId)
	organizationPermissionService := services_organization.OrganizationPermissionService{DB: db}
	hasPermission, err := organizationPermissionService.GetUserHasPermission(userSession.ID, organizationUUID, permission)
	if hasPermission {
		return next(ctx)
	}

	if err != nil {
		return nil, err
	}

	return nil, gqlerror.Errorf("Access Denied")

	/*
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
	*/

}
