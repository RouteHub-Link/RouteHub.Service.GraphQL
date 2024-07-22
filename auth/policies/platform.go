package policies

import (
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	"github.com/casbin/casbin/v2"
	"github.com/google/uuid"
)

func (pb *PolicyBuilder) PlatformRead(platformId uuid.UUID) *PolicyBuilder {
	pb.AddPolicy(pb.sub.String(), platformId.String(), database_enums.PlatformPermissionRead.String())
	return pb
}

func (pb *PolicyBuilder) PlatformUpdate(platformId uuid.UUID) *PolicyBuilder {
	pb.AddPolicy(pb.sub.String(), platformId.String(), database_enums.PlatformPermissionUpdate.String())
	return pb
}

func (pb *PolicyBuilder) PlatformLinkCreate(platformId uuid.UUID) *PolicyBuilder {
	pb.AddPolicy(pb.sub.String(), platformId.String(), database_enums.PlatformPermissionLinkCreate.String())
	return pb
}

func (pb *PolicyBuilder) PlatformLinkUpdate(platformId uuid.UUID) *PolicyBuilder {
	pb.AddPolicy(pb.sub.String(), platformId.String(), database_enums.PlatformPermissionLinkUpdate.String())
	return pb
}

func (pb *PolicyBuilder) PlatformLinkDelete(platformId uuid.UUID) *PolicyBuilder {
	pb.AddPolicy(pb.sub.String(), platformId.String(), database_enums.PlatformPermissionLinkDelete.String())
	return pb
}

func (pb *PolicyBuilder) PlatformLinkRead(platformId uuid.UUID) *PolicyBuilder {
	pb.AddPolicy(pb.sub.String(), platformId.String(), database_enums.PlatformPermissionLinkRead.String())
	return pb
}

func EnforcePlatformPermissions(e *casbin.Enforcer, userId uuid.UUID, platformId uuid.UUID, permissions []database_enums.PlatformPermission) ([]PermissionActExplained, error) {
	permissionsAsStrings := []string{}
	for _, permission := range permissions {
		permissionsAsStrings = append(permissionsAsStrings, permission.String())
	}

	return EnforcePermissions(e, userId, platformId, permissionsAsStrings)
}
