package main

import (
	"fmt"
	"time"
	"net/http"
	"encoding/json"

	"backend/models"
)

type OpenWeather struct {
	id string
	apiKey string
	baseUrl string
	httpClient *http.Client
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

func (ow *OpenWeather) GetCurrent(city string) (models.WeatherData, error) {
	url := fmt.Sprintf("%s/weather?q=%s&appid=%s&units=matric", ow.baseUrl, city, ow.apiKey)

	resp, err := ow.httpClient.Get(url)
	if err != nil {
		return models.WeatherData{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.WeatherData{}, fmt.Errorf("openweather error: %s", resp.Status)
	}

	var data models.RawWeatherData

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return models.WeatherData{}, err
	}

	if len(data.Weather) == 0 {
		return models.WeatherData{}, fmt.Errorf("No weather data")
	}

	wd := models.NewWeatherData()
	wd.TransformWeatherValues(&data)

	return *wd, nil
}