package worker

import (
	"log"

	configuration "github.com/RouteHub-Link/routehub-service-graphql/config"
	"github.com/hibiken/asynq"
)

const (
	QueueCritical = "critical"
	QueueDefault  = "default"
	QueueLow      = "low"
)

var (
	config *configuration.ApplicationConfig
)

func Init() {
	config = configuration.ConfigurationService{}.Get()

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: config.Redis.Addres},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 3,

			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				QueueCritical: 6,
				QueueDefault:  3,
				QueueLow:      1,
			},
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(TaskNameCrawlURL, HandleCrawlURL)
	//mux.Handle(tasks.TypeImageResize, tasks.NewImageProcessor())

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
