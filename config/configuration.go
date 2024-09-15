package config

import (
	"context"
	"log"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type ApplicationConfig struct {
	GraphQL      GraphqlConfig  `koanf:"graphql"`
	Database     DatabaseConfig `koanf:"database"`
	RedisConfig  RedisConfig    `koanf:"redis"`
	Services     ServicesConfig `koanf:"services"`
	CasbinConfig CasbinConfig   `koanf:"casbin"`
	AuthConfig   AuthConfig     `koanf:"zitadel"`
	Host         string         `koanf:"host"`
}

type GraphqlConfig struct {
	Playground   bool                    `koanf:"playground"`
	Port         uint32                  `koanf:"port"`
	PortAsString string                  `koanf:"port"`
	Dataloader   GraphQlDataLoaderConfig `koanf:"dataloader"`
}

type GraphQlDataLoaderConfig struct {
	Cache bool          `koanf:"cache"`
	Wait  time.Duration `koanf:"wait"`
	Lrue  LrueConfig    `koanf:"lrue"`
}

type LrueConfig struct {
	Size   int           `koanf:"size"`
	Expire time.Duration `koanf:"expire"`
}

type DatabaseConfig struct {
	Host            string             `koanf:"host"`
	Port            uint32             `koanf:"port"`
	PortAsString    string             `koanf:"port"`
	User            string             `koanf:"user"`
	Password        string             `koanf:"password"`
	Database        string             `koanf:"database"`
	ApplicationName string             `koanf:"application_name"`
	Type            DatabaseTypeConfig `koanf:"type"`
	Seed            *Seed              `koanf:"seed"`
}
type DatabaseTypeConfig struct {
	Migrate  bool     `koanf:"migrate"`
	Seed     bool     `koanf:"seed"`
	Provider Provider `koanf:"provider"`
}

type RedisConfig struct {
	Addres string `koanf:"addr"`
}

type ServicesConfig struct {
	DomainUtilsHost string `koanf:"domain_utils_host"`
}

type CasbinConfig struct {
	Model    string `koanf:"model"`
	LogLevel string `koanf:"log_level"`
}

type AuthConfig struct {
	ClientID      string   `koanf:"client_id"`
	Callback      string   `koanf:"callback"`
	Scopes        []string `koanf:"scope"`
	AuthorizerURL string   `koanf:"authorizer_url"`
	TokenURL      string   `koanf:"token_url"`
	Issuer        string   `koanf:"issuer"`
	Domain        string   `koanf:"domain"`
	Port          string   `koanf:"port"`
	Insecure      bool     `koanf:"insecure"`
}

type CasbinMongoConfig struct {
	URI        string `koanf:"uri"`
	Database   string `koanf:"database"`
	Collection string `koanf:"collection"`
}

type ConfigurationService struct{}

var (
	_appConfig    *ApplicationConfig
	onceConfigure sync.Once
	conf          = koanf.Conf{
		Delim:       ".",
		StrictMerge: true,
	}
	logger *slog.Logger
	k      = koanf.NewWithConf(conf)
)

const (
	yamlPath = "config/config.yaml"
)

func (ConfigurationService) Get() *ApplicationConfig {
	onceConfigure.Do(func() {
		_appConfig = &ApplicationConfig{}
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
		configParsed := false
		if err := k.Load(file.Provider(yamlPath), yaml.Parser()); err != nil {
			logger.Log(context.Background(), slog.LevelError, "Error loading config.yaml", slog.String("error", err.Error()))
		} else {
			configParsed = koanfUnmarshall()
		}

		if !configParsed {
			configParsed = parseFromEnv(_appConfig)
		}

		if !configParsed || !_appConfig.Database.Type.Provider.IsValid() {
			log.Fatalf("error database provider is not valid: %v valid providers are : %+v", _appConfig.Database.Type.Provider, AllProviders)
		}

	})

	return _appConfig
}

func koanfUnmarshall() (configParsed bool) {
	err := k.Unmarshal("", _appConfig)
	if err != nil {
		logger.Log(context.Background(), slog.LevelError, "Error unmarshalling config.yaml", slog.String("error", err.Error()))
	} else {
		configParsed = true
	}
	return configParsed
}

func parseFromEnv(_appConfig *ApplicationConfig) (configParsed bool) {
	var firstAdminSubject string
	var secondAdminSubject string

	err := k.Load(env.Provider("", ".", func(s string) string {
		return s
	}), nil)

	if err != nil {
		logger.Log(context.Background(), slog.LevelError, "Error loading config from environment variables", slog.String("error", err.Error()))
		return false
	}

	if k.Bool("DATABASE.TYPE.SEED") {
		firstAdminSubject = k.String("DATABASE.SEED.ADMINS.0.SUBJECT")
		secondAdminSubject = k.String("DATABASE.SEED.ADMINS.1.SUBJECT")
	}

	logger.Log(context.Background(), slog.LevelInfo, "Loading config from environment variables")

	configParsed = koanfUnmarshall()
	if !configParsed {
		logger.Log(context.Background(), slog.LevelError, "Error unmarshalling config from environment variables")
		return false
	}

	var admins []SeedAdmin

	if firstAdminSubject != "" {
		admins = append(admins, SeedAdmin{Subject: firstAdminSubject})
	}

	if secondAdminSubject != "" {
		admins = append(admins, SeedAdmin{Subject: secondAdminSubject})
	}

	if len(admins) > 0 {
		_appConfig.Database.Seed = &Seed{
			Admins: &admins,
		}
	}

	return true
}
