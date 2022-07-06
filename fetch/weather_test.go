package fetch

import (
	"fmt"
	"testing"
)

func TestGetWeather(t *testing.T) {
	fmt.Println(GetWeather("beijing"))
	fmt.Println(GetWeather("tianjin"))
	fmt.Println(GetWeather("xian"))
}
func TestGetLocation(t *testing.T) {
	if lat, lon := GetLocation("beijing"); lat != 39.906217 && lon != 116.3912757 {
		t.Error("获得经纬度失败")
	}
}
