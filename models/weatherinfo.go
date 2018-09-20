package models

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"fmt"
)


type PrimaryData struct {
	Date    string        `json:"date"`
	Message string        `json:"message"`
	Status  int           `json:"status"`
	City    string        `json:"city"`
	Count   int           `json:"count"`
	Data    SecondaryData `json:"data"`
}
type SecondaryData struct {
	Humidity    string    `json:"shidu"`
	Pm25        float64       `json:"pm25"`
	Pm10        float64       `json:"pm10"`
	Quality     string    `json:"quality"`
	Temperature string    `json:"wendu"`
	CatchCold   string    `json:"ganmao"`
	Yesterday   Yesterday `json:"yesterday"`
	Forecast   []*SomeDay  `json:"forecast"`
}
type Yesterday struct {
	SomeDay
}

type SomeDay struct {
	Date          string `json:"date"`
	Sunrise       string `json:"sunrise"`
	High          string `json:"high"`
	Low           string `json:"low"`
	Sunset        string `json:"sunset"`
	Aqi           float64    `json:"aqi"`
	WindDirection string `json:"fx"`
	WindLevel     string `json:"fl"`
	WeatherType   string `json:"type"`
	Notice        string `json:"notice"`
}

type Today struct {
	Humidity string `json:"shidu"`
	Pm25 float64 `json:"pm_25"`
	Pm10 float64 `json:"pm_10"`
	Quality string `json:"quality"`
	Temperature string `json:"wendu"`
	CatchCold string `json:"ganmao"`
	Someday SomeDay `json:"someday"`
}

//获取全部数据:PrimaryData
func GetTotalJSON(url string) (*PrimaryData, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return &PrimaryData{}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("request result failure")
	}

	var primaryData PrimaryData

	readAll, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return &PrimaryData{}, nil
	}

	err = json.Unmarshal(readAll, &primaryData)
	if err != nil {
		log.Fatal(err)
		return &PrimaryData{}, nil
	}
	fmt.Println(primaryData)
	return &primaryData, nil
}

