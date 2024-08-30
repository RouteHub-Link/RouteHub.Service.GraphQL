package database_models

import (
	"time"

	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_types "github.com/RouteHub-Link/routehub-service-graphql/database/types"
	"github.com/google/uuid"
)

type LinkCrawl struct {
	ID          uuid.UUID                       `json:"id"`
	LinkId      uuid.UUID                       `json:"link" gorm:"type:uuid;not null;field:link_id"`
	TaskId      uuid.UUID                       `json:"task,omitempty" gorm:"type:uuid;"`
	Target      string                          `json:"target,omitempty" gorm:"not null;"`
	CrawlStatus database_enums.CrawlStatus      `json:"crawlStatus,omitempty" gorm:"serializer:json;not null;field:crawl_status"`
	Logs        []*database_types.Log           `json:"logs,omitempty" gorm:"serializer:json;not null;"`
	Result      *database_types.MetaDescription `json:"result,omitempty" gorm:"serializer:json;field:open_graph"`
	CreatedBy   uuid.UUID                       `gorm:"type:uuid;not null;"`
	CreatedAt   time.Time                       `gorm:"autoCreateTime"`
	UpdatedAt   *time.Time                      `gorm:"autoUpdateTime:milli"`
	StartAt     *time.Time                      `gorm:"field:updated_at;"`
	EndAt       *time.Time                      `gorm:"field:end_at;"`
}

func (lc *LinkCrawl) Requested(link *Link, userId uuid.UUID, log *database_types.Log) {
	lc.ID = uuid.New()
	requestedTime := time.Now()
	lc.LinkId = link.ID
	lc.Target = link.Target
	lc.CrawlStatus = database_enums.CrawlStatusRequested
	lc.Logs = *new([]*database_types.Log)
	if log == nil {
		log = &database_types.Log{
			ID:        uuid.New(),
			CreatedAt: requestedTime,
			Message:   "Crawl requested",
		}
	} else {
		lc.Logs = append(lc.Logs, log)
	}

	lc.CreatedAt = requestedTime
	lc.CreatedBy = userId
}

func (lc *LinkCrawl) Started(log *database_types.Log) {
	startedTime := time.Now()
	lc.CrawlStatus = database_enums.CrawlStatusStarted
	if log == nil {
		log = &database_types.Log{
			ID:        uuid.New(),
			CreatedAt: startedTime,
			Message:   "Crawl started",
		}
	}
	lc.Logs = append(lc.Logs, log)
	lc.StartAt = &startedTime
}

func (lc *LinkCrawl) Finished(result *database_types.MetaDescription, log *database_types.Log, isSuccess bool) {
	finishedTime := time.Now()
	if isSuccess {
		lc.CrawlStatus = database_enums.CrawlStatusSuccess
	} else {
		lc.CrawlStatus = database_enums.CrawlStatusFailed
	}
	lc.Result = result
	if log == nil {
		log = &database_types.Log{
			ID:        uuid.New(),
			CreatedAt: finishedTime,
			Message:   "Crawl finished",
		}
	}
	lc.Logs = append(lc.Logs, log)
	lc.EndAt = &finishedTime
}
