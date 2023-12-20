package database_relations

import (
	"time"

	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	"github.com/google/uuid"
)

type PlatformUser struct {
	ID          uuid.UUID                           `gorm:"primaryKey;type:uuid;not null"`
	UserID      uuid.UUID                           `json:"userId"`
	PlatformID  uuid.UUID                           `json:"platformId"`
	Permissions []database_enums.PlatformPermission `gorm:"serializer:json;field:platform_permissions;not null;"`
	CreatedAt   time.Time                           `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   *time.Time                          `json:"updatedAt,omitempty"`
	DeletedAt   *time.Time                          `json:"deletedAt,omitempty"`
}
