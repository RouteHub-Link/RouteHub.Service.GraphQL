package database

import (
	"log"

	auth_casbin "github.com/RouteHub-Link/routehub-service-graphql/auth/casbin"
	"github.com/RouteHub-Link/routehub-service-graphql/auth/policies"
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	database_relations "github.com/RouteHub-Link/routehub-service-graphql/database/relations"
	database_types "github.com/RouteHub-Link/routehub-service-graphql/database/types"
	"github.com/google/uuid"
)

func seed() {
	log.Println("Database Seed Called")
	seedData := config.Database.Seed

	var user_count int64
	DB.Model(&database_models.User{}).Count(&user_count)

	var user *database_models.User
	var user2 *database_models.User
	if user_count == 0 {
		// seeder1@r4l.cc
		// seeder2@r4l.cc
		// passwords;
		// L-5sEaqufG4.p4V1

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

		/*
			"alternateImage": "cdn.link.r4l.cc/alternate",
			"favIcon": "cdn.link.r4l.cc/favicon",
			"image": "cdn.link.r4l.cc/image",
			"site": "https://routehub.link",
			"url": "https://s.r4l.cc"
		*/

		brandImg := database_types.ImageDescription{
			SRC:    "https://example.com/icon.png",
			Alt:    "Example Platform",
			Height: "30px",
			Width:  "30px",
		}

		navbarItems := []database_types.NavbarItem{{Text: "Home", URL: "/", Target: "_self", Icon: "home"}, {Text: "About", URL: "/about", Target: "_self", Icon: "info", Dropdown: &[]database_types.NavbarItem{{Text: "Contact", URL: "/contact", Target: "_self", Icon: "contact_mail"}}}}
		navbarEndButtons := []database_types.NavbarButton{{Text: "Login", URL: "/login", Target: "_self", ColorClass: "is-secondary"}, {Text: "Sign Up", URL: "/signup", Target: "_self", ColorClass: "is-primary"}}
		navbar_description := database_types.NavbarDescription{BrandName: "Example", BrandURL: "https://example.com", BrandImg: &brandImg, StartItems: &navbarItems, EndButtons: &navbarEndButtons}

		socialMediaList := []database_types.ASocialMedia{{Icon: "facebook", Link: "https://www.facebook.com", Target: "_blank"}, {Icon: "twitter", Link: "https://www.twitter.com", Target: "_blank"}, {Icon: "instagram", Link: "https://www.instagram.com", Target: "_blank"}, {Icon: "linkedin", Link: "https://www.linkedin.com", Target: "_blank"}}
		socialMediaContainer := database_types.SocialMediaContainer{SocialMediaLinks: &socialMediaList, SocialMediaPeddingClass: "pt-5", SocialMediaSizeClass: "is-medium", SocialMediaColorClass: "has-text-white"}
		footer_description := database_types.FooterDescription{ShowRouteHubBranding: true, CompanyBrandingHtml: "<strong>Example Company</strong> <a href=''> Example Company</a> Has Rights of this site since 1111</strong>", SocialMediaContainer: &socialMediaContainer}

		platform_meta_description := database_types.MetaDescription{
			Title:         "Example Platform",
			FavIcon:       "https://example.com/icon.png",
			Description:   "This is an example platform",
			Locale:        "en-US",
			OGTitle:       "Example Platform",
			OGDescription: "This is an example platform",
			OGURL:         "https://example.com",
			OGSiteName:    "Example Platform",
			OGMetaType:    "website",
			OGLocale:      "en-US",
			OGBigImage:    "https://example.com/image",
			OGBigWidth:    "1200",
			OGBigHeight:   "630",
			OGSmallImage:  "https://example.com/image",
			OGSmallWidth:  "600",
			OGSmallHeight: "315",
			OGCard:        "summary_large_image",
			OGSite:        "https://example.com",
			OGType:        "website",
			OGCreator:     "@example",
		}

		platform_description := &database_types.PlatformDescription{
			MetaDescription:   platform_meta_description,
			NavbarDescription: navbar_description,
			FooterDescription: footer_description,
		}

		platform = &database_models.Platform{
			ID:                  uuid.MustParse("12058bdf-8940-43b3-bd90-13487e4c8fc4"),
			Name:                "Example Platform",
			DomainId:            domain.ID,
			RedirectionChoice:   database_enums.RedirectionOptionsTimed,
			PlatformDescription: platform_description,
			CreatedBy:           user.ID,
			Status:              database_enums.StatusStateActive,
			TCPAddr:             "localhost:1883",
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
			PlatformUpdate(platform.ID).
			PlatformLinkRead(platform.ID).
			PlatformLinkUpdate(platform.ID).
			PlatformLinkCreate(platform.ID)
	}

	var link_count int64
	DB.Model(&database_models.Link{}).Count(&link_count)

	if link_count == 0 {
		redirectionDelay := 5
		link_subtitle := "As described in RFC 2606 and RFC 6761, a number of domains such as example.com and example.org are maintained for documentation purposes."

		link := &database_models.Link{
			ID:                uuid.MustParse("c7d3a1e0-0c4d-4c1f-8f2c-6c4e6d6f7f3c"),
			CreatedBy:         user.ID,
			PlatformID:        platform.ID,
			Target:            "https://www.iana.org/help/example-domains",
			Path:              "example-domains",
			State:             database_enums.StatusStateActive,
			RedirectionChoice: database_enums.RedirectionOptionsTimed,
			LinkContent: &database_types.LinkContent{
				Title:              "Example Domains",
				Subtitle:           link_subtitle,
				RedirectionURLText: "IANA.ORG Example Domains",
				RedirectionDelay:   &redirectionDelay,
				MetaDescription: &database_types.MetaDescription{
					Title:         "Example Domains",
					FavIcon:       "https://bucket.r4l.cc/routehub-images/logoipsum-218.svg",
					Description:   link_subtitle,
					Locale:        "en-US",
					OGTitle:       "Example Domains",
					OGDescription: link_subtitle,
					OGURL:         "https://bucket.r4l.cc/routehub-images/logoipsum-218.svg",
					OGSiteName:    "Example Domains",
					OGMetaType:    "website",
					OGLocale:      "en-US",
					OGBigImage:    "https://bucket.r4l.cc/routehub-images/logoipsum-218.svg",
					OGBigWidth:    "1200",
					OGBigHeight:   "630",
					OGSmallImage:  "https://bucket.r4l.cc/routehub-images/logoipsum-218.svg",
					OGSmallWidth:  "600",
					OGSmallHeight: "315",
					OGCard:        "summary_large_image",
					OGSite:        "https://www.iana.org",
					OGType:        "website",
					OGCreator:     "@ExampleDomains",
				},
			},
		}

		DB.Create(link)
	}
}
