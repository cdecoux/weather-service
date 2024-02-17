package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type WeatherDetails struct {
	Condition   string
	Temperature int
}

type WeatherUtil interface {
	GetWeatherDetailsFromCoordinates(lat float32, lon float32) (WeatherDetails, error)
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

func parseOpenWeatherResponse(response *http.Response) WeatherDetails {
	// Read the response body
	var openWeatherResponse openWeatherResponse
	if err := json.NewDecoder(response.Body).Decode(&openWeatherResponse); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return WeatherDetails{}
	}

	// Extract fields (weather.main, main.temp)
	details := WeatherDetails{
		Temperature: int(openWeatherResponse.Main.Temp),
	}

	if len(openWeatherResponse.Weather) == 0 {
		fmt.Println("No weather data for provided inputs")
		return details
	} else {
		details.Condition = openWeatherResponse.Weather[0].Main
	}

	return details
}

func (owu *OpenWeatherUtil) GetWeatherDetailsFromCoordinates(lat float32, lon float32) (WeatherDetails, error) {
	// Setup call for OpenWeather service
	req, err := url.Parse(openWeatherBaseURL)
	if err != nil {
		return WeatherDetails{}, err
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
		return WeatherDetails{}, err
	}
	defer response.Body.Close()

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		var serverErr error
		if b, err := io.ReadAll(response.Body); err != nil {
			serverErr = errors.New("unknown error from OpenWeather backend")
		} else {
			serverErr = fmt.Errorf("error from OpenWeather backend: %v", string(b))
		}
		fmt.Println("Unexpected response from server", serverErr)
		return WeatherDetails{}, serverErr
	}

	return parseOpenWeatherResponse(response), nil
}
