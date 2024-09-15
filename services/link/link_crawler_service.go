package services_link

import (
	"errors"
	"log"

	"github.com/RouteHub-Link/routehub-service-graphql/database"
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	database_types "github.com/RouteHub-Link/routehub-service-graphql/database/types"
)

type LinkCrawlerService struct {
	link          *database_models.Link
	crawl         *database_models.LinkCrawl
	scrapedResult *database_types.MetaDescription
}

func NewLinkCrawlerService(link *database_models.Link, crawl *database_models.LinkCrawl) *LinkCrawlerService {
	return &LinkCrawlerService{link: link, crawl: crawl}
}

func (lcs *LinkCrawlerService) Crawl(updateFromDatabase bool) (err error) {
	//collyClient := clients.CollyClient{}

	lcs.Started()

	//lcs.link.OpenGraph = collyClient.VisitScrapeOG(lcs.link.Target)

	if updateFromDatabase {
		err = lcs.updateLink(lcs.link)
		if err != nil {
			return
		}
	}

	err = lcs.Finished(nil)

	return
}

func (lcs *LinkCrawlerService) updateLink(link *database_models.Link) (err error) {
	//ogJson, err := link.OpenGraph.AsJson()
	if err != nil {
		return
	}

	//err = database.DB.Model(&link).Update("open_graph", ogJson).Error
	if err != nil {
		return
	}

	link.State = database_enums.StatusStateActive
	err = database.DB.Save(&link).Error
	return
}

func (lcs *LinkCrawlerService) Started() (err error) {
	log.Printf("lcs.crawl %+v\n", lcs.crawl)
	lcs.crawl.Started(nil)
	err = database.DB.Save(lcs.crawl).Error
	return
}

func (lcs *LinkCrawlerService) Finished(isSuccess *bool) (err error) {
	//lcs.scrapedResult = lcs.link.OpenGraph
	success := false
	if isSuccess == nil {
		if lcs.scrapedResult == nil {
			lcs.scrapedResult = &database_types.MetaDescription{}
			success = false
		} else {
			success = true
		}
	} else {
		success = *isSuccess
	}

	lcs.crawl.Finished(lcs.scrapedResult, nil, success)
	err = database.DB.Save(lcs.crawl).Error
	return
}

func (lcs *LinkCrawlerService) CheckAlreadyCrawling() (err error) {
	var count int64
	err = database.DB.Model(&database_models.LinkCrawl{}).
		Where("link_id = ? AND end_at IS NULL", lcs.link.ID).
		Count(&count).Error
	isAlreadyCrawling := count > 0
	if isAlreadyCrawling {
		err = errors.New("link is already crawling")
	}
	return
}
