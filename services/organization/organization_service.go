package services_organization

import (
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	database_relations "github.com/RouteHub-Link/routehub-service-graphql/database/relations"
	"github.com/RouteHub-Link/routehub-service-graphql/graph/model"
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

func (o OrganizationService) CreateOrganization(userId uuid.UUID, input *model.OrganizationCreateInput) (organization *database_models.Organization, err error) {
	organization = &database_models.Organization{
		ID:          uuid.New(),
		Name:        input.Name,
		Website:     input.Website,
		Description: input.Description,
		Location:    input.Location,
	}

	err = o.DB.Create(&organization).Error
	if err != nil {
		return
	}

	userOrganization := &database_relations.OrganizationUser{
		ID:             uuid.New(),
		UserID:         userId,
		OrganizationID: organization.ID,
		Permissions:    database_enums.AllOrganizationPermission,
	}

	err = o.DB.Create(&userOrganization).Error
	return
}

func (o OrganizationService) UpdateOrganization(input model.OrganizationUpdateInput) (organization *database_models.Organization, err error) {
	organization = &database_models.Organization{
		ID:          input.OrganizationID,
		Name:        input.Name,
		Website:     input.Website,
		Description: input.Description,
		Location:    input.Location,
	}

	err = o.DB.Save(&organization).Error
	return
}
