package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ApiGet(path string, targetStruct interface{}) {
	url := "http://localhost:4000/api" + path

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

	jsonErr := json.Unmarshal(body, &targetStruct)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
}
