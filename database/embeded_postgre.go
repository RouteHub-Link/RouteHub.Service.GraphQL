package database

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
)

var embeddedPostgres *embeddedpostgres.EmbeddedPostgres

func RunEmbeddedPostgres() {
	embeddedPostgres = embeddedpostgres.NewDatabase(
		embeddedpostgres.DefaultConfig().
			Database("postgres").
			Username("postgres").
			Password("1234").
			Port(5432),
	)
	if err := embeddedPostgres.Start(); err != nil {
		if err.Error() == "process already listening on port 5432" {
			log.Printf("\nEmbedded Postgres already running on port 5432\n continueing...\n")
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
