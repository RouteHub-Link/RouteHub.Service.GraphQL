package database

import (
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	database_relations "github.com/RouteHub-Link/routehub-service-graphql/database/relations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(
		postgres.Open("host=127.0.0.1 user=postgres password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=UTC"),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		panic(err)
	}

	migrate(db)
	DB = db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		&database_models.User{},
		&database_models.Organization{},
		&database_relations.UserOrganization{},
		&database_relations.UserInvite{},
		&database_relations.PlatformUser{},
		&database_models.Domain{},
		&database_models.Platform{},
	)
}
