package fetch

import (
	"context"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/z-spring/get/myredis"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Weather struct {
	Description string
}

type Main struct {
	Temperature float64
	Humidity    int64
}

var (
	client = myredis.NewRedis()
	ctx    = context.Background()
)

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

type Location struct {
	Name    string
	Lat     float64
	Lon     float64
	Country string
	State   string
}

// GetLocation get city's lat and lon
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
	addLocationToRedis(city, lat, lon)
	return
}

func addLocationToRedis(city string, lat float64, lon float64) {
	client.LPush(ctx, city, lat, lon)
}

func getLocationFromRedis(city string) (float64, float64) {
	data := client.LRange(ctx, city, 0, 1)
	result, err := data.Result()
	if err != nil {
		fmt.Printf("\r%s", err)
	}
	if len(result) != 0 {
		lat, _ := strconv.ParseFloat(result[0], 64)
		lon, _ := strconv.ParseFloat(result[1], 64)
		return lat, lon
	}
	return 0, 0
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
		fmt.Println("\rfetch weather timeout, please try again later.")
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return body
}

func GetWeather(city string) (name, description string, temp string, humidity int64) {
	var lat, lon float64
	lon, lat = getLocationFromRedis(city)
	if lat == 0 && lon == 0 {
		lat, lon = GetLocation(city)
	}
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
