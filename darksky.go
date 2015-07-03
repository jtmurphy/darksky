package darksky

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

var PrecipTypes = [...]string{
	"rain",
	"snow",
	"sleet", // freezing rain, ice pellets, wintery mix
	"hail",
}

var Icons = [...]string{
	"clear-day",
	"clear-night",
	"rain",
	"snow",
	"sleet",
	"wind",
	"fog",
	"cloudy",
	"partly-cloudy-day",
	"partly-cloudy-night",
}

type Forecast struct {
	Latitude  float64
	Longitude float64
	Timezone  string
	Offset    int
	Currently DataPoint
	Minutely  DataPoint
	Hourly    DataBlock
	Daily     DataBlock
	Alerts    []Alert
}

type DataBlock struct {
	Summary string
	Icon    string      // See Icons
	Data    []DataPoint // Ordered by time
}

type DataPoint struct {
	ApparentTemperature float32 // "Feels like" temperature in ºF
	CloudCover          float32 // Percentage (0 to 1.0) of cloud cover
	DewPoint            float32 // The dew point in ºF
	Humidity            float32 // Percentage (0 to 1.0) of humidity
	Icon                string  // See Icons
	Ozone               float32 // The columnar density in Dobson units.
	// PrecipIntensity: The average expected intensity of precipitation
	// in inches of liquid water per hour. A very rough semantic guide:
	//      0.000 in./hr. coreesponds to no precipitation,
	//      0.002 in./hr. corresponds to very light precipitation,
	//      0.017 in./hr. corresponds to light precipitation,
	//      0.100 in./hr. corresponds to moderate precipitation, and
	//      0.400 in./hr. corresponds to heavy precipitation.
	PrecipIntensity   float32 //
	PrecipType        string  // See PrecipTypes
	PrecipProbability float32 // Probability (0 to 1.0) of precipitation
	Pressure          float32 // in millibars
	Summary           string  //
	Temperature       float64 // in ºF (not defined in Daily!)
	Time              int64   //
	Visibility        float32 // in miles, capped at 10 miles
	WindBearing       float32 // direction coming from. North is 0º, progressing clockwise
	WindSpeed         float32 // in miles per hour

	// The following fields are *only* available in Daily
	ApparentTemperatureMin     float32 // Min "Feels like" temperature in ºF
	ApparentTemperatureMinTime int64   //
	ApparentTemperatureMax     float32 // Max "Feels like" temperature in ºF
	ApparentTemperatureMaxTime int64   //
	MoonPhase                  float32 // 0.0 is new, 1.0 is full
	PrecipIntensityMax         float64 // See PrecipIntensity
	PrecipIntensityMaxTime     int64   //
	SunriseTime                int64   //
	SunsetTime                 int64   //
	TemperatureMin             float32 // Min temperature in ºF
	TemperatureMinTime         int64   //
	TemperatureMax             float32 // Max temperature in ºF
	TemperatureMaxTime         int64   //

	// The following are *only* defined in Hourly and Daily
	PrecipAccumulation float32 // Snowfall in inches

	// The following are *only* defined in Currently
	NearestStormBearing  float32 // in degrees, North is º0, progress clockwise
	NearestStormDistance float32 // in miles (not very accurate)

}

type Alert struct {
	Title       string
	Expires     int64
	Description string
	Uri         string
}

func Get(key string, lat, long float64) (*Forecast, error) {
	var forecast Forecast
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{},
	}
	client := &http.Client{Transport: tr}
	uri := fmt.Sprintf("https://api.forecast.io/forecast/%s/%f,%f",
		key, lat, long)
	resp, err := client.Get(uri)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&forecast)
	if err != nil {
		return nil, err
	}
	return &forecast, nil
}
