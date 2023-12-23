package directives

import (
	"context"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/RouteHub-Link/routehub-service-graphql/auth"
	"github.com/RouteHub-Link/routehub-service-graphql/database"
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	database_relations "github.com/RouteHub-Link/routehub-service-graphql/database/relations"
	"github.com/google/uuid"
	"github.com/vektah/gqlparser/gqlerror"
)

func PlatformPermissionDirectiveHandler(ctx context.Context, obj interface{}, next graphql.Resolver, permission database_enums.PlatformPermission) (res interface{}, err error) {
	userSession := auth.ForContext(ctx)
	db := database.DB
	if userSession == nil {
		return nil, gqlerror.Errorf("Access Denied")
	}

	platformUUID, err := getPlatformId(obj)
	if err != nil {
		return nil, err
	}

	var platformUser database_relations.PlatformUser
	err = db.Where("user_id = ? AND platform_id = ?", userSession.ID, platformUUID).First(&platformUser).Error
	if err != nil {
		return nil, gqlerror.Errorf("Access Denied")
	}

	log.Printf("\n \norganizationUser: %+v\narg permission : %+v", platformUser, permission)

	for _, organization_user_permission := range platformUser.Permissions {
		if organization_user_permission == permission {
			return next(ctx)
		}
	}
	return nil, gqlerror.Errorf("Access Denied")
}

func getPlatformId(obj interface{}) (*uuid.UUID, error) {
	platformId, platformIdOk := obj.(map[string]interface{})["platformId"].(string)
	if platformIdOk {
		uid := (uuid.MustParse(platformId))
		return &uid, nil
	}

	domainId, domainIdOk := obj.(map[string]interface{})["domainId"].(string)
	if domainIdOk {
		domainPlatformId, err := platformIdFromDomain(domainId)
		if err != nil {
			return nil, err
		}
		return domainPlatformId, nil
	}

	linkId, linkIdOk := obj.(map[string]interface{})["linkId"].(string)
	if linkIdOk {
		linkPlatformId, err := platformIdFromLink(linkId)
		if err != nil {
			return nil, err
		}
		return linkPlatformId, nil
	}

	return nil, gqlerror.Errorf("platformId, domainId or linkId not found in obj")

}

func platformIdFromLink(linkId string) (platformId *uuid.UUID, err error) {
	db := database.DB

	link := &database_models.Link{}
	err = db.Where("id = ?", linkId).First(&link).Error
	if err != nil {
		return nil, gqlerror.Errorf("link cannot be found from linkId")
	}

	return &link.PlatformID, nil
}

func platformIdFromDomain(domainId string) (platformId *uuid.UUID, err error) {
	db := database.DB

	platform := &database_models.Platform{}
	err = db.Where("domain_id = ?", domainId).First(&platform).Error
	if err != nil {
		return nil, gqlerror.Errorf("domain cannot be found from domainId")
	}

	return &platform.ID, nil
}
