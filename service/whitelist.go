package service

import (
	"export-mqtt/api"
	"export-mqtt/config"
	"export-mqtt/subscriber"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func SubscribeWhiteListService() {
	subscriber.SubscribeWhiteList(config.Mac, whiteListCallback)
}

func whiteListCallback(client MQTT.Client, msg MQTT.Message) {
	data := msg.Payload()
	api.PostWhiteList(data)
}
