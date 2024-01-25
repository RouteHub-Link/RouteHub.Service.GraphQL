package services_organization

import (
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_relations "github.com/RouteHub-Link/routehub-service-graphql/database/relations"
	"github.com/google/uuid"
	"github.com/vektah/gqlparser/gqlerror"
	"gorm.io/gorm"
)

type OrganizationPermissionService struct {
	DB *gorm.DB
}

func (o OrganizationPermissionService) GetOrganizationPermissions(userId uuid.UUID, organizationId uuid.UUID) (permissions []database_enums.OrganizationPermission, err error) {
	var organizationUser database_relations.OrganizationUser
	err = o.DB.Where("user_id = ? AND organization_id = ?", userId, organizationId).First(&organizationUser).Error
	if err != nil {
		return nil, gqlerror.Errorf("Access Denied")
	}
	return organizationUser.Permissions, nil

	/*
	   log.Printf("\n \norganizationUser: %+v\narg permission : %+v", organizationUser, permission)

	   	for _, organization_user_permission := range organizationUser.Permissions {
	   		if organization_user_permission == permission {
	   			return next(ctx)
	   		}
	   	}
	*/
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
