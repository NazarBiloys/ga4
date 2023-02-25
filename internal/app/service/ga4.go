package service

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
)

var (
	baseUri = "https://www.google-analytics.com/mp/collect"
)

func Push(client *http.Client, config *Config, ratio float64) {
	// generate a unique client ID
	clientID := uuid.New().String()

	url := fmt.Sprintf("%s?api_secret=%s&firebase_app_id=%s&measurement_id=%s", baseUri, config.ApiSecret, config.AppId, config.MeasurementId)

	// create the GA4 event payload
	payload := []byte(fmt.Sprintf(`{
    "client_id": "%s",
    "events": [
      {
        "name": "%s",
        "params": {
          "ratio": %.4f
        }
      }
    ]
  }`, clientID, config.Event, ratio))

	// create a new HTTP POST request with the payload
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		logrus.Errorf("Failed to make request: %v", err)
	}
	req.Header.Set("Content-Length", "0")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "My App Name/1.0")

	// send the HTTP request
	res, err := client.Do(req)
	if err != nil {
		logrus.Errorf("failed to send request: %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		logrus.Errorf("Error %v", err)
	}

	logrus.Info(string(body))

	defer res.Body.Close()

	// check the response status code
	if res.StatusCode != http.StatusOK {
		logrus.Errorf("Unexpected response status code: %d\n", res.StatusCode)
	}

	logrus.Println("Event sent successfully!")
}
