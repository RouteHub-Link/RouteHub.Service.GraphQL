package database_mock

import (
	"github.com/RouteHub-Link/routehub-service-graphql/database"
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	database_relations "github.com/RouteHub-Link/routehub-service-graphql/database/relations"
	database_types "github.com/RouteHub-Link/routehub-service-graphql/database/types"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	var user_count int64
	database.DB.Model(&database_models.User{}).Count(&user_count)

	if user_count == 0 {
		pass := "admin"
		mail := "runaho@r4l.com"
		hash, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
		password := string(hash)

		user := &database_models.User{
			ID:           uuid.New(),
			Email:        mail,
			PasswordHash: password,
			Fullname:     "Runaho",
			Verified:     true,
		}

		database.DB.Create(user)
	}

	var organization_count int64
	database.DB.Model(&database_models.Organization{}).Count(&organization_count)

	if organization_count == 0 {
		Organization := &database_models.Organization{
			ID:          uuid.New(),
			Name:        "RouteHub",
			Website:     "https://routehub.link",
			Description: "RouteHub is a platform that connects Organizations and talents in the tech industry.",
			Location:    "Jakarta, Indonesia",
			Industry: []*database_types.Industry{
				{Name: "Software Development"}, {Name: "Information Technology"},
			},
			SocialMedias: []*database_types.SocialMedia{{Name: "LinkedIn", URL: "https://www.linkedin.com/test", Icon: "LinkedIn"}},
		}

		database.DB.Create(Organization)
	}

	var user_organization_count int64
	database.DB.Model(&database_relations.UserOrganization{}).Count(&user_organization_count)

	if user_organization_count == 0 {
		var user database_models.User
		var organization database_models.Organization

		database.DB.First(&user)
		database.DB.First(&organization)

		user_Organization := &database_relations.UserOrganization{
			ID:             uuid.New(),
			UserID:         user.ID,
			OrganizationID: organization.ID,
			Permissions:    database_enums.AllOrganizationPermission,
		}

		database.DB.Create(user_Organization)
	}
}
