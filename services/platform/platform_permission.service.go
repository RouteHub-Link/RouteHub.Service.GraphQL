package services_platform

import (
	"log"

	"github.com/RouteHub-Link/routehub-service-graphql/auth"
	"github.com/RouteHub-Link/routehub-service-graphql/auth/policies"
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	"github.com/google/uuid"
	"github.com/vektah/gqlparser/gqlerror"
	"gorm.io/gorm"
)

type PlatformPermissionService struct {
	DB *gorm.DB
}

func (p PlatformPermissionService) GetPlatformPermissions(userId uuid.UUID, platformId uuid.UUID) (permissions []database_enums.PlatformPermission, err error) {
	res, err := policies.EnforcePlatformPermissions(auth.CasbinEnforcer, userId, platformId, database_enums.AllPlatformPermission)
	for _, permission := range res {
		permissions = append(permissions, database_enums.PlatformPermission(permission.Permission))
	}

	return permissions, err
}

func (p PlatformPermissionService) GetUserHasPermission(userId uuid.UUID, platformId uuid.UUID, permission database_enums.PlatformPermission) (hasPermission bool, err error) {
	e := auth.CasbinEnforcer
	hasPermission, exp, err := e.EnforceEx(userId.String(), platformId, permission.String())
	log.Printf("\nOrganizationPermissionDirectiveHandler EnforceEX;\nres: %+v\nexp: %+v\nerr: %+v\n\n", hasPermission, exp, err)

	if err != nil {
		return false, gqlerror.Errorf("Access Denied")
	}

	return
}
