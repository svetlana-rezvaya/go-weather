package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		log.Fatal("OpenWeather API Key is missing; " +
			"please, specify the OPENWEATHER_API_KEY environment variable")
	}

	city := flag.String("city", "New York", "city name")
	flag.Parse()

	weatherData := weatherInfo{}
	url := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/weather?appid=%s&q=%s&units=metric",
		apiKey,
		*city,
	)
	if err := loadJSONData(&http.Client{}, url, &weatherData); err != nil {
		log.Fatalf("unable to load the weather data: %s", err)
	}

	fmt.Printf("weather: %+v\n", weatherData)
}
