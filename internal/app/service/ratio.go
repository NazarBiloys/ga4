package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Ratios []struct {
	CurrencyCode string  `json:"CurrencyCodeL"`
	Amount       float64 `json:"Amount"`
}

func GetCurrentRatio(client *http.Client) (float64, error) {
	// create a new HTTP POST request with the payload
	req, err := http.NewRequestWithContext(context.Background(), "GET", "https://bank.gov.ua/NBU_Exchange/exchange?json", nil)
	if err != nil {
		return 0.0, err
	}

	// send the HTTP request
	res, err := client.Do(req)
	if err != nil {
		return 0.0, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return 0.0, err
	}

	var ratios Ratios

	err = json.Unmarshal([]byte(body), &ratios)
	if err != nil {
		return 0.0, err
	}

	for _, rate := range ratios {
		if rate.CurrencyCode == "USD" {
			return rate.Amount, nil
		}
	}

	return 0.0, fmt.Errorf("Not found usd ratio")
}
