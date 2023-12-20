package graph_inputs

import (
	database_relations "github.com/RouteHub-Link/routehub-service-graphql/database/relations"
)

type UserInviteInput struct {
	Email                    string                                            `json:"email"`
	OrganizationsPermissions []database_relations.OrganizationsWithPermissions `json:"organizationsPermissions"`
	PlatformsWithPermissions []database_relations.PlatformsWithPermissions     `json:"platformsWithPermissions"`
}
