package mq

type MQHandler interface {
	Set(payload []byte) error
	Delete(payload []byte) error
	Get(payload []byte) ([]byte, error)
	Fetch(payload []byte) ([]byte, error)
}
