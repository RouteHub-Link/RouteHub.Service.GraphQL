package graph_inputs

import (
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	"github.com/cloudmatelabs/gorm-gqlgen-relay/relay"
)

type PageInfo = relay.PageInfo

type LinkEdge = relay.Edge[database_models.Link]
type LinkConnection = relay.Connection[database_models.Link]
