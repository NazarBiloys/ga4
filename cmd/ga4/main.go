package main

import (
	"time"
	"net/http"
	"github.com/sirupsen/logrus"

	"ratios-pusher/internal/app/service"
)

func main() {
	// create a new http client with timeout
	client := &http.Client{Timeout: 10 * time.Second}

	config, err := service.ReadConfig()

	if err != nil {
		logrus.Errorf("Error when read file: %v", err)
	}

	// get the current UAH/USD ratio (replace with your own implementation)
	ratio, err := service.GetCurrentRatio(client)

	if err != nil {
		logrus.Errorf("Failed to get ratio: %v", err.Error())
	}

	service.Push(client, config, ratio)
}
