package service

import (
	"export-mqtt/api"
	"export-mqtt/subscriber"

	"github.com/conthing/utils/common"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func SubscribeWhiteListService() {
	subscriber.SubscribeWhiteList(common.GetSerialNumber(), whiteListCallback)
}

func whiteListCallback(client MQTT.Client, msg MQTT.Message) {
	data := msg.Payload()
	api.PostWhiteList(data)
}
