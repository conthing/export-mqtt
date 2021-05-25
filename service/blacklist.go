package service

import (
	"export-mqtt/api"
	"export-mqtt/subscriber"

	"github.com/conthing/utils/common"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func SubscribeBlackListService() {
	subscriber.SubscribeBlackList(common.GetSerialNumber(), blackListCallback)
}

func blackListCallback(client MQTT.Client, msg MQTT.Message) {
	data := msg.Payload()
	api.PostBlackList(data)
}
