package services_user

import (
	"errors"

	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	database_relations "github.com/RouteHub-Link/routehub-service-graphql/database/relations"
	"github.com/RouteHub-Link/routehub-service-graphql/graph/model"
	services_organization "github.com/RouteHub-Link/routehub-service-graphql/services/organization"
	services_utils "github.com/RouteHub-Link/routehub-service-graphql/services/utils"
	"github.com/google/uuid"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (u UserService) User(id uuid.UUID) (user *database_models.User, err error) {
	err = u.DB.Where("id = ?", id).First(&user).Error
	return
}

func (u UserService) CreateUser(input model.UserCreateInput) (user *database_models.User, err error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	password := string(hash)

	user = &database_models.User{
		ID:           uuid.New(),
		Email:        input.Email,
		PasswordHash: password,
		Fullname:     input.Fullname,
		Verified:     false,
	}
	err = u.DB.Create(user).Error
	if err != nil {
		return
	}

	if input.Organization != nil {
		organizationService := services_organization.OrganizationService{DB: u.DB}
		_, err := organizationService.CreateOrganization(user.ID, input.Organization)
		if err != nil {
			return nil, gqlerror.Errorf("User created but organization cannot be created %v", err)
		}
	}

	return
}

func (u UserService) Users() (users []*database_models.User, err error) {
	err = u.DB.Find(&users).Error
	return
}

func (u UserService) UsersByIds(ids []uuid.UUID) (users []*database_models.User, err error) {
	err = u.DB.Where("id IN ?", ids).Find(&users).Error
	return
}

func (u UserService) UsersByOrganization(organizationId uuid.UUID) (users []*database_models.User, err error) {
	organizationUsers := []database_relations.OrganizationUser{}
	err = u.DB.Where("organization_id = ?", organizationId).Find(&organizationUsers).Error
	if err != nil {
		return
	}

	joinQuery := u.DB.Model(&database_relations.OrganizationUser{}).Select("user_id").Where("organization_id = ?", organizationId)
	err = u.DB.Where("id IN ?", joinQuery).Find(&users).Error
	return
}

func (u UserService) UserIdsByOrganization(organizationId uuid.UUID) (userIds []uuid.UUID, err error) {
	organizationUsers := []database_relations.OrganizationUser{}
	err = u.DB.Where("organization_id = ?", organizationId).Find(&organizationUsers).Error
	if err != nil {
		return
	}

	for _, organizationUser := range organizationUsers {
		userIds = append(userIds, organizationUser.UserID)
	}

	return
}

func (u UserService) OrganizationUser(userId uuid.UUID) (Organization []*database_models.Organization, err error) {
	organizationUsers := []database_relations.OrganizationUser{}
	err = u.DB.Where("user_id = ?", userId).Find(&organizationUsers).Error
	if err != nil {
		return
	}

	OrganizationIDs := []uuid.UUID{}
	for _, organizationUser := range organizationUsers {
		OrganizationIDs = append(OrganizationIDs, organizationUser.OrganizationID)
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
