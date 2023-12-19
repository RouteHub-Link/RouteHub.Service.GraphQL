package services

import (
	"errors"

	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_relations "github.com/RouteHub-Link/routehub-service-graphql/database/relations"
	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func (u UserService) InviteUser(email string, invitedById uuid.UUID, organizationPermissions []database_relations.OrganizationsWithPermissions, platformsWithPermissions []database_relations.PlatformsWithPermissions) (userInvite *database_relations.UserInvite, err error) {
	id, err := gonanoid.New(24)
	if err != nil {
		id = uuid.New().String()
	}

	// TODO: Check the organizationId, invitedById, and platformsWithAccess in platform id's to make sure they are valid

	userInvite = &database_relations.UserInvite{
		Email:                    email,
		InvitedByID:              invitedById,
		OrganizationPermissions:  organizationPermissions,
		PlatformsWithPermissions: platformsWithPermissions,
		Code:                     id,
	}

	err = u.DB.Create(&userInvite).Error
	return
}

func (u UserService) UpdateInvitation(code string, status database_enums.InvitationStatus) (userInvite *database_relations.UserInvite, err error) {
	err = u.DB.Where("code = ?", code).First(&userInvite).Error
	if err != nil {
		return
	}

	if userInvite.Status != database_enums.InvitationStatusPending {
		return nil, errors.New("invitation is already used")
	}

	userInvite.Status = status
	err = u.DB.Save(&userInvite).Error

	// TODO 1. Create User

	// TODO 2. make organization users and platform users relations

	return
}
