package api

import (
	"bytes"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var BlackListURL = "http://localhost/parklock/blacklist"

func PostBlackList(body []byte) {
	_, PostErr := http.Post(BlackListURL,
		"application/json",
		bytes.NewBuffer(body))

	if PostErr != nil {
		log.Error("PostBlackList 失败", PostErr)
	}

}
