package service

import (
	"export-mqtt/api"
	"export-mqtt/config"
	"export-mqtt/subscriber"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func SubscribeBlackListService() {
	subscriber.SubscribeBlackList(config.Mac, blackListCallback)
}

func blackListCallback(client MQTT.Client, msg MQTT.Message) {
	data := msg.Payload()
	api.PostBlackList(data)
}
