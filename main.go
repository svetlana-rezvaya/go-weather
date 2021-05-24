package main

import (
	"flag"
	"fmt"
	"log"
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

	fmt.Printf("apiKey: %s\n", apiKey)
	fmt.Printf("city: %s\n", *city)
}
