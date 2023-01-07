package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	mcsrvstatURL = "https://api.mcsrvstat.us/2/%s"
)

type mcsrvstatStats struct {
	Online bool   `json:"online"`
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	Debug  struct {
		Ping          bool `json:"ping"`
		Query         bool `json:"query"`
		Srv           bool `json:"srv"`
		Querymismatch bool `json:"querymismatch"`
		Ipinsrv       bool `json:"ipinsrv"`
		Cnameinsrv    bool `json:"cnameinsrv"`
		Animatedmotd  bool `json:"animatedmotd"`
		Cachetime     int  `json:"cachetime"`
		Cacheexpire   int  `json:"cacheexpire"`
		Apiversion    int  `json:"apiversion"`
	} `json:"debug"`
	Motd struct {
		Raw   []string `json:"raw"`
		Clean []string `json:"clean"`
		HTML  []string `json:"html"`
	} `json:"motd"`
	Players struct {
		Online int      `json:"online"`
		Max    int      `json:"max"`
		List   []string `json:"list"`
		UUID   struct {
			Spirit55555 string `json:"Spirit55555"`
			Sarsum33    string `json:"sarsum33"`
		} `json:"uuid"`
	} `json:"players"`
	Version  string `json:"version"`
	Protocol int    `json:"protocol"`
	Hostname string `json:"hostname"`
	Icon     string `json:"icon"`
	Software string `json:"software"`
	Map      string `json:"map"`
	Gamemode string `json:"gamemode"`
	Serverid string `json:"serverid"`
	Plugins  struct {
		Names []string `json:"names"`
		Raw   []string `json:"raw"`
	} `json:"plugins"`
	Mods struct {
		Names []string `json:"names"`
		Raw   []string `json:"raw"`
	} `json:"mods"`
	Info struct {
		Raw   []string `json:"raw"`
		Clean []string `json:"clean"`
		HTML  []string `json:"html"`
	} `json:"info"`
}

func GetMcsrvstatStats(serverDomain string) (result mcsrvstatStats, err error) {

	reqURL := fmt.Sprintf(mcsrvstatURL, serverDomain)
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
