package main

import "fmt"

func displayAsPlainText(info weatherInfo) {
	fmt.Printf("Location: %s\n", info.location())
	if len(info.Kind) > 0 {
		fmt.Printf("Weather description: %s\n", info.Kind[0].Description)
	}
	fmt.Printf("Temperature: %.2f °C\n", info.Main.Temperature)
	fmt.Printf("Temperature feels like: %.2f °C\n", info.Main.TemperatureFeelsLike)
	fmt.Printf("Wind direction: %d°\n", info.Wind.Direction)
	fmt.Printf("Wind speed: %.2f km/h\n", convertToKmPerH(info.Wind.Speed))
	fmt.Printf(
		"Wind gust speed: %.2f km/h\n",
		convertToKmPerH(info.Wind.GustSpeed),
	)
	fmt.Printf("Visibility: %d m\n", info.Visibility)
	fmt.Printf("Precipitation volume: %.2f mm/h\n", info.precipitationVolume())
	fmt.Printf("Pressure: %d hPa\n", info.Main.Pressure)
	fmt.Printf("Humidity: %d %%\n", info.Main.Humidity)
}

func convertToKmPerH(value float32) float32 {
	return value * 3.6
}
