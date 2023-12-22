package database_relations

import (
	"time"

	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	"github.com/google/uuid"
)

type UserInvite struct {
	ID                       uuid.UUID                      `gorm:"primaryKey;type:uuid;not null"`
	Email                    string                         `gorm:"type:varchar(255);not null"`
	OrganizationID           uuid.UUID                      `gorm:"type:uuid;not null;field:organization_id"`
	InvitedByID              uuid.UUID                      `gorm:"type:uuid;not null";field:invited_by_id`
	OrganizationPermissions  []OrganizationsWithPermissions `gorm:"serializer:json;not null"`
	PlatformsWithPermissions []PlatformsWithPermissions     `gorm:"serializer:json;not null"`
	Code                     string                         `gorm:"type:varchar(60);not null"`
	Status                   database_enums.InvitationStatus
	CreatedAt                time.Time `gorm:"autoCreateTime"`
	UpdatedAt                *time.Time
}

func (ui UserInvite) ToOrganizationsPermissions(userId uuid.UUID) []OrganizationUser {
	var organizationUsers []OrganizationUser
	for _, organization := range ui.OrganizationPermissions {
		organizationUsers = append(organizationUsers, organization.ToOrganizationUsers(userId))
	}
	return organizationUsers
}

func (ui UserInvite) ToPlatformUsers(userId uuid.UUID) []PlatformUser {
	var platformUsers []PlatformUser
	for _, platform := range ui.PlatformsWithPermissions {
		platformUsers = append(platformUsers, platform.ToPlatformUsers(userId))
	}
	return platformUsers
}

type PlatformsWithPermissions struct {
	PlatformID          uuid.UUID
	PlatformPermissions []database_enums.PlatformPermission
}

type OrganizationsWithPermissions struct {
	OrganizationID          uuid.UUID
	OrganizationPermissions []database_enums.OrganizationPermission
}

func (owp OrganizationsWithPermissions) ToOrganizationUsers(userId uuid.UUID) OrganizationUser {
	return OrganizationUser{
		ID:             uuid.New(),
		UserID:         userId,
		OrganizationID: owp.OrganizationID,
		Permissions:    owp.OrganizationPermissions,
	}
}

func (pwp PlatformsWithPermissions) ToPlatformUsers(userId uuid.UUID) PlatformUser {
	return PlatformUser{
		ID:          uuid.New(),
		UserID:      userId,
		PlatformID:  pwp.PlatformID,
		Permissions: pwp.PlatformPermissions,
	}
}
