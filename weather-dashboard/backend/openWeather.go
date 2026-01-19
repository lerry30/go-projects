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

func (ow *OpenWeather) GetCurrent(city string) (models.CurrentWeatherData, error) {
	url := fmt.Sprintf("%s/weather?q=%s&appid=%s&units=metric", ow.baseUrl, city, ow.apiKey)

	resp, err := ow.httpClient.Get(url)
	if err != nil {
		return models.CurrentWeatherData{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.CurrentWeatherData{}, fmt.Errorf("openweather error: %s", resp.Status)
	}

	var data models.RawCurrentWeatherData

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return models.CurrentWeatherData{}, err
	}

	if len(data.Weather) == 0 {
		return models.CurrentWeatherData{}, fmt.Errorf("No weather data.")
	}

	wd := models.NewCurrentWeatherData()
	wd.TransformCurrentWeatherValues(&data)

	return *wd, nil
}

func (ow *OpenWeather) GetForecast(city string) (models.ForecastWeatherData, error) {
	url := fmt.Sprintf("%s/forecast?q=%s&appid=%s&units=metric", ow.baseUrl, city, ow.apiKey)

	resp, err := ow.httpClient.Get(url)
	if err != nil {
		return models.ForecastWeatherData{}, fmt.Errorf("openweather error: %s.", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.ForecastWeatherData{}, fmt.Errorf("failed to get weather forecast.")
	}

	var data models.RawForecastWeatherData

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return models.ForecastWeatherData{}, fmt.Errorf("No weather data.")
	}

	if data.COD != "200" || len(data.List) == 0 {
		return models.ForecastWeatherData{}, fmt.Errorf("No weather data.")
	}

	fw := models.NewForecastWeatherData()
	fw.TransformForecastWeatherValues(&data)

	return *fw, nil
}