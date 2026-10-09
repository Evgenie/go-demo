package main

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"flag"
	"fmt"
)

func main() {
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")
	flag.Parse()
	fmt.Println(*city)

	geoData, err := geo.GetUserLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(geoData)
	weatherData := weather.GetWeatherData(*geoData, *format)
	fmt.Printf("%s\n", weatherData)
}
