package mq

type MessageBus interface {
	Subscribe()
	SubscribeQueue()
	Publish()
	PublishQueue()
}

var msgBusMap = map[string]MessageBus{}

func NewMessageBus(name string) MessageBus {
	msgBus, ok := msgBusMap[name]
	if !ok {

	}
	return msgBus
}
