package main

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/schachmat/wego/iface"
)

func Test_displayAsASCIIArt(test *testing.T) {
	type args struct {
		info weatherInfo
	}
	type data struct {
		name         string
		args         args
		wantedOutput string
	}

	tests := []data{
		data{
			name: "success",
			args: args{
				info: weatherInfo{
					City: "City",
					System: systemInfo{
						Country: "Country",
					},
					Kind: []kindInfo{},
					Main: mainInfo{
						Temperature:          1.2,
						TemperatureFeelsLike: 2.3,
						Pressure:             34,
						Humidity:             45,
					},
					Wind: windInfo{
						Direction: 56,
						Speed:     10,
						GustSpeed: 100,
					},
					Visibility: 67,
					Rain: &precipitationInfo{
						Volume: 1000,
					},
					Snow: nil,
				},
			},
			wantedOutput: "Weather for City, Country\n" +
				"\n" +
				"     .-.       \n" +
				"      __)      1 (2) °C       \n" +
				"     (         ↙ 36 – 360 km/h\n" +
				"      `-᾿      67 m           \n" +
				"       •       1.0 m/h        \n" +
				"               34 hPa\n" +
				"               45 %\n",
		},
	}
	for _, testData := range tests {
		receivedOutput, err := captureOutput(test, func() {
			displayAsASCIIArt(testData.args.info)
		})
		if err != nil {
			test.Logf("failed %q: %s", testData.name, err)
			test.FailNow()
		}

		receivedOutput = regexp.MustCompile(`\x1b\[.*?m`).
			ReplaceAllString(receivedOutput, "")

		if !reflect.DeepEqual(receivedOutput, testData.wantedOutput) {
			test.Logf(
				"failed %q:\n  expected: %+v\n  actual: %+v",
				testData.name,
				testData.wantedOutput,
				receivedOutput,
			)
			test.Fail()
		}
	}
}

func Test_convertToWegoData(test *testing.T) {
	type args struct {
		info weatherInfo
	}
	type data struct {
		name           string
		args           args
		wantedWegoData iface.Data
	}

	tests := []data{
		data{
			name: "with the weather kind",
			args: args{
				info: weatherInfo{
					City: "City",
					System: systemInfo{
						Country: "Country",
					},
					Kind: []kindInfo{
						kindInfo{
							ID:          200,
							Description: "Description",
						},
					},
					Main: mainInfo{
						Temperature:          1.2,
						TemperatureFeelsLike: 2.3,
						Pressure:             34,
						Humidity:             45,
					},
					Wind: windInfo{
						Direction: 56,
						Speed:     10,
						GustSpeed: 100,
					},
					Visibility: 67,
					Rain: &precipitationInfo{
						Volume: 1000,
					},
					Snow: nil,
				},
			},
			wantedWegoData: iface.Data{
				Location: "City, Country",
				Current: iface.Cond{
					TempC:         pointerToFloat32(1.2),
					FeelsLikeC:    pointerToFloat32(2.3),
					WinddirDegree: pointerToInt(56),
					WindspeedKmph: pointerToFloat32(36),
					WindGustKmph:  pointerToFloat32(360),
					VisibleDistM:  pointerToFloat32(67),
					PrecipM:       pointerToFloat32(1),
					Code:          iface.CodeThunderyShowers,
					Desc:          "Description",
				},
			},
		},
		data{
			name: "without the weather kind",
			args: args{
				info: weatherInfo{
					City: "City",
					System: systemInfo{
						Country: "Country",
					},
					Kind: []kindInfo{},
					Main: mainInfo{
						Temperature:          1.2,
						TemperatureFeelsLike: 2.3,
						Pressure:             34,
						Humidity:             45,
					},
					Wind: windInfo{
						Direction: 56,
						Speed:     10,
						GustSpeed: 100,
					},
					Visibility: 67,
					Rain: &precipitationInfo{
						Volume: 1000,
					},
					Snow: nil,
				},
			},
			wantedWegoData: iface.Data{
				Location: "City, Country",
				Current: iface.Cond{
					TempC:         pointerToFloat32(1.2),
					FeelsLikeC:    pointerToFloat32(2.3),
					WinddirDegree: pointerToInt(56),
					WindspeedKmph: pointerToFloat32(36),
					WindGustKmph:  pointerToFloat32(360),
					VisibleDistM:  pointerToFloat32(67),
					PrecipM:       pointerToFloat32(1),
					Code:          iface.CodeUnknown,
					Desc:          "",
				},
			},
		},
	}
	for _, testData := range tests {
		receivedWegoData := convertToWegoData(testData.args.info)

		if !reflect.DeepEqual(receivedWegoData, testData.wantedWegoData) {
			test.Logf(
				"failed %q:\n  expected: %+v\n  actual: %+v",
				testData.name,
				testData.wantedWegoData,
				receivedWegoData,
			)
			test.Fail()
		}
	}
}

func pointerToInt(value int) *int {
	return &value
}

func pointerToFloat32(value float32) *float32 {
	return &value
}
