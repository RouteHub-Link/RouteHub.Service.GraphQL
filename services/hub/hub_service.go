package hub

import (
	"github.com/RouteHub-Link/routehub-service-graphql/clients/mq"
	"github.com/RouteHub-Link/routehub-service-graphql/database"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	"github.com/google/uuid"
)

type HubService struct {
	MQC *mq.MQTTClient
}

func NewHubService(platform database_models.Platform) (*HubService, error) {
	mqc, err := mq.NewMQTTClient(platform.TCPAddr)
	if err != nil {
		return nil, err
	}

	rc := HubService{
		MQC: mqc,
	}
	return &rc, nil
}

func NewHubServiceFromPlatformId(platformId uuid.UUID) (*HubService, error) {
	var platform database_models.Platform
	database.DB.Where("id = ?", platformId).Select("tcp_addr").Find(&platform)

	mqc, err := mq.NewMQTTClient(platform.TCPAddr)
	if err != nil {
		return nil, err
	}

	rc := HubService{
		MQC: mqc,
	}
	return &rc, nil

}
