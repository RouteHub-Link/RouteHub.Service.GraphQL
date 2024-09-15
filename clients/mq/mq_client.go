package mq

import (
	"os"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTClient struct {
	tcpAddr string
	Timeout time.Duration
	Client  *mqtt.Client
}

func NewMQTTClient(tcpAddr string) (*MQTTClient, error) {
	rc := MQTTClient{
		tcpAddr: tcpAddr,
		Timeout: 30 * time.Second,
	}

	hostName := os.Getenv("HOSTNAME")

	options := mqtt.NewClientOptions().AddBroker(rc.tcpAddr)
	options.SetClientID(strings.Join([]string{"ROUTEHUB.SERVICE.GRAPHQL", hostName}, "/"))

	client := mqtt.NewClient(options)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	rc.Client = &client

	return &rc, nil
}
