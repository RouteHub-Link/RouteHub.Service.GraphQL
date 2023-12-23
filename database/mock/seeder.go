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

	var user *database_models.User
	if user_count == 0 {
		pass := "admin"
		mail := "runaho@r4l.com"
		hash, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
		password := string(hash)

		user = &database_models.User{
			ID:           uuid.MustParse("927bb153-ed0a-4686-b8d5-5f1ced408ae4"),
			Email:        mail,
			PasswordHash: password,
			Fullname:     "Runaho",
			Verified:     true,
		}

		database.DB.Create(user)
	}

	var organization_count int64
	database.DB.Model(&database_models.Organization{}).Count(&organization_count)

	var organization *database_models.Organization
	if organization_count == 0 {
		organization = &database_models.Organization{
			ID:          uuid.MustParse("4e3b2a19-603c-4537-889e-04e64b3a8168"),
			Name:        "RouteHub",
			Website:     "https://routehub.link",
			Description: "RouteHub is a platform that connects Organizations and talents in the tech industry.",
			Location:    "Jakarta, Indonesia",
			Industry: []*database_types.Industry{
				{Name: "Software Development"}, {Name: "Information Technology"},
			},
			SocialMedias: []*database_types.SocialMedia{{Name: "LinkedIn", URL: "https://www.linkedin.com/test", Icon: "LinkedIn"}},
		}

		database.DB.Create(organization)
	}

	var user_organization_count int64
	database.DB.Model(&database_relations.OrganizationUser{}).Count(&user_organization_count)

	if user_organization_count == 0 {

		database.DB.First(&user)
		database.DB.First(&organization)

		user_Organization := &database_relations.OrganizationUser{
			ID:             uuid.New(),
			UserID:         user.ID,
			OrganizationID: organization.ID,
			Permissions:    database_enums.AllOrganizationPermission,
		}

		database.DB.Create(user_Organization)
	}

	var domain_count int64
	domain := &database_models.Domain{}
	database.DB.Model(&database_models.Domain{}).Count(&domain_count)

	if domain_count == 0 {
		domain = &database_models.Domain{
			ID:             uuid.MustParse("71a61b35-b143-4ba3-9345-22f829eedfd5"),
			Name:           "RouteHub Public Shortener",
			OrganizationId: organization.ID,
			URL:            "https://s.r4l.cc",
		}

		database.DB.Create(domain)
	}

	var platform_count int64
	database.DB.Model(&database_models.Platform{}).Count(&platform_count)

	var platform *database_models.Platform
	if platform_count == 0 {
		og := &database_types.OpenGraph{}
		ogData := `{
			"title": "Route Hub Shortener",
			"description": "R4L Route Hub B2B Shortener",
			"alternateImage": "cdn.link.r4l.cc/alternate",
			"favIcon": "cdn.link.r4l.cc/favicon",
			"image": "cdn.link.r4l.cc/image",
			"locale": "tr-TR",
			"siteName": "Shorten Route for Links CC",
			"type": "website",
			"url": "https://s.r4l.cc",
			"x": {
			  "creator": "@runaho",
			  "card": "R4L Shorten Link",
			  "description": "Route Hub Link Shortener B2B Main",
			  "image": "s.r4l.cc/image",
			  "site": "https://routehub.link",
			  "title": "Route Hub Shortener",
			  "type": "website",
			  "url": "https://s.r4l.cc"
			}
		}`
		og.ParseFromJson(ogData)

		platform = &database_models.Platform{
			ID:                uuid.MustParse("12058bdf-8940-43b3-bd90-13487e4c8fc4"),
			Name:              "RouteHub Public Shortener",
			DomainId:          domain.ID,
			RedirectionChoice: database_enums.RedirectionOptionsTimed,
			OpenGraph:         og,
			CreatedBy:         user.ID,
			Status:            database_enums.StatusStateActive,
		}

		database.DB.Create(&platform)

		organization_relation := database_relations.OrganizationPlatform{
			ID:             uuid.New(),
			OrganizationID: organization.ID,
			PlatformID:     platform.ID,
		}

		database.DB.Create(&organization_relation)

		platform_user := database_relations.PlatformUser{
			ID:          uuid.New(),
			UserID:      user.ID,
			PlatformID:  platform.ID,
			Permissions: database_enums.AllPlatformPermission,
		}

		database.DB.Create(&platform_user)
	}

	var link_count int64
	database.DB.Model(&database_models.Link{}).Count(&link_count)

	if link_count == 0 {
		ogData := `{
			"title": "My 10 Day GO Learning Journey",
			"description": "In this blog i will share my experiences on GO & also i will give you whole 10 day learning exp.",
			"alternateImage": "cdn.link.r4l.cc/alternate",
			"favIcon": "cdn.link.r4l.cc/favicon",
			"image": "cdn.link.r4l.cc/image",
			"locale": "tr-TR",
			"siteName": "Shorten Route for Links CC",
			"type": "blog",
			"url": "https://blog.guneskorkmaz.net/my-go-learning-journey",
			"x": {
			  "creator": "@runaho",
			  "card": "R4L Shorten Link",
			  "description": "Route Hub Link Shortener B2B Main",
			  "image": "s.r4l.cc/image",
			  "site": "https://routehub.link",
			  "title": "Route Hub Shortener",
			  "type": "website",
			  "url": "https://blog.guneskorkmaz.net/my-go-learning-journey"
			}
		  }`

		og := &database_types.OpenGraph{}
		og.ParseFromJson(ogData)

		link := &database_models.Link{
			ID:                uuid.MustParse("c7d3a1e0-0c4d-4c1f-8f2c-6c4e6d6f7f3c"),
			CreatedBy:         user.ID,
			PlatformID:        platform.ID,
			Target:            "https://blog.guneskorkmaz.net/my-go-learning-journey",
			Path:              "my-go-learning-journey",
			Status:            database_enums.StatusStateActive,
			RedirectionChoice: database_enums.RedirectionOptionsTimed,
			OpenGraph:         og,
		}

		database.DB.Create(link)
	}
}
