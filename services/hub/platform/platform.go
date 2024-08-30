package platform

import (
	"github.com/RouteHub-Link/routehub-service-graphql/services/hub/layout"
	"github.com/RouteHub-Link/routehub-service-graphql/services/hub/redirection"
	"github.com/RouteHub-Link/routehub-service-graphql/services/hub/status"
)

type Platform struct {
	Name               string             `gorm:"type:varchar(255);not null;"`
	Slug               string             `gorm:"type:varchar(255);not null;"`
	DefaultRedirection redirection.Option `gorm:"type:varchar(255);not null;"`
	Status             status.State       `gorm:"type:varchar(255);not null;"`
	LayoutDescription  LayoutDescription  `gorm:"foreignKey:PlatformID;"`
}

type LayoutDescription struct {
	MetaDescription   layout.MetaDescription
	NavbarDescription layout.NavbarDescription
	FooterDescription layout.FooterDescription
}
