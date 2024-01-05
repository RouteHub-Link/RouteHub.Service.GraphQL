package database

import (
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	database_relations "github.com/RouteHub-Link/routehub-service-graphql/database/relations"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&database_models.User{},
		&database_models.Organization{},
		&database_relations.OrganizationUser{},
		&database_relations.OrganizationPlatform{},
		&database_relations.UserInvite{},
		&database_relations.PlatformUser{},
		&database_models.Domain{},
		&database_models.Platform{},
		&database_models.Link{},
		&database_models.LinkCrawl{},
	)
}
