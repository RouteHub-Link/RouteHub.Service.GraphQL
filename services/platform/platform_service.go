package services_platform

import (
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	database_relations "github.com/RouteHub-Link/routehub-service-graphql/database/relations"
	graph_inputs "github.com/RouteHub-Link/routehub-service-graphql/graph/model/inputs"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlatformService struct {
	DB *gorm.DB
}

func (ps PlatformService) CreatePlatform(input graph_inputs.PlatformCreateInput, userId uuid.UUID) (platform database_models.Platform, err error) {
	// TODO Check that domain in user relations organization -> Permissions permissions can be checked by directive

	platform = database_models.Platform{
		ID:                uuid.New(),
		Name:              input.Name,
		DomainId:          input.DomainID,
		RedirectionChoice: input.RedirectionChoice,
		OpenGraph:         input.OpenGraph,
		CreatedBy:         userId,
		Status:            database_enums.StatusStatePasive,
	}

	err = ps.DB.Create(&platform).Error
	if err != nil {
		return
	}

	organization_relation := database_relations.OrganizationPlatform{
		ID:             uuid.New(),
		OrganizationID: input.OrganizationID,
		PlatformID:     platform.ID,
	}

	err = ps.DB.Create(&organization_relation).Error
	if err != nil {
		return
	}

	platform_user := database_relations.PlatformUser{
		ID:          uuid.New(),
		UserID:      userId,
		PlatformID:  platform.ID,
		Permissions: database_enums.AllPlatformPermission,
	}

	err = ps.DB.Create(&platform_user).Error

	return
}

func (ps PlatformService) GetPlatform(id uuid.UUID) (platform *database_models.Platform, err error) {
	err = ps.DB.Where("id = ?", id).First(&platform).Error
	return
}

func (ps PlatformService) GetPlatformsByOrganization(organizationId uuid.UUID) (platforms []*database_models.Platform, err error) {
	joinQuery := ps.DB.Model(&database_relations.OrganizationPlatform{}).Select("platform_id").Where("organization_id = ?", organizationId)
	err = ps.DB.Where("id IN (?)", joinQuery).Find(&platforms).Error
	return
}

func (ps PlatformService) GetPlatformsByUser(userId uuid.UUID) (platforms []*database_models.Platform, err error) {
	joinQuery := ps.DB.Model(&database_relations.PlatformUser{}).Select("platform_id").Where("user_id = ?", userId)
	err = ps.DB.Where("id IN (?)", joinQuery).Find(&platforms).Error
	return
}

func (ps PlatformService) GetPlatforms() (platforms []*database_models.Platform, err error) {
	err = ps.DB.Find(&platforms).Error
	return
}
