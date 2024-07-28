package config

import (
	"log"
	"sync"
	"time"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type ApplicationConfig struct {
	GraphQL          GraphqlConfig    `koanf:"graphql"`
	Database         DatabaseConfig   `koanf:"database"`
	Redis            RedisConfig      `koanf:"redis"`
	Services         ServicesConfig   `koanf:"services"`
	CasbinConfig     CasbinConfig     `koanf:"casbin"`
	AuthorizerConfig AuthorizerConfig `koanf:"authorizer"`
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
	Model    string            `koanf:"model"`
	LogLevel string            `koanf:"log_level"`
	Mongo    CasbinMongoConfig `koanf:"mongodb"`
}

type AuthorizerConfig struct {
	ClientID      string             `koanf:"client_id"`
	ClientSecret  string             `koanf:"client_secret"`
	Issuer        string             `koanf:"issuer"`
	AuthorizerURL string             `koanf:"authorizer_url"`
	TokenURL      string             `koanf:"token_url"`
	ResponseMode  string             `koanf:"response_mode"`
	ResponseType  string             `koanf:"response_type"`
	Scopes        []string           `koanf:"scope"`
	Callback      string             `koanf:"redirect_url"`
	ExtraHedars   *map[string]string `koanf:"extra_headers"`
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
	k = koanf.NewWithConf(conf)
)

const (
	yamlPath = "config/config.yaml"
)

func (ConfigurationService) Get() *ApplicationConfig {
	onceConfigure.Do(func() {
		_appConfig = &ApplicationConfig{}

		if err := k.Load(file.Provider(yamlPath), yaml.Parser()); err != nil {
			log.Fatalf("error loading config: %v", err)
		}

		err := k.Unmarshal("", _appConfig)
		if err != nil {
			log.Fatalf("error unmarshal config: %v", err)
		}

		if !_appConfig.Database.Type.Provider.IsValid() {
			log.Fatalf("error database provider is not valid: %v valid providers are : %+v", _appConfig.Database.Type.Provider, AllProviders)
		}

	})

	return _appConfig
}
