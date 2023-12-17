package services

import (
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	database_relations "github.com/RouteHub-Link/routehub-service-graphql/database/relations"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (u UserService) Users() (users []*database_models.User, err error) {
	err = u.DB.Find(&users).Error
	return
}

func (u UserService) UserCompany(userId uuid.UUID) (company []*database_models.Company, err error) {
	userCompanies := []database_relations.UserCompany{}
	err = u.DB.Where("user_id = ?", userId).Find(&userCompanies).Error
	if err != nil {
		return
	}

	companyIDs := []uuid.UUID{}
	for _, userCompany := range userCompanies {
		companyIDs = append(companyIDs, userCompany.CompanyID)
	}

	err = u.DB.Where("id IN ?", companyIDs).Find(&company).Error
	return
}
