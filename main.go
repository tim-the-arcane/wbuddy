package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

func GetJson(url string, target interface{}) error {
	res, err := client.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	return json.NewDecoder(res.Body).Decode(&target)
}

type WeatherResponse struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		IsDay     int     `json:"is_day"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Hour []struct {
			TimeEpoch int     `json:"time_epoch"`
			TempC     float64 `json:"temp_c"`
		} `json:"hour"`
	} `json:"forecast"`
}

var client *http.Client

func main() {
	client = &http.Client{Timeout: 10 * time.Second}
	apiKey := "00c5a8f6443d42958af41628233107"
	q := "Heilbronn"

	if len(os.Args) > 1 {
		q = ""

		for _, arg := range os.Args[1:] {
			q += " " + arg
		}
	}

	var posts WeatherResponse

	err := GetJson("https://api.weatherapi.com/v1/forecast.json?key="+apiKey+"&q="+url.QueryEscape(q)+"&aqi=no", &posts)

	if err != nil {
		panic(err)
	}

	fmt.Printf("It's currently %s with %.0f°C in %s, %s\n", posts.Current.Condition.Text, posts.Current.TempC, posts.Location.Name, posts.Location.Country)

	for _, hourlyForecast := range posts.Forecast.Hour {
		fmt.Printf("%0f\n", hourlyForecast.TempC)
	}
}
