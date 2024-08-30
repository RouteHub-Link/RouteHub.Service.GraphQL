package publish

import (
	"encoding/json"

	"github.com/RouteHub-Link/routehub-service-graphql/clients/mq"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	"github.com/RouteHub-Link/routehub-service-graphql/services/hub"
	"github.com/RouteHub-Link/routehub-service-graphql/services/hub/link"
)

type LinkPub struct {
	hubService *hub.HubService
}

func NewLinkEvents(hubService *hub.HubService) *LinkPub {
	return &LinkPub{
		hubService: hubService,
	}
}

func (lp *LinkPub) PubSet(dbLink database_models.Link) error {
	_link := &link.Link{}
	_link.MapFromDatabaseLink(dbLink)
	jsonLink, err := json.Marshal(_link)
	if err != nil {
		return err
	}
	mqc := lp.hubService.MQC
	client := *mqc.Client
	token := client.Publish(mq.MQE_LINK_SET.AsTopic(), 3, true, jsonLink)
	token.WaitTimeout(mqc.Timeout)

	return token.Error()
}

func (lp *LinkPub) PubDel(dbLink database_models.Link) error {
	linkTarget := dbLink.Path
	client := *lp.hubService.MQC.Client
	token := client.Publish(mq.MQE_LINK_DEL.AsTopic(), 3, true, linkTarget)
	token.WaitTimeout(lp.hubService.MQC.Timeout)
	return token.Error()
}
