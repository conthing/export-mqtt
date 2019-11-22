package subscriber

import (
	"export-mqtt/client"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func SubscribeBlackList(mac string, callback mqtt.MessageHandler) {
	topic := "/parklock/" + mac + "/blacklist"
	client.Subscribe(topic, callback)
}
