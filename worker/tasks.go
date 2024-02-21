package worker

import "github.com/google/uuid"

// A list of task types.
const (
	TaskNameSendVerificationEmail = "email:verification:send"
	TaskNameCrawlURL              = "crawl:url"
)

type CrawlURLPayload struct {
	LinkId  uuid.UUID
	CrawlId uuid.UUID
	LinkUrl string
}
