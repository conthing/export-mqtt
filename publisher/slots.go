package publisher

import (
	"encoding/json"
	"export-mqtt/client"
	"export-mqtt/dto"
	log "github.com/sirupsen/logrus"
)

func PublishSlots(topic string, slots []dto.Slot) {
	byteSlots, err := json.Marshal(slots)
	if err != nil {
		log.Error("PublishSlots 序列化失败", err)
		return
	}
	client.Publish(topic, byteSlots)
}
