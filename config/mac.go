package config

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
)

var Mac string

func SetMac(name string) {
	netInterface, err := net.InterfaceByName(name)
	if err != nil {
		log.Fatal("无法获取mac地址", name)
	}

	Mac = fmt.Sprintf("%x", []byte(netInterface.HardwareAddr))
}
