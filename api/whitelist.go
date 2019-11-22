package api

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var WhiteListURL = "http://localhost/parklock/blacklist"

func PostWhiteList(body []byte) {
	_, PostErr := http.Post(WhiteListURL,
		"application/json",
		bytes.NewBuffer(body))

	if PostErr != nil {
		log.Error("PostWhiteList 失败", PostErr)
	}
}
