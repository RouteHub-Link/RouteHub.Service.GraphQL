package directives

import "github.com/RouteHub-Link/routehub-service-graphql/graph"

func Assign(config *graph.Config) {
	config.Directives.Auth = AuthDirectiveHandler
	config.Directives.OrganizationPermission = OrganizationPermissionDirectiveHandler
	config.Directives.PlatformPermission = PlatformPermissionDirectiveHandler
	config.Directives.DomainURLCheck = DomainURLCheckDirectiveHandler
}
