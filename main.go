package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

//const PAGESIZE = 100
//const STATSSIZE = 6

func getSLLdata(config *Configuration) []Response {
	sessionurl, err := url.Parse(config.ServerRoot)
	q := sessionurl.Query()
	q.Set("key", config.APIKey)
	q.Set("siteid", config.SiteID)
	q.Set("timewindow", "60")
	sessionurl.RawQuery = q.Encode()

	buf := bytes.NewBufferString("")
	req, err := http.NewRequest(http.MethodGet, sessionurl.String(), buf)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var data ResponseOfDepartures
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(string(body))
		panic(err)
	}

	switch config.TravelType {
	case Metros:
		return data.ResponseData.Metros
	case Buses:
		return data.ResponseData.Buses
	case Trains:
		return data.ResponseData.Trains
	case Trams:
		return data.ResponseData.Trams
	case Ships:
		return data.ResponseData.Ships
	}

	log.Fatalf("Unknown travel type %d", config.TravelType)
	return nil
}

func main() {
	config := getConfig()
	ticker := time.NewTicker(time.Second * 60)
	write2influx(config, getSLLdata(config))
	for _ = range ticker.C {
		write2influx(config, getSLLdata(config))
	}
}
