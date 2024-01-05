package services_domain

import (
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"

	"github.com/RouteHub-Link/routehub-service-graphql/graph/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DomainService struct {
	DB *gorm.DB
}

func (ds DomainService) CreateDomain(input model.DomainCreateInput) (domain database_models.Domain, err error) {
	// TODO Check that domain in user relations organization -> Permissions permissions can be checked by directive

	domain = database_models.Domain{
		ID:             uuid.New(),
		Name:           input.Name,
		OrganizationId: input.OrganizationID,
		URL:            input.URL,
		State:          database_enums.StatusStatePasive,
	}

	err = ds.DB.Create(&domain).Error
	return
}

func (ds DomainService) GetDomain(id uuid.UUID) (domain *database_models.Domain, err error) {
	err = ds.DB.Where("id = ?", id).First(&domain).Error
	return
}

func (ds DomainService) GetDomainByPlatformId(platformId uuid.UUID) (domain *database_models.Domain, err error) {
	joinQuery := ds.DB.Model(&database_models.Platform{}).Select("domain_id").Where("id = ?", platformId)

	err = ds.DB.Where("id = (?)", joinQuery).First(&domain).Error
	return
}

func (ds DomainService) GetDomainsByOrganization(organizationId uuid.UUID) (domains []*database_models.Domain, err error) {
	err = ds.DB.Where("organization_id = ?", organizationId).Find(&domains).Error
	return
}
