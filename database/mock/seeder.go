package database_mock

import (
	"github.com/RouteHub-Link/routehub-service-graphql/database"
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

	var company_count int64
	database.DB.Model(&database_models.Company{}).Count(&company_count)

	if company_count == 0 {
		company := &database_models.Company{
			ID:          uuid.New(),
			Name:        "RouteHub",
			Website:     "https://routehub.link",
			Description: "RouteHub is a platform that connects companies and talents in the tech industry.",
			Location:    "Jakarta, Indonesia",
			Industry: []*database_types.Industry{
				{Name: "Software Development"}, {Name: "Information Technology"},
			},
			SocialMedias: []*database_types.SocialMedia{{Name: "LinkedIn", URL: "https://www.linkedin.com/test", Icon: "LinkedIn"}},
		}

		database.DB.Create(company)
	}

	var user_company_count int64
	database.DB.Model(&database_relations.UserCompany{}).Count(&user_company_count)

	if user_company_count == 0 {
		var user database_models.User
		var company database_models.Company

		database.DB.First(&user)
		database.DB.First(&company)

		user_company := &database_relations.UserCompany{
			UserID:    user.ID,
			CompanyID: company.ID,
		}

		database.DB.Create(user_company)
	}
}
