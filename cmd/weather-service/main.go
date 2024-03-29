// This is an example of implementing the Pet Store from the OpenAPI documentation
// found at:
// https://github.com/OAI/OpenAPI-Specification/blob/master/examples/v3.0/petstore.yaml

package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/cdecoux/weather-service/api"
	"github.com/cdecoux/weather-service/internal/weather"
	"github.com/go-chi/chi/v5"
	middleware "github.com/oapi-codegen/nethttp-middleware"
)

const (
	WEATHER_SERVICE_PORT = "8080"
)

var (
	OPEN_WEATHER_API_KEY = os.Getenv("OPEN_WEATHER_API_KEY")
)

func main() {

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}
	swagger.Servers = nil

	// Setup weather utility
	weatherUtil := weather.NewOpenWeatherUtil(OPEN_WEATHER_API_KEY)

	// Create an instance of our handler which satisfies the generated interface
	weatherService := api.NewWeatherService(api.WeatherServiceConfig{
		WeatherUtil: weatherUtil,
		HotTemp:     90,
		ColdTemp:    60,
	})
	weatherServiceHandler := api.NewStrictHandler(weatherService, nil)

	// Setting up basic Chi router with validation middleware
	r := chi.NewRouter()
	r.Use(middleware.OapiRequestValidator(swagger))
	api.HandlerFromMux(weatherServiceHandler, r)

	// Start listening for server
	s := &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("0.0.0.0", WEATHER_SERVICE_PORT),
	}

	log.Fatal(s.ListenAndServe())
}
