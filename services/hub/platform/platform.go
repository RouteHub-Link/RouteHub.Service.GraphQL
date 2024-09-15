package platform

import (
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	database_types "github.com/RouteHub-Link/routehub-service-graphql/database/types"
	"github.com/RouteHub-Link/routehub-service-graphql/services/hub/redirection"
	"github.com/RouteHub-Link/routehub-service-graphql/services/hub/status"
)

type Platform struct {
	Name               string
	Slug               string
	DefaultRedirection redirection.Option
	Status             status.State
	LayoutDescription  *database_types.PlatformDescription
}

func (p *Platform) MapFromDatabasePlatform(dbPlatform database_models.Platform) {
	p.Name = dbPlatform.Name
	p.Slug = dbPlatform.Slug

	switch dbPlatform.RedirectionChoice {
	case database_enums.RedirectionOptionsTimed:
		p.DefaultRedirection = redirection.OptionTimed
	case database_enums.RedirectionOptionsDirectHTTPRedirect:
		p.DefaultRedirection = redirection.OptionDirectHTTP
	case database_enums.RedirectionOptionsConfirmRedirect:
		p.DefaultRedirection = redirection.OptionConfirm
	case database_enums.RedirectionOptionsNotAutoRedirect:
		p.DefaultRedirection = redirection.OptionNotAuto
	case database_enums.RedirectionOptionsCustom:
		p.DefaultRedirection = redirection.OptionCustom
	}

	p.LayoutDescription = dbPlatform.PlatformDescription
}
