package database_relations

import (
	"time"

	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	"github.com/google/uuid"
)

type UserOrganization struct {
	ID             uuid.UUID                               `gorm:"primaryKey;type:uuid;not null"`
	UserID         uuid.UUID                               `json:"user_id" gorm:"field:user_id;primaryKey;type:uuid;not null"`
	OrganizationID uuid.UUID                               `json:"organization_id"  gorm:"field:organization_id;primaryKey;type:uuid;not null"`
	Permissions    []database_enums.OrganizationPermission `gorm:"serializer:json;field:organization_permissions;not null;"`
	CreatedAt      time.Time                               `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt      *time.Time                              `json:"updatedAt,omitempty"`
	DeletedAt      *time.Time                              `json:"deletedAt,omitempty"`
}
