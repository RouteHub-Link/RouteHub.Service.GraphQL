package database

import (
	"os"
	"os/signal"
	"syscall"

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
