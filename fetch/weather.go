package fetch

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

type Weather struct {
	Description string
}

type Main struct {
	Temperature float64
	Humidity    int64
}

var apiKey = "4e2d8f44d84be2c264a492cae6101ffb"

func getLocationData(city string) []byte {
	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&appid=%s", city, apiKey)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	return body
}

func getWeatherData(lat, lon float64) []byte {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&lang=zh_cn&units=metric", lat, lon, apiKey)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return body
}

type Location struct {
	Name    string
	Lat     float64
	Lon     float64
	Country string
	State   string
}

func GetLocation(city string) (lat float64, lon float64) {
	content := getLocationData(city)

	doc := gjson.ParseBytes(content).Array()
	for _, i := range doc {
		var location Location
		location = Location{
			Name:    i.Get("name").String(),
			Country: i.Get("country").String(),
			State:   i.Get("state").String(),
			Lon:     i.Get("lon").Float(),
			Lat:     i.Get("lat").Float(),
		}
		lat = location.Lat
		lon = location.Lon
	}
	return
}

func GetWeather(city string) (name, description string, temp string, humidity int64) {
	lat, lon := GetLocation(city)
	content := getWeatherData(lat, lon)

	doc := gjson.ParseBytes(content)

	name = doc.Get("name").String()
	main := doc.Get("main").Array()
	var mains Main
	for _, v := range main {
		mains = Main{
			Temperature: v.Get("temp").Float(),
			Humidity:    v.Get("humidity").Int(),
		}
	}

	temp1 := mains.Temperature
	temp = fmt.Sprintf("%.f°C", temp1)
	humidity = mains.Humidity

	w := doc.Get("weather").Array()
	var weather Weather
	for _, v := range w {
		weather = Weather{
			Description: v.Get("description").String(),
		}
	}

	description = weather.Description
	return
}
