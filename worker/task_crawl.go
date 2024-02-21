package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/RouteHub-Link/routehub-service-graphql/database"
	services_link "github.com/RouteHub-Link/routehub-service-graphql/services/link"
	"github.com/hibiken/asynq"
)

func NewCrawlURLTask(paload CrawlURLPayload) (*asynq.Task, error) {
	payload, err := json.Marshal(paload)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TaskNameCrawlURL, payload), nil
}

func HandleCrawlURL(ctx context.Context, t *asynq.Task) (err error) {
	var p CrawlURLPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	log.Printf("Crawling URL: %s", p.LinkUrl)

	db := database.DB
	linkService := &services_link.LinkService{DB: db}
	link, err := linkService.GetLinkById(p.LinkId)
	if err != nil {
		return
	}
	_crawl, err := linkService.GetCrawlById(p.CrawlId)
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
