//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=weather-service-api.gen.yml weather-service-api.yml
package api

import "context"

type WeatherService struct {
	config WeatherServiceConfig
}

// Help conform to interface
var _ StrictServerInterface = (*WeatherService)(nil)

type WeatherServiceConfig struct {
	// Anything below specified ColdTemp will be summarized as "COLD"
	ColdTemp int
	// Anything above HotTemp will be summarized as "HOT" (this must be greater than ColdTemp )
	HotTemp int
}

func NewWeatherService() *WeatherService {
	return &WeatherService{}
}

// GetWeather implements StrictServerInterface.
func (*WeatherService) GetWeather(ctx context.Context, request GetWeatherRequestObject) (GetWeatherResponseObject, error) {
	panic("unimplemented")
	// Call Weather API

	// Extract fields (weather.main, main.temp)

	// Summarize temp into MODERATE, COLD, HOT
}
