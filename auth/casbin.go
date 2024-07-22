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
	CasbinAdapter     persist.Adapter
	CasbinEnforcer    *casbin.SyncedCachedEnforcer
	onceLogger        sync.Once
	onceAdapter       sync.Once
	onceEnforcer      sync.Once
	casbinMongoLogger *slog.Logger
	casbinSlog        *CasbinSlogLogger
	level             slog.Level
)

type CasbinConfigurer struct {
	CasbinConfig config.CasbinConfig
}

func NewCasbinConfigurer(casbinConfig config.CasbinConfig) CasbinConfigurer {
	cc := CasbinConfigurer{CasbinConfig: casbinConfig}

	cc.getLogger()
	cc.getAdapter()
	cc.getEnforcer()
	return cc
}

func GetPolicyBuilder(userId uuid.UUID, e *casbin.SyncedCachedEnforcer) *policies.PolicyBuilder {
	return policies.NewPolicyBuilder(e, userId, "allow")
}

func (cc CasbinConfigurer) getLogger() *slog.Logger {
	onceLogger.Do(func() {
		var err = level.UnmarshalText([]byte(cc.CasbinConfig.LogLevel))
		if err != nil {
			level = slog.LevelInfo
		}

		opts := &slog.HandlerOptions{
			Level: slog.Level(level),
		}

		casbinMongoLogger = slog.New(slog.NewJSONHandler(os.Stdout, opts))
		casbinSlog = NewCasbinSlogLogger(casbinMongoLogger)
	})

	return casbinMongoLogger
}

func (cc CasbinConfigurer) getAdapter() persist.Adapter {
	onceAdapter.Do(func() {
		ctx := context.Background()

		mongoClientOption := mongooptions.Client().ApplyURI(cc.CasbinConfig.Mongo.URI)

		monitor := &event.CommandMonitor{
			Started: func(_ context.Context, e *event.CommandStartedEvent) {
				if e.CommandName != "endSessions" {
					casbinMongoLogger.LogAttrs(ctx, slog.LevelDebug, "Casbin MongoDB Command", slog.String("Database", e.DatabaseName), slog.String("CommandName", e.CommandName), slog.Any("Command", e.Command))
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
		e, _ := casbin.NewSyncedCachedEnforcer(cc.CasbinConfig.Model, cc.getAdapter())
		e.SetLogger(casbinSlog)
		e.EnableLog(true)
		CasbinEnforcer = e
	})

	return CasbinEnforcer
}
