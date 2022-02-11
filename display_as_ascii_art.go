package main

import (
	"fmt"
	"strings"

	_ "github.com/schachmat/wego/frontends"
	"github.com/schachmat/wego/iface"
)

func displayAsASCIIArt(info weatherInfo) {
	wegoData := convertToWegoData(info)
	iface.AllFrontends["ascii-art-table"].Render(wegoData, iface.UnitsMetric)

	indent := strings.Repeat(" ", 15)
	fmt.Printf(indent+"%d hPa\n", info.Main.Pressure)
	fmt.Printf(indent+"%d %%\n", info.Main.Humidity)
}

func convertToWegoData(info weatherInfo) iface.Data {
	windSpeed := convertToKmPerH(info.Wind.Speed)
	windGustSpeed := convertToKmPerH(info.Wind.GustSpeed)
	visibility := float32(info.Visibility)
	precipitationVolumeInM := info.precipitationVolume() / 1000
	wegoData := iface.Data{
		Location: info.location(),
		Current: iface.Cond{
			TempC:         &info.Main.Temperature,
			FeelsLikeC:    &info.Main.TemperatureFeelsLike,
			WinddirDegree: &info.Wind.Direction,
			WindspeedKmph: &windSpeed,
			WindGustKmph:  &windGustSpeed,
			VisibleDistM:  &visibility,
			PrecipM:       &precipitationVolumeInM,
		},
	}
	if len(info.Kind) > 0 {
		wegoData.Current.Code = convertToWeatherCode(info.Kind[0].ID)
		wegoData.Current.Desc = info.Kind[0].Description
	} else {
		wegoData.Current.Code = iface.CodeUnknown
	}

	return wegoData
}

func convertToWeatherCode(weatherID int) iface.WeatherCode {
	weatherCode := iface.CodeUnknown
	switch weatherID {
	// thunderstorm
	case 200, 201, 210, 211, 230, 231:
		weatherCode = iface.CodeThunderyShowers
	case 202, 212, 221, 232:
		weatherCode = iface.CodeThunderyHeavyRain

	// drizzle
	case 300, 301, 310, 311, 313, 321:
		weatherCode = iface.CodeLightRain
	case 302, 312, 314:
		weatherCode = iface.CodeHeavyRain

	// rain
	case 500, 501, 520, 521:
		weatherCode = iface.CodeLightShowers
	case 502, 503, 504, 522, 531:
		weatherCode = iface.CodeHeavyShowers

	// sleet
	case 511, 611, 615, 616:
		weatherCode = iface.CodeLightSleet
	case 612, 613:
		weatherCode = iface.CodeLightSleetShowers

	// snow
	case 600, 601:
		weatherCode = iface.CodeLightSnow
	case 602:
		weatherCode = iface.CodeHeavySnow
	case 620, 621:
		weatherCode = iface.CodeLightSnowShowers
	case 622:
		weatherCode = iface.CodeHeavySnowShowers

	// fog
	case 701, 721, 741:
		weatherCode = iface.CodeFog

	// clear
	case 800:
		weatherCode = iface.CodeSunny

	// clouds
	case 801:
		weatherCode = iface.CodePartlyCloudy
	case 802:
		weatherCode = iface.CodeCloudy
	case 803, 804:
		weatherCode = iface.CodeVeryCloudy
	}

	return weatherCode
}
