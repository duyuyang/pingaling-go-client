package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type HealthStatus struct {
	Url     string `json:"url"`
	Updated string `json:"updated"`
	Type    string `json:"type"`
	Status  string `json:"status"`
	Name    string `json:"name"`
}

type HealthStatuses struct {
	Data []HealthStatus
}

func HealthStatusSummary() []HealthStatus {
	url := "http://localhost:4000/api/health/summary"

	pingalingClient := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := pingalingClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var statuses HealthStatuses
	jsonErr := json.Unmarshal(body, &statuses)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return statuses.Data
}
