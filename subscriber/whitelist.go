package subscriber

import (
	"export-mqtt/client"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func SubscribeWhiteList(mac string, callback mqtt.MessageHandler) {
	topic := "/parklock/" + mac + "/whitelist"
	client.Subscribe(topic, callback)
}
