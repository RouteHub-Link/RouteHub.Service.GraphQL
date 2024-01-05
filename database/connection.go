package database

import (
	"log"
	"os"
	"time"

	configuration "github.com/RouteHub-Link/routehub-service-graphql/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB     *gorm.DB
	config *configuration.ApplicationConfig
)

func Init() {
	config = configuration.ConfigurationService{}.Get()

	connectionSelector()

	if config.Database.Type.Migrate {
		Migrate(DB)
	}

	if config.Database.Type.Seed {
		Seed()
	}
}

func connectionSelector() {
	_logger := getLogger()
	gormConfig := &gorm.Config{
		Logger: _logger,
	}

	switch config.Database.Type.Provider {
	case configuration.Postgres:
		setupPostgres(gormConfig, config)
	case configuration.Embed:
		setupEmbeded(gormConfig, config)
	case configuration.Mysql:
		setupMysql(gormConfig, config)
	default:
		log.Fatal("Database provider not implemented")
	}

}

func getLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Minute,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			Colorful:                  true,
		},
	)
}

func setupPostgres(gormConfig *gorm.Config, config *configuration.ApplicationConfig) {
	dsn := "postgresql://" + config.Database.User + ":" + config.Database.Password + "@" + config.Database.Host + ":" + config.Database.PortAsString + "/" + config.Database.Database + "?application_name=" + config.Database.ApplicationName
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}

func setupEmbeded(gormConfig *gorm.Config, config *configuration.ApplicationConfig) {
	RunEmbeddedPostgres()
	go InterruptEmbedded()
	db, err := gorm.Open(
		postgres.Open("host=127.0.0.1 user=postgres password=1234 dbname=postgres port="+config.Database.PortAsString+" sslmode=disable TimeZone=UTC"),
		gormConfig,
	)
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}

func setupMysql(gormConfig *gorm.Config, config *configuration.ApplicationConfig) {
	dsn := config.Database.User + ":" + config.Database.Password + "@tcp(" + config.Database.Host + ":" + config.Database.PortAsString + ")/" + config.Database.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}
