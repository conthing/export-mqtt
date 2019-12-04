package service

import (
	"bytes"
	"export-mqtt/api"
	"export-mqtt/client"
	"export-mqtt/config"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const configFile = "config.yaml"

// Config 配置模型
type Config struct {
	Host    string `yaml:"host"`
	MQTTURL string `yaml:"mqtturl"`
	NetName string `yaml:"net"`
}

var conf = Config{
	Host:    "localhost",
	MQTTURL: "tcp://mqtt.conthing.com:1883",
	NetName: "eth1",
}

func ConfigService() {

	if !exists(configFile) {
		createConfigFile()
	}
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal("读取配置文件失败", err)
	}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatal("配置文件序列化失败", err)
	}

	//todo
	config.SetMac(conf.NetName)

	api.BlackListURL = "http://" + conf.Host + "/parklock/blacklist"
	api.WhiteListURL = "http://" + conf.Host + "/parklock/whitelist"
	api.SlotsURL = "http://" + conf.Host + "/parklock/slots"
	api.CommandURL = "http://" + conf.Host + "/parklock/slots/"

	client.MQTTURL = conf.MQTTURL
}

// exists 判断所给路径文件/文件夹是否存在
func exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			log.Println(err)
			return true
		}
		return false
	}
	return true
}

func createConfigFile() {
	buf := new(bytes.Buffer)
	err := yaml.NewEncoder(buf).Encode(conf)
	if err != nil {
		log.Fatal("配置文件编码失败", err)
	}

	f, err := os.Create(configFile)
	if err != nil {
		log.Fatal("配置文件创建失败", err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Error("配置文件关闭失败", err)
		}

	}()

	_, err = f.Write(buf.Bytes())
	if err != nil {
		log.Fatal("配置文件写入失败", err)
	}
}
