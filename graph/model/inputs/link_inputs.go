package graph_inputs

import (
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_types "github.com/RouteHub-Link/routehub-service-graphql/database/types"
	"github.com/google/uuid"
)

type LinkUpdateInput struct {
	LinkID             uuid.UUID                          `json:"linkId"`
	Target             *string                            `json:"target,omitempty"`
	Path               *string                            `json:"path,omitempty"`
	RedirectionOptions *database_enums.RedirectionOptions `json:"redirectionOptions,omitempty"`
	LinkContent        *database_types.LinkContent        `json:"LinkContent,omitempty"`
}
