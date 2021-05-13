package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Weather struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`

	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`

	Base string `json:"base"`

	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`

	Visibility int `json:"visibility"`

	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`

	Clouds struct {
		All int `json:"deg"`
	} `json:"clouds"`

	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		Id      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`

	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Code     int    `json:"code"`
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter City name to get Weather information: ")
	scanner.Scan()
	input := scanner.Text()
	BaseURI := `http://api.openweathermap.org/data/2.5/weather?q=` + input + `&units=metric&appid=0ad3305bb5a406f138f141381b488ff3`
	res, err := http.Get(BaseURI)

	if err != nil {
		fmt.Printf("Error during the request to URI and its: %q\n", err)
	}
	defer res.Body.Close()
	jsonDataFromHttp, _ := ioutil.ReadAll(res.Body)

	var myWeather Weather

	marsharlErr := json.Unmarshal([]byte(jsonDataFromHttp), &myWeather)

	if marsharlErr != nil {
		fmt.Print("There is error with convert data ", marsharlErr)
	}
	roundedTemp := math.Round(myWeather.Main.Temp)
	main := myWeather.Weather[0].Main
	description := myWeather.Weather[0].Description
	humidity := myWeather.Main.Humidity
	windSpeed := myWeather.Wind.Speed
	cloudsPercent := myWeather.Clouds.All
	pressure := myWeather.Main.Pressure
	lastUpdated := getDate()
	fmt.Printf("\n\nCurrent Temp is %v and its: %v\nWeather description: %v\nLast updated: %q\n", roundedTemp, main, description, lastUpdated)

	fmt.Printf("Humidity: %v %%\tWind Speed: %v m/s\tClouds: %v %%\tPressure: %v hPa\n", humidity, windSpeed, cloudsPercent, pressure)
}

func getDate() string {
	d := time.Now()
	months := []string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December"}

	days := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saterday",
		"Sunday"}

	day := days[int(d.Weekday())-1]
	date := d.Day()
	month := months[d.Month()-1]
	year := d.Year()

	res := string(day) + ", " + strconv.Itoa(date) + " " + month + ", " + strconv.Itoa(year)
	return res
}
