package directives

import (
	"context"
	"reflect"

	"github.com/99designs/gqlgen/graphql"
	"github.com/RouteHub-Link/routehub-service-graphql/auth"
	auth_casbin "github.com/RouteHub-Link/routehub-service-graphql/auth/casbin"
	"github.com/RouteHub-Link/routehub-service-graphql/database"
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	"github.com/google/uuid"
	"github.com/vektah/gqlparser/gqlerror"
)

func PlatformPermissionDirectiveHandler(ctx context.Context, obj interface{}, next graphql.Resolver, permission database_enums.PlatformPermission) (res interface{}, err error) {
	userSession := auth.ForContext(ctx)
	var platformId string
	if userSession == nil {
		return nil, gqlerror.Errorf("Access Denied")
	}

	fc := graphql.GetFieldContext(ctx)

	if obj == nil {
		return next(ctx)
	}

	if fc.Parent.Object == "Mutation" {
		uuid, err := getPlatformId(obj)
		if err != nil {
			return nil, err
		}

		platformId = uuid.String()
	} else {
		reflectFields := reflect.ValueOf(obj).Elem()
		if reflectFields.Type() != reflect.TypeOf(database_models.Platform{}) {
			return next(ctx)
		}

		reflectField := reflectFields.FieldByName("ID")
		if reflectField.IsValid() {
			bytes := reflectField.Bytes()
			platformId = uuid.Must(uuid.FromBytes(bytes)).String()
		}

		if platformId == "" {
			return nil, gqlerror.Errorf("platformId not found in obj")
		}
	}

	e := auth_casbin.CasbinEnforcer
	hasPermission, _, err := e.EnforceEx(userSession.ID.String(), platformId, permission.String())

	if hasPermission {
		return next(ctx)
	} else if err != nil {
		return nil, err
	}

	return nil, gqlerror.Errorf("Access Denied")
}

func getPlatformId(obj interface{}) (*uuid.UUID, error) {
	platform, platformOk := obj.(*database_models.Platform)
	if platformOk {
		return &platform.ID, nil
	}

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
