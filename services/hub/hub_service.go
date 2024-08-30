package hub

import (
	"github.com/RouteHub-Link/routehub-service-graphql/clients/mq"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
)

type HubService struct {
	platform database_models.Platform
	MQC      *mq.MQTTClient
}

func NewHubService(platform database_models.Platform) (*HubService, error) {
	mqc, err := mq.NewMQTTClient(platform.TCPAddr)
	if err != nil {
		return nil, err
	}

	rc := HubService{
		platform: platform,
		MQC:      mqc,
	}
	return &rc, nil
}
