package weather

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

func NewOpenWeatherUtil(apiKey string) *OpenWeatherUtil {
	return &OpenWeatherUtil{apiKey: apiKey}
}

func (owu *OpenWeatherUtil) GetWeatherDetailsFromCoordinates(lat float32, lon float32) WeatherDetails {
	// Call Weather API

	// Extract fields (weather.main, main.temp)

	return WeatherDetails{
		Condition:   "",
		Temperature: 0,
	}
}
