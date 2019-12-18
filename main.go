package main

import (
	"encoding/json"
	"export-mqtt/client"
	"export-mqtt/dto"
	"export-mqtt/service"
	"net/http"

	log "github.com/sirupsen/logrus"
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
	http.Handle("/api/v1/ping", http.HandlerFunc(PingHandler))
	http.Handle("/api/v1/status", http.HandlerFunc(StatusHandler))
	err := http.ListenAndServe(":52018", nil)
	if err != nil {
		log.Fatal("http服务器启动失败")
	}
}
func PingHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("pong"))
}

func StatusHandler(w http.ResponseWriter, req *http.Request) {
	var resp dto.StatusInfo
	resp.Status = client.Status
	data, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
	}
	w.Write(data)
}
