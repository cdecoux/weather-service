//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=weather-service-api.gen.yml weather-service-api.yml
package api

import (
	"context"

	"github.com/cdecoux/weather-service/internal/weather"
)

type weatherService struct {
	config WeatherServiceConfig
}

// Help conform to interface
var _ StrictServerInterface = (*weatherService)(nil)

type WeatherServiceConfig struct {
	WeatherUtil weather.WeatherUtil
	// Anything below specified ColdTemp will be summarized as "COLD"
	ColdTemp int
	// Anything above HotTemp will be summarized as "HOT" (this must be greater than ColdTemp )
	HotTemp int
}

func NewWeatherService(config WeatherServiceConfig) *weatherService {
	// TODO: Validate configs (such as HotTemp > ColdTemp)
	return &weatherService{config: config}
}

// GetWeather implements StrictServerInterface.
func (w *weatherService) GetWeather(ctx context.Context, request GetWeatherRequestObject) (GetWeatherResponseObject, error) {
	// Get Weather Details
	weatherDetails := w.config.WeatherUtil.GetWeatherDetailsFromCoordinates(*request.Params.Lat, *request.Params.Lon)
	// Summary temperature
	temperatureSummary := MODERATE
	if weatherDetails.Temperature <= w.config.ColdTemp {
		temperatureSummary = COLD
	} else if weatherDetails.Temperature >= w.config.HotTemp {
		temperatureSummary = HOT
	}

	// Summarize temp into MODERATE, COLD, HOT
	return GetWeather200JSONResponse{
		TemperatureSummary: temperatureSummary,
		WeatherCondition:   weatherDetails.Condition,
	}, nil
}
