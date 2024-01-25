package services_user

import (
	"errors"

	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	database_relations "github.com/RouteHub-Link/routehub-service-graphql/database/relations"
	database_types "github.com/RouteHub-Link/routehub-service-graphql/database/types"
	"github.com/RouteHub-Link/routehub-service-graphql/graph/model"
	graph_inputs "github.com/RouteHub-Link/routehub-service-graphql/graph/model/inputs"
	services_utils "github.com/RouteHub-Link/routehub-service-graphql/services/utils"
	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func (u UserService) InviteUser(input graph_inputs.UserInviteInput, invitedById uuid.UUID) (userInvite *database_relations.UserInvite, err error) {
	id, err := gonanoid.New(24)
	if err != nil {
		id = uuid.New().String()
	}

	userInvite = &database_relations.UserInvite{
		ID:                       uuid.New(),
		Email:                    input.Email,
		InvitedByID:              invitedById,
		OrganizationPermissions:  input.OrganizationsPermissions,
		PlatformsWithPermissions: input.PlatformsWithPermissions,
		Status:                   database_enums.InvitationStatusPending,
		Code:                     id,
	}

	err = u.DB.Create(&userInvite).Error
	return
}

func (u UserService) UpdateInvitation(updateUserInviteInput model.UpdateUserInviteInput) (userInvite *database_relations.UserInvite, err error) {
	// TODO check email already invited or not
	err = u.DB.Where("code = ?", updateUserInviteInput.Code).First(&userInvite).Error
	if err != nil {
		return
	}

	if userInvite.Status != database_enums.InvitationStatusPending {
		return nil, errors.New("invitation is already used")
	}

	hashedPassword, err := services_utils.HashPassword(updateUserInviteInput.User.Password)
	if err != nil {
		return
	}

	userInvite.Status = updateUserInviteInput.Status
	err = u.DB.Save(&userInvite).Error
	if err != nil {
		return
	}

	user := &database_models.User{
		ID:       uuid.New(),
		Email:    userInvite.Email,
		Fullname: updateUserInviteInput.User.Fullname,
		Phone: &database_types.AccountPhone{
			CountryCode: updateUserInviteInput.User.Phone.CountryCode,
			Number:      updateUserInviteInput.User.Phone.Number,
		},
		PasswordHash: hashedPassword,
		Verified:     true,
	}

	err = u.DB.Create(&user).Error
	if err != nil {
		return
	}

	organizationUsers := userInvite.ToOrganizationsPermissions(user.ID)
	err = u.DB.Create(&organizationUsers).Error
	if err != nil {
		return
	}

	platformUsers := userInvite.ToPlatformUsers(user.ID)
	err = u.DB.Create(&platformUsers).Error

	return
}

func (u UserService) GetInvitesByInvitedById(invitedById uuid.UUID) (userInvites []*database_relations.UserInvite, err error) {
	err = u.DB.Where("invited_by_id = ?", invitedById).Find(&userInvites).Error
	return
}

func (u UserService) GetInvitedUserByInvitation(invitation database_relations.UserInvite) (userId *uuid.UUID, err error) {
	if invitation.Status != database_enums.InvitationStatusAccepted {
		return nil, errors.New("invitation is not accepted")
	}

	user := &database_models.User{}
	err = u.DB.Where("email = ?", invitation.Email).First(&user).Error

	return &user.ID, err
}
