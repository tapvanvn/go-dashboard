package system

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/tapvanvn/go-dashboard/entity"
	"github.com/tapvanvn/goutil"
)

func LoadConfig(configPath string) {

	file, err := os.Open(configPath)

	if err != nil {
		log.Panic("load config file error.", err)
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)

	if err != nil {

		log.Panic("read config file error.", err)
	}
	bytes = goutil.TripJSONComment(bytes)

	config := &entity.Config{}

	err = json.Unmarshal(bytes, &config)

	if err != nil {

		log.Panic("parse config error.", err)
	}
	Config = config
}
