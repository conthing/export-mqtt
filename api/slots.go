package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"export-mqtt/dto"
	"io/ioutil"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

var SlotsURL = "http://localhost:52032/api/v1/slots"
var AuthURL = "http://localhost:52000/api/v1/login"
var token string

// GetSlots 获取车位
func GetSlots() (*dto.SlotResp, error) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", SlotsURL, nil)
	if err != nil {
		log.Error("NewRequest:", err)
		return nil, err
	}
	request.Header["Authorization"] = []string{token}
	resp, err := client.Do(request)
	if err != nil {
		log.Error("DoRequest:", err)
		return nil, err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		client := &http.Client{}
		loginParam := dto.LoginParam{
			UserName: "admin",
			Password: "admin",
		}
		data, err := json.Marshal(loginParam)
		if err != nil {
			return nil, err
		}
		request, err := http.NewRequest("POST", AuthURL, bytes.NewReader(data))
		if err != nil {
			log.Error("RE-NewRequest:", err)
			return nil, err
		}

		resp, err := client.Do(request)
		if err != nil {
			log.Error("RE-DoRequest:", err)
			return nil, err
		}

		byteArray, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			log.Error("RE-ReadAll:", err)
			return nil, err
		}

		var tokenResp dto.TokenResp
		err = json.Unmarshal(byteArray, &tokenResp)
		if err != nil {
			log.Error("Unmarshal:", err)
			return nil, err
		}

		token = tokenResp.Token
		return GetSlots()
	}

	if resp.StatusCode != http.StatusOK {
		log.Error("code:", resp.StatusCode)
		return nil, errors.New(strconv.FormatInt(int64(resp.StatusCode), 10))
	}

	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Error("Body:", err)
		return nil, err
	}

	log.Info("res:", string(data))

	var slotResp dto.SlotResp
	err = json.Unmarshal(data, &slotResp)
	if err != nil {
		log.Error("Unmarshal:", err)
		return nil, err
	}

	return &slotResp, nil
}
