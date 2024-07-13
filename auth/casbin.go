package auth

import (
	"sync"

	"github.com/RouteHub-Link/routehub-service-graphql/auth/policies"
	"github.com/RouteHub-Link/routehub-service-graphql/config"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	mongodbadapter "github.com/casbin/mongodb-adapter/v3"
	"github.com/google/uuid"
	mongooptions "go.mongodb.org/mongo-driver/mongo/options"
)

var (
	adp           persist.Adapter
	onceConfigure sync.Once
)

type CasbinConfigurer struct {
	CasbinConfig config.CasbinConfig
}

func (cc CasbinConfigurer) getAdapter() persist.Adapter {
	onceConfigure.Do(func() {

		mongoClientOption := mongooptions.Client().ApplyURI(cc.CasbinConfig.Mongo.URI)

		a, err := mongodbadapter.NewAdapterWithCollectionName(mongoClientOption, cc.CasbinConfig.Mongo.Database, cc.CasbinConfig.Mongo.Collection)
		if err != nil {
			panic(err)
		}
		adp = a

	})

	return adp
}

func (cc CasbinConfigurer) initTestPolicy(e *casbin.Enforcer) (*casbin.Enforcer, error) {
	userId := uuid.New()
	organizationId := uuid.New()
	platformId := uuid.New()

	pb := policies.NewPolicyBuilder(e, userId, organizationId, "allow")

	pb.EnforceWhenAdded(true).
		OrganizationRead(organizationId).
		OrganizationUpdate(organizationId).
		OrganizationDelete(organizationId).
		OrganizationPlatformCreate(organizationId).
		OrganizationUserInvite(organizationId).
		PlatformRead(platformId).
		PlatformUpdate(platformId).
		LinkCreate(organizationId).
		LinkRead(organizationId).
		LinkUpdate(organizationId).
		LinkDelete(organizationId)

	return e, nil
}

func (cc CasbinConfigurer) GetEnforcer() *casbin.Enforcer {
	e, _ := casbin.NewEnforcer(cc.CasbinConfig.Model, cc.getAdapter())

	//policy, err := e.GetPolicy()
	//lenPolicy := len(policy)
	//if err != nil || lenPolicy == 0 {

	_, _ = cc.initTestPolicy(e)

	return e
}
