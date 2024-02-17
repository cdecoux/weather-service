package weather

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func newTestServer(jsonResponse string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate a JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonResponse))
	}))
	return server
}

func Test_parseOpenWeatherResponse_Success(t *testing.T) {
	type args struct {
		jsonResponse string
	}
	tests := []struct {
		name string
		args args
		want WeatherDetails
	}{
		{
			name: "Success Parse of OW Response",
			args: args{jsonResponse: `{"coord":{"lon":10.99,"lat":44.34},"weather":[{"id":803,"main":"Clouds","description":"broken clouds","icon":"04n"}],"base":"stations","main":{"temp":276.03,"feels_like":276.03,"temp_min":274.41,"temp_max":279.96,"pressure":1025,"humidity":74,"sea_level":1025,"grnd_level":936},"visibility":10000,"wind":{"speed":1,"deg":230,"gust":0.91},"clouds":{"all":74},"dt":1708149693,"sys":{"type":2,"id":2044440,"country":"IT","sunrise":1708150448,"sunset":1708188380},"timezone":3600,"id":3163858,"name":"Zocca","cod":200}`},
			want: WeatherDetails{
				Condition:   "Clouds",
				Temperature: 276,
			},
		},
		{
			name: "Success Parse of OW Response - Empty Condition",
			args: args{jsonResponse: `{"coord":{"lon":10.99,"lat":44.34},"base":"stations","main":{"temp":276.03,"feels_like":276.03,"temp_min":274.41,"temp_max":279.96,"pressure":1025,"humidity":74,"sea_level":1025,"grnd_level":936},"visibility":10000,"wind":{"speed":1,"deg":230,"gust":0.91},"clouds":{"all":74},"dt":1708149693,"sys":{"type":2,"id":2044440,"country":"IT","sunrise":1708150448,"sunset":1708188380},"timezone":3600,"id":3163858,"name":"Zocca","cod":200}`},
			want: WeatherDetails{
				Condition:   "",
				Temperature: 276,
			},
		},
	}
	for _, tt := range tests {
		// Create a mock HTTP server
		server := newTestServer(tt.args.jsonResponse)
		defer server.Close()
		resp, _ := http.Get(server.URL)

		t.Run(tt.name, func(t *testing.T) {
			if got := parseOpenWeatherResponse(resp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseOpenWeatherResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
