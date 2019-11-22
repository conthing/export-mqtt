package api

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"net/http"
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
