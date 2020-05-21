package service

import (
	"export-mqtt/api"
	"export-mqtt/config"
	"export-mqtt/dto"
	"export-mqtt/publisher"
	"export-mqtt/subscriber"
	"strconv"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"

	"sort"
	"time"
)

var StoredSlots = make([]dto.Slot, 0)

func SubscribeGetSlots() {
	subscriber.SubscribeGetSlots(config.Mac, func(client MQTT.Client, message MQTT.Message) {
		LoadSlots()
		publishSlots(StoredSlots)
	})
}

func LoadSlots() {
	err := storeSlots()
	for err != nil {
		time.Sleep(3 * time.Second)
		err = storeSlots()
	}
	sort.Sort(dto.ByAddr(StoredSlots))
	subscribeAllCommand(StoredSlots)
}

func subscribeAllCommand(slots []dto.Slot) {
	for _, slot := range slots {
		addr := strconv.Itoa(int(slot.Addr))
		callback := generateCommandCallback(addr)
		subscriber.SubscribeCommand(config.Mac, addr, callback)
	}
}

func generateCommandCallback(addr string) func(MQTT.Client, MQTT.Message) {
	return func(client MQTT.Client, msg MQTT.Message) {
		data := msg.Payload()
		api.PostCommand(addr, data)
	}
}

func DiffSlots() {
	for {
		time.Sleep(time.Second * 5)
		newSlots := make([]dto.Slot, 0)
		storeSlots()
		sort.Sort(dto.ByAddr(newSlots))
		for i, slot := range newSlots {
			if slot.Addr != StoredSlots[i].Addr ||
				slot.Name != StoredSlots[i].Name ||
				slot.IP != StoredSlots[i].IP ||
				slot.Position != StoredSlots[i].Position ||
				slot.State != StoredSlots[i].State ||
				slot.Battery != StoredSlots[i].Battery {
				// 只要有一个不同，就会发布
				publishSlots(newSlots)
			}
		}

		if len(StoredSlots) != len(newSlots) {
			subscribeAllCommand(newSlots)
		}
		StoredSlots = newSlots

	}

}

func publishSlots(slots []dto.Slot) {
	topic := "/parklock/" + config.Mac + "/slots"
	publisher.PublishSlots(topic, slots)
}

func storeSlots() error {
	res, err := api.GetSlots()
	if err != nil {
		log.Error("获取 Slots 失败", err)
		return err
	}

	StoredSlots = res.Data
	return nil

}
