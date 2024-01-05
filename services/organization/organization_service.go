package services_organization

import (
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrganizationService struct {
	DB *gorm.DB
}

func (o OrganizationService) GetOrganization(id uuid.UUID) (organization *database_models.Organization, err error) {
	err = o.DB.Where("id = ?", id).First(&organization).Error
	return
}

func (o OrganizationService) GetOrganizations() (organizations []*database_models.Organization, err error) {
	err = o.DB.Find(&organizations).Error
	return
}

func (o OrganizationService) GetOrganizationsByIds(ids []uuid.UUID) (organizations []*database_models.Organization, err error) {
	err = o.DB.Where("id IN ?", ids).Find(&organizations).Error
	return
}
