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
	CreatedAt                time.Time `gorm:"autoCreateTime"`
	UpdatedAt                *time.Time
}

func (ui UserInvite) ToOrganizationsPermissions(userId uuid.UUID) []UserOrganization {
	var userOrganizations []UserOrganization
	for _, organization := range ui.OrganizationPermissions {
		userOrganizations = append(userOrganizations, organization.ToUserOrganizations(userId))
	}
	return userOrganizations
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

func (owp OrganizationsWithPermissions) ToUserOrganizations(userId uuid.UUID) UserOrganization {
	return UserOrganization{
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
