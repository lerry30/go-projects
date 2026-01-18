package main

import (
	"fmt"
	"time"
	"net/http"
	"encoding/json"
)

type OpenWeather struct {
	id string
	apiKey string
	baseUrl string
	httpClient *http.Client
}

type WeatherData struct {
	Temp float64
	Description string
}

// ----
// Initialize
func NewOpenWeather(id, key string) *OpenWeather {
	return &OpenWeather{
		id: id,
		apiKey: key,
		baseUrl: "https://api.openweathermap.org/data/2.5",
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (ow *OpenWeather) GetCurrent(city string) (WeatherData, error) {
	url := fmt.Sprintf("%s/weather?q=%s&appid=%s&units=matric", ow.baseUrl, city, ow.apiKey)

	resp, err := ow.httpClient.Get(url)
	if err != nil {
		return WeatherData{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return WeatherData{}, fmt.Errorf("openweather error: %s", resp.Status)
	}

	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return WeatherData{}, err
	}

	if len(data.Weather) == 0 {
		return WeatherData{}, fmt.Errorf("No weather data")
	}

	return WeatherData{
		Temp: data.Main.Temp,
		Description: data.Weather[0].Description,
	}, nil
}

/*{
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
}*/