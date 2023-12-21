package services_user

import (
	"errors"

	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	database_relations "github.com/RouteHub-Link/routehub-service-graphql/database/relations"
	"github.com/RouteHub-Link/routehub-service-graphql/graph/model"
	services_utils "github.com/RouteHub-Link/routehub-service-graphql/services/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (u UserService) User(id uuid.UUID) (user *database_models.User, err error) {
	err = u.DB.Where("id = ?", id).First(&user).Error
	return
}

func (u UserService) Users() (users []*database_models.User, err error) {
	err = u.DB.Find(&users).Error
	return
}

func (u UserService) UserOrganization(userId uuid.UUID) (Organization []*database_models.Organization, err error) {
	userOrganizations := []database_relations.UserOrganization{}
	err = u.DB.Where("user_id = ?", userId).Find(&userOrganizations).Error
	if err != nil {
		return
	}

	OrganizationIDs := []uuid.UUID{}
	for _, userOrganization := range userOrganizations {
		OrganizationIDs = append(OrganizationIDs, userOrganization.OrganizationID)
	}

	err = u.DB.Where("id IN ?", OrganizationIDs).Find(&Organization).Error
	return
}

func (u UserService) Login(input model.LoginInput) (user *database_models.User, err error) {
	err = u.DB.Where("email = ?", input.Email).First(&user).Error
	if err != nil {
		return
	}

	if !services_utils.CheckPasswordHash(input.Password, user.PasswordHash) {
		return nil, errors.New("invalid password")
	}

	return
}
