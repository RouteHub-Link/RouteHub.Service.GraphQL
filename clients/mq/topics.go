package mq

import "strings"

type MQEvent string

var (
	MQE_LINK_SET     MQEvent = "link/set"
	MQE_LINK_DEL     MQEvent = "link/del"
	MQE_LINK_GET     MQEvent = "link/get"
	MQE_LINK_ALL     MQEvent = "link/all"
	MQE_PLATFORM_SET MQEvent = "platform/set"
	MQE_PLATFORM_GET MQEvent = "platform/get"
)

func (e MQEvent) AsTopic() string {
	return strings.Join([]string{"topic", string(e)}, "/")
}

func (e MQEvent) IsValid() bool {
	switch e {
	case MQE_LINK_SET, MQE_LINK_DEL, MQE_LINK_GET, MQE_LINK_ALL, MQE_PLATFORM_SET, MQE_PLATFORM_GET:
		return true
	}
	return false
}
