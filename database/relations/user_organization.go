package database_relations

import (
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	"github.com/google/uuid"
)

type UserOrganization struct {
	UserID         uuid.UUID                               `json:"user_id" gorm:"field:user_id;primaryKey;type:uuid;not null"`
	OrganizationID uuid.UUID                               `json:"organization_id"  gorm:"field:organization_id;primaryKey;type:uuid;not null"`
	Permissions    []database_enums.OrganizationPermission `gorm:"serializer:json;field:organization_permissions;not null;"`
	CreatedAt      string                                  `json:"createdAt"`
	UpdatedAt      *string                                 `json:"updatedAt,omitempty"`
	DeletedAt      *string                                 `json:"deletedAt,omitempty"`
}
