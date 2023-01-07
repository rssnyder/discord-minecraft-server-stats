package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	mcapiURL = "https://mcapi.us/server/status?ip=%s"
)

type mcapiStats struct {
	Status   string `json:"status"`
	Online   bool   `json:"online"`
	Motd     string `json:"motd"`
	MotdJSON struct {
		Extra []struct {
			Color string `json:"color"`
			Text  string `json:"text"`
		} `json:"extra"`
		Text string `json:"text"`
	} `json:"motd_json"`
	Favicon string      `json:"favicon"`
	Error   interface{} `json:"error"`
	Players struct {
		Max    int `json:"max"`
		Now    int `json:"now"`
		Sample []struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"sample"`
	} `json:"players"`
	Server struct {
		Name     string `json:"name"`
		Protocol int    `json:"protocol"`
	} `json:"server"`
	LastUpdated string `json:"last_updated"`
	Duration    string `json:"duration"`
}

func GetMcapiStats(serverDomain string) (result mcapiStats, err error) {

	reqURL := fmt.Sprintf(mcapiURL, serverDomain)
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return
	}

	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}

	results, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(results, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
