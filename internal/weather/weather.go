package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type WeatherDetails struct {
	Condition   string
	Temperature int
}

type WeatherUtil interface {
	GetWeatherDetailsFromCoordinates(lat float32, lon float32) WeatherDetails
}

type OpenWeatherUtil struct {
	apiKey string
}

const (
	openWeatherBaseURL = "https://api.openweathermap.org/data/2.5/weather"
)

func NewOpenWeatherUtil(apiKey string) *OpenWeatherUtil {
	return &OpenWeatherUtil{apiKey: apiKey}
}

type openWeatherResponse struct {
	Weather []struct {
		Main string `json:"main"`
	} `json:"weather"`
	Main struct {
		Temp float32 `json:"temp"`
	}
}

// TODO: This should return an error for upstream's discretion
func (owu *OpenWeatherUtil) GetWeatherDetailsFromCoordinates(lat float32, lon float32) WeatherDetails {
	// Setup call for OpenWeather service
	req, err := url.Parse(openWeatherBaseURL)
	if err != nil {
		return WeatherDetails{}
	}
	// Query params
	params := url.Values{}
	params.Add("lon", fmt.Sprintf("%.2f", lon))
	params.Add("lat", fmt.Sprintf("%.2f", lat))
	params.Add("appid", owu.apiKey)
	params.Add("units", "imperial")
	req.RawQuery = params.Encode()

	// Call API
	response, err := http.Get(req.String())
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return WeatherDetails{}
	}
	defer response.Body.Close()

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: Unexpected status code", response.StatusCode)
		return WeatherDetails{}
	}

	// Read the response body
	var openWeatherResponse openWeatherResponse
	err = json.NewDecoder(response.Body).Decode(&openWeatherResponse)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return WeatherDetails{}
	}

	if len(openWeatherResponse.Weather) == 0 {
		fmt.Println("No weather data for provided inputs")
		return WeatherDetails{}
	}

	// Extract fields (weather.main, main.temp)
	return WeatherDetails{
		Condition:   openWeatherResponse.Weather[0].Main,
		Temperature: int(openWeatherResponse.Main.Temp),
	}
}
