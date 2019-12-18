package client

import (
	"export-mqtt/config"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

var client MQTT.Client

// MQTTURL MQTTURL
var MQTTURL = "tcp://mqtt.conthing.com:1883"

// Status MQTT连接状态
var Status string

// Connect 连接
func Connect() {
	opts := MQTT.NewClientOptions().AddBroker(MQTTURL)
	opts.SetClientID(config.Mac)

	client = MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("client连接失败", token.Error())
	}
	Status = "connected"
	go CheckIsConnected()
}

// CheckIsConnected 检查是否连接
func CheckIsConnected() {
	for {
		time.Sleep(time.Second * 60)
		isConnected := client.IsConnected()
		if isConnected {
			Status = "connected"
		} else {
			Status = "disconnected"
			// 重新连接
			Connect()
		}
	}
}

func Publish(topic string, payload interface{}) {
	client.Publish(topic, 0, false, payload)
	log.Infof("topic:%s 发布成功", topic)
}

func Subscribe(topic string, callback MQTT.MessageHandler) {
	if token := client.Subscribe(topic, 0, callback); token.Wait() && token.Error() != nil {
		log.Fatal("订阅失败", token.Error())
	}
	log.Infof("topic:%s 订阅成功", topic)
}
