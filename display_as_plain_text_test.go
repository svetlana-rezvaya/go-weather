package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func Test_displayAsPlainText(test *testing.T) {
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
			wantedOutput: "Location: City, Country\n" +
				"Weather description: Description\n" +
				"Temperature: 1.20 °C\n" +
				"Temperature feels like: 2.30 °C\n" +
				"Wind direction: 56°\n" +
				"Wind speed: 36.00 km/h\n" +
				"Wind gust speed: 360.00 km/h\n" +
				"Visibility: 67 m\n" +
				"Precipitation volume: 1000.00 mm/h\n" +
				"Pressure: 34 hPa\n" +
				"Humidity: 45 %\n",
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
			wantedOutput: "Location: City, Country\n" +
				"Temperature: 1.20 °C\n" +
				"Temperature feels like: 2.30 °C\n" +
				"Wind direction: 56°\n" +
				"Wind speed: 36.00 km/h\n" +
				"Wind gust speed: 360.00 km/h\n" +
				"Visibility: 67 m\n" +
				"Precipitation volume: 1000.00 mm/h\n" +
				"Pressure: 34 hPa\n" +
				"Humidity: 45 %\n",
		},
	}
	for _, testData := range tests {
		receivedOutput, err := captureOutput(test, func() {
			displayAsPlainText(testData.args.info)
		})
		if err != nil {
			test.Logf("failed %q: %s", testData.name, err)
			test.FailNow()
		}

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

func captureOutput(test *testing.T, handler func()) (output string, err error) {
	temporaryFile, err := ioutil.TempFile("", "test.*")
	if err != nil {
		return "", fmt.Errorf("unable to create the temporary output file: %w", err)
	}
	defer os.Remove(temporaryFile.Name())
	defer temporaryFile.Close()

	previousStdout := os.Stdout
	os.Stdout = temporaryFile
	defer func() {
		os.Stdout = previousStdout
	}()

	handler()

	outputBytes, err := ioutil.ReadFile(temporaryFile.Name())
	if err != nil {
		return "", fmt.Errorf("unable to read the temporary output file: %w", err)
	}

	return string(outputBytes), nil
}
