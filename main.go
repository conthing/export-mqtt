package main

import (
	"export-mqtt/client"
	"export-mqtt/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
}

func main() {
	service.ConfigService()
	client.Connect()
	service.LoadSlots()
	service.SubscribeBlackListService()
	service.SubscribeWhiteListService()
	service.SubscribeGetSlots()
	startHttpServer()

}

func startHttpServer() {
	http.Handle("/ping", http.HandlerFunc(PingHandler))
	err := http.ListenAndServe(":52018", nil)
	if err != nil {
		log.Fatal("http服务器启动失败")
	}
}
func PingHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("pong"))
}
