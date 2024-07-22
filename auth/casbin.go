package auth

import (
	"context"
	"log/slog"
	"os"
	"sync"

	"github.com/RouteHub-Link/routehub-service-graphql/auth/policies"
	"github.com/RouteHub-Link/routehub-service-graphql/config"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	mongodbadapter "github.com/casbin/mongodb-adapter/v3"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/event"
	mongooptions "go.mongodb.org/mongo-driver/mongo/options"
)

var (
	CasbinAdapter  persist.Adapter
	CasbinEnforcer *casbin.SyncedCachedEnforcer
	onceAdapter    sync.Once
	onceEnforcer   sync.Once
)

type CasbinConfigurer struct {
	CasbinConfig config.CasbinConfig
}

func NewCasbinConfigurer(casbinConfig config.CasbinConfig) CasbinConfigurer {
	cc := CasbinConfigurer{CasbinConfig: casbinConfig}
	cc.getAdapter()
	cc.getEnforcer()
	return cc
}

func GetPolicyBuilder(userId uuid.UUID, e *casbin.SyncedCachedEnforcer) *policies.PolicyBuilder {
	return policies.NewPolicyBuilder(e, userId, "allow")
}

func (cc CasbinConfigurer) getAdapter() persist.Adapter {
	onceAdapter.Do(func() {

		opts := &slog.HandlerOptions{
			Level: slog.Level(-8),
		}

		logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

		ctx := context.Background()

		mongoClientOption := mongooptions.Client().ApplyURI(cc.CasbinConfig.Mongo.URI)

		monitor := &event.CommandMonitor{
			Started: func(_ context.Context, e *event.CommandStartedEvent) {
				if e.CommandName != "endSessions" {
					logger.Log(ctx, opts.Level.Level(), "MongoDB Trace Command", e.CommandName, e.Command)
				}
			},
		}

		mongoClientOption.SetMonitor(monitor)

		a, err := mongodbadapter.NewAdapterWithCollectionName(mongoClientOption, cc.CasbinConfig.Mongo.Database, cc.CasbinConfig.Mongo.Collection)
		if err != nil {
			panic(err)
		}
		CasbinAdapter = a

	})

	return CasbinAdapter
}

func (cc CasbinConfigurer) initTestPolicy(e *casbin.SyncedCachedEnforcer) (*casbin.SyncedCachedEnforcer, error) {
	userId := uuid.New()
	organizationId := uuid.New()
	platformId := uuid.New()

	pb := policies.NewPolicyBuilder(e, userId, "allow")

	pb.EnforceWhenAdded(true).
		OrganizationRead(organizationId).
		OrganizationUpdate(organizationId).
		OrganizationDelete(organizationId).
		OrganizationPlatformCreate(organizationId).
		OrganizationUserInvite(organizationId).
		PlatformRead(platformId).
		PlatformUpdate(platformId).
		PlatformLinkCreate(platformId).
		PlatformLinkDelete(platformId).
		PlatformLinkRead(platformId).
		PlatformLinkUpdate(platformId)

	return e, nil
}

func (cc CasbinConfigurer) getEnforcer() *casbin.SyncedCachedEnforcer {
	onceEnforcer.Do(func() {
		//e, _ := casbin.NewEnforcer(cc.CasbinConfig.Model, cc.getAdapter())
		e, _ := casbin.NewSyncedCachedEnforcer(cc.CasbinConfig.Model, cc.getAdapter())
		CasbinEnforcer = e

		//_, _ = cc.initTestPolicy(e)
	})

	return CasbinEnforcer
}
