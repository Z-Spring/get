package registry

import (
	"fmt"
	"github.com/Z-Spring/get/fetch"

	"github.com/spf13/cobra"
	"time"
)

func NewWeatherCommand() *cobra.Command {
	weatherCmd := &cobra.Command{
		Use:   "weather",
		Short: "you can get weather infos use this command.",
		Args:  cobra.ExactArgs(1),
		Run:   runWeather,
	}
	return weatherCmd
}

func runWeather(cmd *cobra.Command, args []string) {
	go Spinner(100 * time.Millisecond)

	city := args[0]
	n, d, t, h := fetch.GetWeather(city)
	var weather string
	if n == "" && h == 0 {
		weather = ""
		fmt.Printf("\r%s", weather)
		return
	}
	weather = fmt.Sprintf("\r%s %s %s  湿度：%d", n, d, t, h)
	fmt.Println(weather)
}
