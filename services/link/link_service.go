package services_link

import (
	"github.com/RouteHub-Link/routehub-service-graphql/database"
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	database_types "github.com/RouteHub-Link/routehub-service-graphql/database/types"
	"github.com/RouteHub-Link/routehub-service-graphql/graph/model"
	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type LinkService struct {
	DB *gorm.DB
}

func (ls LinkService) CreateLink(input model.LinkCreateInput, userId uuid.UUID, validationTaskId uuid.UUID) (link *database_models.Link, err error) {
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
		ValidationTaskID:  validationTaskId,
		Target:            input.Target,
		Path:              *path,
		RedirectionChoice: *input.RedirectionOptions,
		State:             database_enums.StatusStatePasive,
		CreatedBy:         userId,
		OpenGraph:         input.OpenGraph,
	}

	err = ls.DB.Create(&link).Error
	if err != nil {
		return
	}

	return
}

func (ls LinkService) SaveCrawlRequest(link *database_models.Link, userId uuid.UUID) (crawlId uuid.UUID, err error) {
	crawl := &database_models.LinkCrawl{}
	crawl.Requested(link, userId, nil)
	err = database.DB.Create(crawl).Error

	crawlId = crawl.ID
	return
}

func (ls LinkService) GetLinkById(id uuid.UUID) (link *database_models.Link, err error) {
	err = ls.DB.First(&link, id).Error
	return
}

func (ls LinkService) GetLinksByPlatformId(platformId uuid.UUID) (link []*database_models.Link, err error) {
	err = ls.DB.Where("platform_id = ?", platformId).Find(&link).Error
	return
}

func (ls LinkService) UpdateLinkStatus(link *database_models.Link, status database_enums.StatusState) (err error) {
	link.State = status
	err = ls.DB.Save(&link).Error
	return
}

func (ls LinkService) GetCrawls(linkId uuid.UUID) (crawls []*database_models.LinkCrawl, err error) {
	err = ls.DB.Where("link_id = ?", linkId).Find(&crawls).Error
	return
}

func (ls LinkService) GetCrawlById(id uuid.UUID) (crawl *database_models.LinkCrawl, err error) {
	err = ls.DB.First(&crawl, id).Error
	return
}

func (ls LinkService) AddToPinnedLinks(platformId uuid.UUID, linkId, userId uuid.UUID) (err error) {
	platform := &database_models.Platform{}
	err = ls.DB.First(&platform, platformId).Error
	if err != nil {
		return
	}

	pinnedLink := &database_types.PinnedLink{
		LinkID:    linkId,
		PinnedBy:  userId,
		CreatedAt: "now()",
	}

	if platform.PinnedLinks == nil {
		platform.PinnedLinks = &[]database_types.PinnedLink{}
	}

	newPinnedLinks := append(*platform.PinnedLinks, *pinnedLink)
	platform.PinnedLinks = &newPinnedLinks

	err = ls.DB.Save(&platform).Error
	return
}

func (ls LinkService) RemoveFromPinnedLinks(platformId uuid.UUID, linkId uuid.UUID) (err error) {
	platform := &database_models.Platform{}
	err = ls.DB.First(&platform, platformId).Error
	if err != nil {
		return
	}

	if platform.PinnedLinks == nil {
		return
	}

	newPinnedLinks := []database_types.PinnedLink{}
	for _, pinnedLink := range *platform.PinnedLinks {
		if pinnedLink.LinkID != linkId {
			newPinnedLinks = append(newPinnedLinks, pinnedLink)
		}
	}

	platform.PinnedLinks = &newPinnedLinks

	err = ls.DB.Save(&platform).Error
	return
}

func (ls LinkService) GetLinksByIds(ids []uuid.UUID) (links []*database_models.Link, err error) {
	err = ls.DB.Where("id IN ?", ids).Find(&links).Error
	return
}
