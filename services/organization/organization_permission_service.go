package services_organization

import (
	"github.com/RouteHub-Link/routehub-service-graphql/auth"
	"github.com/RouteHub-Link/routehub-service-graphql/auth/policies"
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	"github.com/google/uuid"
	"github.com/vektah/gqlparser/gqlerror"
	"gorm.io/gorm"
)

type OrganizationPermissionService struct {
	DB *gorm.DB
}

func (o OrganizationPermissionService) GetOrganizationPermissions(userId uuid.UUID, organizationId uuid.UUID) (permissions []database_enums.OrganizationPermission, err error) {

	res, err := policies.EnforceOrganizationPermissions(auth.CasbinEnforcer, userId, organizationId, database_enums.AllOrganizationPermission)
	if err != nil {
		return nil, gqlerror.Errorf("Access Denied")
	}

	for _, permission := range res {
		permissions = append(permissions, database_enums.OrganizationPermission(permission.Permission))
	}

	return permissions, nil
}

func (o OrganizationPermissionService) GetUserHasPermission(userId uuid.UUID, organizationId uuid.UUID, permission database_enums.OrganizationPermission) (hasPermission bool, err error) {
	permissions, err := o.GetOrganizationPermissions(userId, organizationId)
	if err != nil {
		return false, err
	}

	for _, organization_user_permission := range permissions {
		if organization_user_permission == permission {
			return true, nil
		}
	}
	return false, nil
}
