package subscriber

import (
	"export-mqtt/client"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func SubscribeGetSlots(mac string, callback mqtt.MessageHandler) {
	topic := "/parklock/" + mac + "/getSlots"
	client.Subscribe(topic, callback)
}
