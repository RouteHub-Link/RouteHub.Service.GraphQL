package services_link

import (
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	"github.com/RouteHub-Link/routehub-service-graphql/graph/model"
	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type LinkService struct {
	DB *gorm.DB
}

func (ls LinkService) CreateLink(input model.LinkCreateInput, userId uuid.UUID) (link *database_models.Link, err error) {
	path := input.Path
	if path == nil {
		genPath, err := gonanoid.New(6)
		if err != nil {
			return link, err
		}
		path = &genPath
	}

	link = &database_models.Link{
		ID:                uuid.New(),
		PlatformID:        input.PlatformID,
		Target:            input.Target,
		Path:              *path,
		RedirectionChoice: *input.RedirectionOptions,
		Status:            database_enums.StatusStateActive,
		CreatedBy:         userId,
		OpenGraph:         input.OpenGraph,
	}

	err = ls.DB.Create(&link).Error

	/* TODO Get OpenGraph data from target URL
	Mergings

	Platform OpenGraph
	User Input OpenGraph
	Crawler OpenGraph

	And Update the OG data in the database
	*/
	return
}

func (ls LinkService) GetLinkById(id uuid.UUID) (link database_models.Link, err error) {
	err = ls.DB.First(&link, id).Error
	return
}

func (ls LinkService) GetLinksByPlatformId(platformId uuid.UUID) (link []*database_models.Link, err error) {
	err = ls.DB.Where("platform_id = ?", platformId).Find(&link).Error
	return
}
