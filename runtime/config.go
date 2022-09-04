package runtime

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tapvanvn/go-dashboard/entity"
	"github.com/tapvanvn/goutil"
)

func ReadConfig(configPath string) (*entity.Config, error) {

	file, err := os.Open(configPath)

	if err != nil {
		fmt.Println("load config file error.", err.Error())
		return nil, err
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)

	if err != nil {

		fmt.Println("read config file error.", err.Error())
		return nil, err
	}
	bytes = goutil.TripJSONComment(bytes)

	config := &entity.Config{}

	err = json.Unmarshal(bytes, &config)

	if err != nil {

		fmt.Println("parse config error.", err.Error())
		return nil, err
	}
	return config, nil
}
