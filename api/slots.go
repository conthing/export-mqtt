package api

import (
	"io/ioutil"
	"net/http"
)

var SlotsURL = "http://localhost/parklock/slots"

func GetSlots() ([]byte, error) {
	resp, err := http.Get(SlotsURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
