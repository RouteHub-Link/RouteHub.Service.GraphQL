package worker

import (
	"log"

	"github.com/RouteHub-Link/routehub-service-graphql/database"
	services_link "github.com/RouteHub-Link/routehub-service-graphql/services/link"
)

func HandleCrawlURL(payload CrawlURLPayload) (err error) {

	log.Printf("Crawling URL: %s", payload.LinkUrl)

	db := database.DB
	linkService := &services_link.LinkService{DB: db}
	link, err := linkService.GetLinkById(payload.LinkId)
	if err != nil {
		return
	}
	_crawl, err := linkService.GetCrawlById(payload.CrawlId)
	if err != nil {
		return
	}

	linkCrawlerService := services_link.NewLinkCrawlerService(link, _crawl)

	err = linkCrawlerService.Crawl(true)
	if err != nil {
		return
	}

	return nil
}
