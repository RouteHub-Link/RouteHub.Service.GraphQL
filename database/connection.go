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
	log.Println("Database Init Called")

	connectionSelector()
}

func Migration() {
	log.Println("Database Migrate Called")
	if config.Database.Type.Migrate {
		log.Println("Database Migrate Started")
		migrate(DB)
	}
}

func Seed() {
	log.Println("Database Seed Called")
	if config.Database.Type.Seed {
		log.Println("Database Seed Started")
		if config.Database.Seed == nil {
			log.Fatal("Seed data not found in config")
		}

		seed()
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
	log.Println("Database Provider: Postgres")

	dsn := config.Database.GetPostgreDSN()
	log.Println("Database Connection String: ", dsn)

	_dialector := postgres.Open(dsn)

	db, err := gorm.Open(_dialector, gormConfig)
	if err != nil {
		log.Fatal(err)
	}
	DB = db

	log.Println("Database Connection Established")
}

func setupEmbeded(gormConfig *gorm.Config, config *configuration.ApplicationConfig) {
	log.Println("Database Provider: Embeded")
	RunEmbeddedPostgres()
	log.Println("Database Running in Embeded Mode")

	go InterruptEmbedded()
	db, err := gorm.Open(
		postgres.Open("host=127.0.0.1 user=postgres password=1234 dbname=postgres port="+config.Database.PortAsString+" sslmode=disable TimeZone=UTC"),
		gormConfig,
	)
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	log.Println("Database Connection Established")
}
