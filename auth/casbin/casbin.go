package auth_casbin

import (
	"log"
	"log/slog"
	"os"
	"sync"

	"github.com/RouteHub-Link/routehub-service-graphql/auth/policies"
	"github.com/RouteHub-Link/routehub-service-graphql/config"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	gormadapter "github.com/casbin/gorm-adapter/v3"

	"github.com/google/uuid"
)

var (
	CasbinAdapter  persist.Adapter
	CasbinEnforcer *casbin.SyncedCachedEnforcer
	onceLogger     sync.Once
	onceAdapter    sync.Once
	onceEnforcer   sync.Once
	casbinLogger   *slog.Logger
	casbinSlog     *CasbinSlogLogger
	level          slog.Level
)

type CasbinConfigurer struct {
	CasbinConfig   config.CasbinConfig
	databaseConfig config.DatabaseConfig
}

func NewCasbinConfigurer(casbinConfig config.CasbinConfig, databaseConfig config.DatabaseConfig) CasbinConfigurer {
	cc := CasbinConfigurer{CasbinConfig: casbinConfig, databaseConfig: databaseConfig}

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

		casbinLogger = slog.New(slog.NewJSONHandler(os.Stdout, opts))
		casbinSlog = NewCasbinSlogLogger(casbinLogger)
	})

	if casbinLogger == nil {
		log.Fatal("Casbin Logger is nil")
	}

	return casbinLogger
}

func (cc CasbinConfigurer) getAdapter() persist.Adapter {
	onceAdapter.Do(func() {
		dsn := cc.databaseConfig.GetPostgreDSN()

		a, _ := gormadapter.NewAdapter(config.Postgres.String(), dsn, true)
		CasbinAdapter = a
	})

	if CasbinAdapter == nil {
		log.Fatal("Casbin Adapter is nil")
	}

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
		e, err := casbin.NewSyncedCachedEnforcer(cc.CasbinConfig.Model, cc.getAdapter())
		if err != nil {
			log.Fatal(err)
		}

		e.SetLogger(casbinSlog)
		e.EnableLog(true)
		CasbinEnforcer = e
	})

	if CasbinEnforcer == nil {
		log.Fatal("Casbin Enforcer is nil")
	}

	return CasbinEnforcer
}
