package api

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestGetSlots(t *testing.T) {
	res, err := GetSlots()
	if err != nil {
		t.Error(err)
	}
	log.Info(res.Data[0].Battery)
}
