package auth

import (
	"log"
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
	// add the policy
	testUserUUID := uuid.New()
	testOrgUUID := uuid.New()
	testPlatformUUID := uuid.New()

	testPolicies := policies.NewOrganizationPolicies(testUserUUID, testOrgUUID, testPlatformUUID)
	log.Printf("testPolicies : %+v", testPolicies)

	userId := uuid.New()
	organizationId := uuid.New()
	platformId := uuid.New()

	e.SavePolicy()

	res, _ := e.Enforce(testUserUUID.String(), testOrgUUID.String(), testPlatformUUID.String(), "read")
	log.Printf("res : %v", res)

	// With Custom Builder
	pb := policies.NewPolicyBuilder(e, userId, organizationId, platformId)

	pb.OrganizationRead().
		OrganizationUpdate().
		OrganizationDelete().
		PlatformRead().
		PlatformUpdate().
		InvitationCreate().
		InvitationRead().
		InvitationDelete().
		LinkCreate().
		LinkRead().
		LinkDelete().
		LinkUpdate().
		GroupingPolicy().
		Build()

	pb.Enforce()

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
