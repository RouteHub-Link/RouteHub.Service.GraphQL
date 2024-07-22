package policies

import (
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	"github.com/casbin/casbin/v2"
	"github.com/google/uuid"
)

func (pb *PolicyBuilder) OrganizationRead(organizationId uuid.UUID) *PolicyBuilder {
	pb.AddPolicy(pb.sub.String(), organizationId.String(), database_enums.OrganizationPermissionRead.String())
	return pb
}

func (pb *PolicyBuilder) OrganizationUpdate(organizationId uuid.UUID) *PolicyBuilder {
	pb.AddPolicy(pb.sub.String(), organizationId.String(), database_enums.OrganizationPermissionUpdate.String())
	return pb
}

func (pb *PolicyBuilder) OrganizationDelete(organizationId uuid.UUID) *PolicyBuilder {
	pb.AddPolicy(pb.sub.String(), organizationId.String(), database_enums.OrganizationPermissionDelete.String())
	return pb
}

func (pb *PolicyBuilder) OrganizationPlatformCreate(organizationId uuid.UUID) *PolicyBuilder {
	pb.AddPolicy(pb.sub.String(), organizationId.String(), database_enums.OrganizationPermissionPlatformCreate.String())
	return pb
}

func (pb *PolicyBuilder) OrganizationUserInvite(organizationId uuid.UUID) *PolicyBuilder {
	pb.AddPolicy(pb.sub.String(), organizationId.String(), database_enums.OrganizationPermissionUserInvite.String())
	return pb
}

func EnforceOrganizationPermissions(e *casbin.SyncedCachedEnforcer, userId uuid.UUID, organizationId uuid.UUID, permissions []database_enums.OrganizationPermission) ([]PermissionActExplained, error) {
	permissionsAsStrings := []string{}
	for _, permission := range permissions {
		permissionsAsStrings = append(permissionsAsStrings, permission.String())
	}

	return EnforcePermissions(e, userId, organizationId, permissionsAsStrings)
}
