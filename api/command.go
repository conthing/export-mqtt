package api

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var CommandURL = "http://localhost/parklock/slots/"

func PostCommand(addr string, body []byte) {
	URL := CommandURL + addr + "/command"
	_, PostErr := http.Post(URL,
		"application/json",
		bytes.NewBuffer(body))

	if PostErr != nil {
		log.Error("Post command 失败", PostErr)
	}

	log.Info("post command:", addr)
}
