package database

import (
	auth_casbin "github.com/RouteHub-Link/routehub-service-graphql/auth/casbin"
	"github.com/RouteHub-Link/routehub-service-graphql/auth/policies"
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	database_relations "github.com/RouteHub-Link/routehub-service-graphql/database/relations"
	database_types "github.com/RouteHub-Link/routehub-service-graphql/database/types"
	"github.com/google/uuid"
)

func Seed() {
	seedData := config.Database.Seed

	var user_count int64
	DB.Model(&database_models.User{}).Count(&user_count)

	var user *database_models.User
	var user2 *database_models.User
	if user_count == 0 {
		// seeder1@r4l.cc
		// seeder2@r4l.cc
		// passwords;
		// L-5sEaqufG4.p4V

		user = &database_models.User{
			ID: uuid.MustParse("927bb153-ed0a-4686-b8d5-5f1ced408ae4"),
		}

		user2 = &database_models.User{
			ID: uuid.MustParse("927bb153-ed0a-4686-b8d5-5f1ced408ae5"),
		}

		if seedData.Admins != nil {
			admins := *seedData.Admins
			user.Subject = admins[0].Subject
			user2.Subject = admins[1].Subject
		}

		DB.Create(user)
		DB.Create(user2)
	}

	var organization_count int64
	DB.Model(&database_models.Organization{}).Count(&organization_count)

	var organization *database_models.Organization
	if organization_count == 0 {
		organization = &database_models.Organization{
			ID:          uuid.MustParse("4e3b2a19-603c-4537-889e-04e64b3a8168"),
			Name:        seedData.Organization.Name,
			Website:     seedData.Organization.Url,
			Description: seedData.Organization.Description,
			Location:    "Web",
			Industry: []*database_types.Industry{
				{Name: "Software Development"}, {Name: "Information Technology"},
			},
			SocialMedias: []*database_types.SocialMedia{{Name: "LinkedIn", URL: "https://www.linkedin.com/test", Icon: "LinkedIn"}},
		}

		DB.Create(organization)
	}

	var user_organization_count int64
	DB.Model(&database_relations.OrganizationUser{}).Count(&user_organization_count)

	if user_organization_count == 0 {

		DB.First(&user)
		DB.First(&organization)

		user_Organization := &database_relations.OrganizationUser{
			ID:             uuid.New(),
			UserID:         user.ID,
			OrganizationID: organization.ID,
			Permissions:    database_enums.AllOrganizationPermission,
		}

		user2_Organization := &database_relations.OrganizationUser{
			ID:             uuid.New(),
			UserID:         user2.ID,
			OrganizationID: organization.ID,
			Permissions:    database_enums.AllOrganizationPermission,
		}

		DB.Create(user_Organization)
		DB.Create(user2_Organization)

		policies.NewPolicyBuilder(auth_casbin.CasbinEnforcer, user.ID, "allow").
			OrganizationPlatformCreate(organization.ID).
			OrganizationDelete(organization.ID).
			OrganizationRead(organization.ID).
			OrganizationUpdate(organization.ID).
			OrganizationUserInvite(organization.ID)

		policies.NewPolicyBuilder(auth_casbin.CasbinEnforcer, user2.ID, "allow").
			OrganizationPlatformCreate(organization.ID).
			OrganizationDelete(organization.ID).
			OrganizationRead(organization.ID).
			OrganizationUpdate(organization.ID).
			OrganizationUserInvite(organization.ID)

	}

	var domain_count int64
	domain := &database_models.Domain{}
	DB.Model(&database_models.Domain{}).Count(&domain_count)

	if domain_count == 0 {
		domain = &database_models.Domain{
			ID:             uuid.MustParse("71a61b35-b143-4ba3-9345-22f829eedfd5"),
			Name:           seedData.Domain.Name,
			OrganizationId: organization.ID,
			URL:            seedData.Domain.Url,
			State:          database_enums.StatusStateActive,
		}

		DB.Create(domain)
	}

	var platform_count int64
	DB.Model(&database_models.Platform{}).Count(&platform_count)

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
			Name:              "First Platform",
			DomainId:          domain.ID,
			RedirectionChoice: database_enums.RedirectionOptionsTimed,
			OpenGraph:         og,
			CreatedBy:         user.ID,
			Status:            database_enums.StatusStateActive,
		}

		DB.Create(&platform)

		organization_relation := database_relations.OrganizationPlatform{
			ID:             uuid.New(),
			OrganizationID: organization.ID,
			PlatformID:     platform.ID,
		}

		DB.Create(&organization_relation)

		platform_user := database_relations.PlatformUser{
			ID:          uuid.New(),
			UserID:      user.ID,
			PlatformID:  platform.ID,
			Permissions: database_enums.AllPlatformPermission,
		}

		DB.Create(&platform_user)

		policies.NewPolicyBuilder(auth_casbin.CasbinEnforcer, user.ID, "allow").
			PlatformRead(platform.ID).
			PlatformUpdate(platform.ID)
	}

	var link_count int64
	DB.Model(&database_models.Link{}).Count(&link_count)

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
			State:             database_enums.StatusStateActive,
			RedirectionChoice: database_enums.RedirectionOptionsTimed,
			OpenGraph:         og,
		}

		DB.Create(link)
	}
}
