package worker

import (
	"log"

	"github.com/hibiken/asynq"
)

type WrokerService struct{}

func (WrokerService) client() *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: config.Redis.Addres})
}

func (ws WrokerService) NewCrawlURL(paload CrawlURLPayload) (*asynq.TaskInfo, error) {
	client := ws.client()
	defer client.Close()

	task, err := NewCrawlURLTask(paload)
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}

	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}

	return info, err
}
