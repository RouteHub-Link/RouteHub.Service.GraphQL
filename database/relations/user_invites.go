package database_relations

import (
	"time"

	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	"github.com/google/uuid"
)

type UserInvite struct {
	ID                       uuid.UUID                      `gorm:"primaryKey;type:uuid;not null"`
	Email                    string                         `gorm:"type:varchar(255);not null"`
	OrganizationID           uuid.UUID                      `gorm:"type:uuid;not null"`
	InvitedByID              uuid.UUID                      `gorm:"type:uuid;not null"`
	OrganizationPermissions  []OrganizationsWithPermissions `gorm:"serializer:json;not null"`
	PlatformsWithPermissions []PlatformsWithPermissions     `gorm:"serializer:json;not null"`
	Code                     string                         `gorm:"type:varchar(60);not null"`
	Status                   database_enums.InvitationStatus
	CreatedAt                time.Time
	UpdatedAt                time.Time
}

type PlatformsWithPermissions struct {
	PlatformID          uuid.UUID
	PlatformPermissions []database_enums.PlatformPermission
}

type OrganizationsWithPermissions struct {
	OrganizationID          uuid.UUID
	OrganizationPermissions []database_enums.OrganizationPermission
}
