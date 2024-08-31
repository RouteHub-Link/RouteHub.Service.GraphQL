package publish

import (
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	"github.com/RouteHub-Link/routehub-service-graphql/services/hub"
	"github.com/RouteHub-Link/routehub-service-graphql/services/hub/platform"
)

type PlatformPub struct {
	hubService *hub.HubService
}

func NewPlatformEvents(hubService *hub.HubService) *PlatformPub {
	return &PlatformPub{
		hubService: hubService,
	}
}

func (pp *PlatformPub) PubSet(p database_models.Platform) error {
	_platform := &platform.Platform{}
	_platform.MapFromDatabasePlatform(p)

	mqc := pp.hubService.MQC
	client := *mqc.Client

	token := client.Publish("platform/set", 3, true, _platform)
	token.WaitTimeout(mqc.Timeout)

	return token.Error()
}
