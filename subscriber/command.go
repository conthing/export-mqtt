package subscriber

import (
	"export-mqtt/client"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func SubscribeCommand(mac string, addr string, callback mqtt.MessageHandler) {
	topic := "/parklock/" + mac + "/" + addr + "/command"
	client.Subscribe(topic, callback)
}
