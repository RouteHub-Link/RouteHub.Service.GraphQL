package worker

import (
	"log"
)

type WrokerService struct{}

func (ws WrokerService) NewCrawlURL(payload CrawlURLPayload) error {
	err := HandleCrawlURL(payload)
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}
	return err
}
