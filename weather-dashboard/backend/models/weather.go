package models

import (
	"fmt"
	"time"

	"backend/utils"
)

type WeatherData struct {
	Name string
	Temp string
	FeelsLike string
	Description string
	Humidity string
	Pressure string
	Icon string
	Country string
	Sunrise time.Time
	Sunset time.Time
	Speed string
	WindDirection string
	Clouds string
	Visibility string
	LastUpdateTime time.Time
}

type RawWeatherData struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Humidity int `json:"humidity"`
		Pressure int `json:"pressure"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
		Icon string `json:"icon"`
	} `json:"weather"`
	Sys struct {
		Country string `json:"country"`
		Sunrise int64 `json:"sunrise"`
		Sunset int64 `json:"sunset"`
	} `json:"sys"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg int `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Visibility int `json:"visibility"`
	DT int64 `json:"dt"`
	Timezone int `json:"timezone"`
}

func NewWeatherData() *WeatherData {
	return &WeatherData{}
}

func (wd *WeatherData) TransformWeatherValues(raw *RawWeatherData) {
	if len(raw.Weather) == 0 {
		return
	}

	const kDiff = 273.15 // C = K - 273.15
	temp := fmt.Sprintf("%.2f°C", raw.Main.Temp - kDiff)
	feelsLike := fmt.Sprintf("%.2f°C", raw.Main.FeelsLike - kDiff)
	humidity := fmt.Sprintf("%d%%", raw.Main.Humidity)
	pressure := fmt.Sprintf("%d%% hPa", raw.Main.Pressure)

	icon := fmt.Sprintf("http://openweathermap.org/img/w/%s", raw.Weather[0].Icon)

	sunrise := utils.ToLocalTime(raw.Sys.Sunrise, raw.Timezone)
	sunset := utils.ToLocalTime(raw.Sys.Sunset, raw.Timezone)
	
	speed := fmt.Sprintf("%.2f m/s", raw.Wind.Speed)
	windDirection := fmt.Sprintf("%d°", raw.Wind.Deg)
	
	clouds := fmt.Sprintf("%d%%", raw.Clouds.All)
	visibility := fmt.Sprintf("%.2f km", float64(raw.Visibility) / 1000.0)
	
	lastUpdateTime := utils.ToLocalTime(raw.DT, raw.Timezone)

	wd.Name = raw.Name
	wd.Temp = temp
	wd.FeelsLike = feelsLike
	wd.Description = raw.Weather[0].Description
	wd.Humidity = humidity
	wd.Pressure = pressure
	wd.Icon = icon
	wd.Country = raw.Sys.Country
	wd.Sunrise = sunrise
	wd.Sunset = sunset
	wd.Speed = speed
	wd.WindDirection = windDirection
	wd.Clouds = clouds
	wd.Visibility = visibility
	wd.LastUpdateTime = lastUpdateTime
}

/*
{
	"coord": {
		"lon": 120.9833,
		"lat": 15.2833
  	},
	"weather": [
		{
			"id": 804,
			"main": "Clouds",
			"description": "overcast clouds",
			"icon": "04n"
		}
	],
	"base": "stations",
	"main": {
		"temp": 296.78,
		"feels_like": 297.52,
		"temp_min": 296.78,
		"temp_max": 296.78,
		"pressure": 1005,
		"humidity": 89,
		"sea_level": 1005,
		"grnd_level": 1005
	},
	"visibility": 10000,
	"wind": {
		"speed": 1.49,
		"deg": 17,
		"gust": 1.62
	},
	"clouds": {
		"all": 100
	},
	"dt": 1768672892,
	"sys": {
		"country": "PH",
		"sunrise": 1768688761,
		"sunset": 1768729572
	},
	"timezone": 28800,
	"id": 1713227,
	"name": "City of Gapan",
	"cod": 200
}
*/