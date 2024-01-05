package config

import (
	"log"
	"sync"
	"time"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type DatabaseConfig struct {
	Host            string             `koanf:"host"`
	Port            uint32             `koanf:"port"`
	PortAsString    string             `koanf:"port"`
	User            string             `koanf:"user"`
	Password        string             `koanf:"password"`
	Database        string             `koanf:"database"`
	ApplicationName string             `koanf:"application_name"`
	Type            DatabaseTypeConfig `koanf:"type"`
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

type DatabaseTypeConfig struct {
	Migrate  bool     `koanf:"migrate"`
	Seed     bool     `koanf:"seed"`
	Provider Provider `koanf:"provider"`
}

type ApplicationConfig struct {
	Database DatabaseConfig `koanf:"database"`
	GraphQL  GraphqlConfig  `koanf:"graphql"`
}

type ConfigurationService struct {
}

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
