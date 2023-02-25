package service

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var confFile = "/app/config.json"

type Config struct {
	MeasurementId string `json:"measurementId"`
	AppId         string `json:"appId"`
	Event         string `json:"event"`
	ApiSecret     string `json:"apiSecret"`
}

func ReadConfig() (*Config, error) {
	// Load the specified config file from the path provided
	jsonFile, err := os.Open(confFile)

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	config := Config{}

	err = json.Unmarshal(byteValue, &config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}
