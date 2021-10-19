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



func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter City name to get Weather information: ")
	scanner.Scan()
	input := scanner.Text()
	BaseURI := `http://api.openweathermap.org/data/2.5/weather?q=` + input + `&units=metric&appid=`
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
