package main

import "testing"

func Test_weatherInfo_precipitationVolume(test *testing.T) {
	type fields struct {
		Rain *precipitationInfo
		Snow *precipitationInfo
	}
	type data struct {
		name                      string
		fields                    fields
		wantedPrecipitationVolume float32
	}

	tests := []data{
		data{
			name: "without precipitation",
			fields: fields{
				Rain: nil,
				Snow: nil,
			},
			wantedPrecipitationVolume: 0,
		},
		data{
			name: "with rain",
			fields: fields{
				Rain: &precipitationInfo{
					Volume: 2.3,
				},
				Snow: nil,
			},
			wantedPrecipitationVolume: 2.3,
		},
		data{
			name: "with snow",
			fields: fields{
				Rain: nil,
				Snow: &precipitationInfo{
					Volume: 2.3,
				},
			},
			wantedPrecipitationVolume: 2.3,
		},
	}
	for _, testData := range tests {
		info := weatherInfo{
			Rain: testData.fields.Rain,
			Snow: testData.fields.Snow,
		}
		receivedPrecipitationVolume := info.precipitationVolume()

		if receivedPrecipitationVolume != testData.wantedPrecipitationVolume {
			test.Logf(
				"failed %q:\n  expected: %+v\n  actual: %+v",
				testData.name,
				testData.wantedPrecipitationVolume,
				receivedPrecipitationVolume,
			)
			test.Fail()
		}
	}
}
