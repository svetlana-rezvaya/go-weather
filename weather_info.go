package main

import "fmt"

type systemInfo struct {
	Country string
}

type kindInfo struct {
	ID          int
	Description string
}

type mainInfo struct {
	Temperature          float32 `json:"temp"`
	TemperatureFeelsLike float32 `json:"feels_like"`
	Pressure             int
	Humidity             int
}

type windInfo struct {
	Direction int `json:"deg"`
	Speed     float32
	GustSpeed float32 `json:"gust"`
}

type precipitationInfo struct {
	Volume float32 `json:"1h"`
}

type weatherInfo struct {
	City       string     `json:"name"`
	System     systemInfo `json:"sys"`
	Kind       []kindInfo `json:"weather"`
	Main       mainInfo
	Wind       windInfo
	Visibility int
	Rain       *precipitationInfo
	Snow       *precipitationInfo
}

func (info weatherInfo) location() string {
	return fmt.Sprintf("%s, %s", info.City, info.System.Country)
}

func (info weatherInfo) precipitationVolume() float32 {
	volume := float32(0)
	if info.Rain != nil {
		volume = info.Rain.Volume
	} else if info.Snow != nil {
		volume = info.Snow.Volume
	}

	return volume
}
