package services_platform

import (
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_relations "github.com/RouteHub-Link/routehub-service-graphql/database/relations"
	"github.com/google/uuid"
	"github.com/vektah/gqlparser/gqlerror"
	"gorm.io/gorm"
)

type PlatformPermissionService struct {
	DB *gorm.DB
}

func (p PlatformPermissionService) GetPlatformPermissions(userId uuid.UUID, platformId uuid.UUID) (platformPermissions []database_enums.PlatformPermission, err error) {
	var platformUser database_relations.PlatformUser
	err = p.DB.Where("user_id = ? AND platform_id = ?", userId, platformId).First(&platformUser).Error
	if err != nil {
		return nil, err
	}
	return platformUser.Permissions, err
}

func (p PlatformPermissionService) GetUserHasPermission(userId uuid.UUID, platformId uuid.UUID, permission database_enums.PlatformPermission) (hasPermission bool, err error) {
	platformPermissions, err := p.GetPlatformPermissions(userId, platformId)
	if err != nil {
		return false, gqlerror.Errorf("Access Denied")
	}

	for _, platformPermission := range platformPermissions {
		if platformPermission == permission {
			hasPermission = true
			return
		}
	}

	return
}
