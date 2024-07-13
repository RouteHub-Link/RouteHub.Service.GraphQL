package directives

import (
	"context"
	"log"
	"reflect"

	"github.com/99designs/gqlgen/graphql"
	"github.com/RouteHub-Link/routehub-service-graphql/auth"
	database_enums "github.com/RouteHub-Link/routehub-service-graphql/database/enums"
	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	"github.com/google/uuid"
	"github.com/vektah/gqlparser/gqlerror"
)

func OrganizationPermissionDirectiveHandler(ctx context.Context, obj interface{}, next graphql.Resolver, permission database_enums.OrganizationPermission) (res interface{}, err error) {
	userSession := auth.ForContext(ctx)
	if userSession == nil {
		return nil, gqlerror.Errorf("Access Denied")
	}

	var organizationId string
	fc := graphql.GetFieldContext(ctx)

	if obj == nil {
		return next(ctx)
	}

	if fc.Parent.Object == "Mutation" {
		_orgId, ok := obj.(map[string]interface{})["organizationId"].(string)
		if !ok {
			return nil, gqlerror.Errorf("organizationId not found in obj")

		}
		organizationId = _orgId
	} else {
		reflectFields := reflect.ValueOf(obj).Elem()
		if reflectFields.Type() != reflect.TypeOf(database_models.Organization{}) {
			return next(ctx)
		}

		reflectField := reflectFields.FieldByName("ID")
		if reflectField.IsValid() {
			bytes := reflectField.Bytes()
			organizationId = uuid.Must(uuid.FromBytes(bytes)).String()
		}

		if organizationId == "" {
			return nil, gqlerror.Errorf("organizationId not found in obj")
		}
	}

	e := auth.CasbinEnforcer
	hasPermission, exp, err := e.EnforceEx(userSession.ID.String(), organizationId, permission.String())
	log.Printf("\nOrganizationPermissionDirectiveHandler EnforceEX;\nres: %+v\nexp: %+v\nerr: %+v\n\n", hasPermission, exp, err)

	if hasPermission {
		return next(ctx)
	}

	if err != nil {
		return nil, err
	}

	return nil, gqlerror.Errorf("Access Denied")
}
