package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
)

var Mac string

func SetMac() {
	netInterface, err := net.InterfaceByName("eth0")
	if err != nil {
		netInterface, err = net.InterfaceByName("eth1")
		if err != nil {
			netInterface, err = net.InterfaceByName("en0")
			if err != nil {
				log.Fatal("无法获取mac地址")
			}
		}
	}

	Mac = fmt.Sprintf("%x", []byte(netInterface.HardwareAddr))
}
