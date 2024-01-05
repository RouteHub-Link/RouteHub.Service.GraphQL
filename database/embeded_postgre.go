package database

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	configuration "github.com/RouteHub-Link/routehub-service-graphql/config"
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
)

var embeddedPostgres *embeddedpostgres.EmbeddedPostgres

func RunEmbeddedPostgres() {
	config = configuration.ConfigurationService{}.Get()
	embeddedPostgres = embeddedpostgres.NewDatabase(
		embeddedpostgres.DefaultConfig().
			Database("postgres").
			Username("postgres").
			Password("1234").
			Port(config.Database.Port),
	)
	if err := embeddedPostgres.Start(); err != nil {
		port := fmt.Sprint(config.Database.Port)
		if err.Error() == "process already listening on port "+port {
			log.Printf("\nEmbedded Postgres already running on port " + port + "\nContinueing old session...")
			time.Sleep(3 * time.Second)
			return
		}
		panic(err)
	}
}

func InterruptEmbedded() {
	sig := make(chan os.Signal, 1)
	signal.Notify(
		sig,
		syscall.SIGTERM,
		syscall.SIGINT,
		os.Interrupt,
	)

	<-sig

	embeddedPostgres.Stop()
	os.Exit(0)
}
