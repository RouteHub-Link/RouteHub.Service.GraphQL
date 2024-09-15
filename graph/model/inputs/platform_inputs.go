package graph_inputs

import (
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_types "github.com/RouteHub-Link/routehub-service-graphql/database/types"
	"github.com/RouteHub-Link/routehub-service-graphql/graph/model"
	"github.com/google/uuid"
)

type PlatformCreateInput struct {
	OrganizationID      uuid.UUID                           `json:"organizationId"`
	DomainID            uuid.UUID                           `json:"domainId"`
	Name                string                              `json:"name"`
	PlatformDescription *database_types.PlatformDescription `json:"openGraph"`
	RedirectionChoice   database_enums.RedirectionOptions   `json:"redirectionChoice"`
	Templates           []model.TemplateInput               `json:"templates,omitempty"`
}

type PlatformUpdateInput struct {
	PlatformID          uuid.UUID                           `json:"platformId"`
	OrganizationID      uuid.UUID                           `json:"organizationId"`
	Name                string                              `json:"name,omitempty"`
	PlatformDescription *database_types.PlatformDescription `json:"platformDescription,omitempty"`
	RedirectionChoice   *database_enums.RedirectionOptions  `json:"redirectionChoice,omitempty"`
	Templates           []model.TemplateInput               `json:"templates,omitempty"`
	Status              database_enums.StatusState          `json:"status"`
}
