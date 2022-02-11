# go-weather

[![Go Report Card](https://goreportcard.com/badge/github.com/svetlana-rezvaya/go-weather)](https://goreportcard.com/report/github.com/svetlana-rezvaya/go-weather)
[![Build Status](https://app.travis-ci.com/svetlana-rezvaya/go-weather.svg?branch=master)](https://app.travis-ci.com/svetlana-rezvaya/go-weather)
[![codecov](https://codecov.io/gh/svetlana-rezvaya/go-weather/branch/master/graph/badge.svg)](https://codecov.io/gh/svetlana-rezvaya/go-weather)

The utility for retrieving current weather data from the [OpenWeather](https://openweathermap.org/) service.

## Installation

```
$ go get github.com/svetlana-rezvaya/go-weather
```

## Usage

```
$ go-weather -h | -help | --help
$ go-weather [options]
```

Options:

- `-h`, `-help`, `--help` &mdash; show the help message and exit;
- `-city STRING` &mdash; city name (default: `New York`);
- `-ascii-art` &mdash; display as [ASCII art](https://en.wikipedia.org/wiki/ASCII_art) (based on the [github.com/schachmat/wego](https://github.com/schachmat/wego) package).

## Output Example

Displaying as plain text:

```
Location: New York, US
Weather description: clear sky
Temperature: -5.87 °C
Temperature feels like: -5.87 °C
Wind direction: 297°
Wind speed: 1.62 km/h
Wind gust speed: 4.82 km/h
Visibility: 10000 m
Precipitation volume: 0.00 mm/h
Pressure: 1028 hPa
Humidity: 61 %
```

Displaying as [ASCII art](https://en.wikipedia.org/wiki/ASCII_art) (based on the [github.com/schachmat/wego](https://github.com/schachmat/wego) package):

```
Weather for New York, US

     \   /     clear sky
      .-.      -5 (-5) °C
   ‒ (   ) ‒   ↘ 1 – 4 km/h
      `-᾿      10 km
     /   \     0.0 mm/h
               1028 hPa
               61 %
```

## License

The MIT License (MIT)

Copyright &copy; 2022 svetlana-rezvaya
