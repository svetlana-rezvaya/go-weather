package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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

	url := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/weather?appid=%s&q=%s&units=metric",
		apiKey,
		*city,
	)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("unable to create the request: %s", err)
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("unable to send the request: %s", err)
	}
	defer response.Body.Close()

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("unable to read the response: %s", err)
	}

	fmt.Printf("response: %s\n", responseBytes)
}
