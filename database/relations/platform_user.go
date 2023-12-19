package database_relations

import database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"

type PlatformUser struct {
	UserID      string                              `json:"userId"`
	PlatformID  string                              `json:"platformId"`
	Permissions []database_enums.PlatformPermission `gorm:"serializer:json;field:platform_permissions;not null;"`
	CreatedAt   string                              `json:"createdAt"`
	UpdatedAt   *string                             `json:"updatedAt,omitempty"`
	DeletedAt   *string                             `json:"deletedAt,omitempty"`
}
