package services_link

import (
	"errors"
	"log"
	"time"

	"github.com/RouteHub-Link/routehub-service-graphql/clients"
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	database_types "github.com/RouteHub-Link/routehub-service-graphql/database/types"
	"gorm.io/gorm"
)

type LinkCrawlerService struct {
	DB            *gorm.DB
	link          *database_models.Link
	user          *database_models.User
	crawl         *database_models.LinkCrawl
	scrapedResult *database_types.OpenGraph
}

func (lcs *LinkCrawlerService) Crawl(updateFromDatabase bool) (err error) {
	collyClient := clients.CollyClient{}

	err = lcs.CheckAlreadyCrawling()
	if err != nil {
		return
	}

	lcs.Requested()

	go func() {
		// TODO Can be implemented queue system
		time.Sleep(10 * time.Second)

		lcs.Started()

		lcs.link.OpenGraph = collyClient.VisitScrapeOG(lcs.link.Target)

		time.Sleep(10 * time.Second)

		if updateFromDatabase {
			err = lcs.updateLink(lcs.link)
			if err != nil {
				return
			}
		}

		err = lcs.Finished(nil)
	}()

	return
}

func (lcs *LinkCrawlerService) updateLink(link *database_models.Link) (err error) {
	ogJson, err := link.OpenGraph.AsJson()
	if err != nil {
		return
	}

	err = lcs.DB.Model(&link).Update("open_graph", ogJson).Error
	if err != nil {
		return
	}

	link.State = database_enums.StatusStateActive
	err = lcs.DB.Save(&link).Error
	return
}

func (lcs *LinkCrawlerService) Requested() (err error) {
	lcs.crawl = &database_models.LinkCrawl{}
	lcs.crawl.Requested(lcs.link, lcs.user.ID, nil)
	log.Printf("LinkCrawlerService.Requested() %+v\n", lcs.crawl)
	err = lcs.DB.Create(lcs.crawl).Error
	log.Printf("lcs.crawl from requested %+v\n", lcs.crawl)
	return
}

func (lcs *LinkCrawlerService) Started() (err error) {
	log.Printf("lcs.crawl %+v\n", lcs.crawl)
	lcs.crawl.Started(nil)
	err = lcs.DB.Save(lcs.crawl).Error
	return
}

func (lcs *LinkCrawlerService) Finished(isSuccess *bool) (err error) {
	lcs.scrapedResult = lcs.link.OpenGraph
	success := false
	if isSuccess == nil {
		if lcs.scrapedResult == nil {
			lcs.scrapedResult = &database_types.OpenGraph{}
			success = false
		} else {
			success = true
		}
	} else {
		success = *isSuccess
	}

	lcs.crawl.Finished(lcs.scrapedResult, nil, success)
	err = lcs.DB.Save(lcs.crawl).Error
	return
}

func (lcs *LinkCrawlerService) CheckAlreadyCrawling() (err error) {
	var count int64
	err = lcs.DB.Model(&database_models.LinkCrawl{}).
		Where("link_id = ? AND end_at IS NULL", lcs.link.ID).
		Count(&count).Error
	isAlreadyCrawling := count > 0
	if isAlreadyCrawling {
		err = errors.New("link is already crawling")
	}
	return
}
