// This Go program fetches and displas Earthquake data.
// Data source: https://earthquake.usgs.gov/ - USGS
// Author: Marcelo Pinheiro - [Twitter](http://twitter.com/mpinheir)
//---------------------------------------------------------------------------------------

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// EarthquakeAPIURL is the URL for earthquake data.
const EarthquakeAPIURL = "https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/significant_month.geojson"

// Metadata contains metadata information.
type Metadata struct {
	Generated int64  `json:"generated"`
	URL       string `json:"url"`
	Title     string `json:"title"`
	Status    int    `json:"status"`
	API       string `json:"api"`
	Count     int    `json:"count"`
}

// Earthquake represents earthquake data.
type Earthquake struct {
	Type     string `json:"type"`
	Meta     Metadata
	Features []struct {
		Type       string `json:"type"`
		Properties struct {
			Mag     float64 `json:"mag"`
			Place   string  `json:"place"`
			Time    int64   `json:"time"`
			Updated int64   `json:"updated"`
			Tz      int     `json:"tz"`
		} `json:"properties"`
	} `json:"features"`
}

// Program will display Eartjquates with magnitude > minimumMagnitude
var minimumMagnitude float64

func main() {

	minimumMagnitude = 0

	if len(os.Args[1:]) > 0 {
		if n, err := strconv.ParseFloat(os.Args[1], 64); err == nil {
			minimumMagnitude = n
		}
	}

	// Fetch earthquake data from the API
	earthquakeData, err := fetchEarthquakeData()
	if err != nil {
		log.Fatal("Failed to fetch earthquake data:", err)
	}

	fmt.Println("-------------------------------------------------------------------")
	fmt.Printf(" Earthquake above %f degrees, in the last 30 dias:\n", minimumMagnitude)
	fmt.Println("-------------------------------------------------------------------")

	for _, feature := range earthquakeData.Features {

		magnitude := feature.Properties.Mag

		if magnitude > minimumMagnitude {
			fmt.Println("Epicenter =", feature.Properties.Place)
			fmt.Println("Magnitude:", magnitude)

			t := time.UnixMilli(feature.Properties.Time)
			fmt.Println("Time:", t.UTC())

			fmt.Println("-------------------------------------------------------------------")
		}
	}

}

func fetchEarthquakeData() (Earthquake, error) {
	// Build the request
	req, err := http.NewRequest("GET", EarthquakeAPIURL, nil)
	if err != nil {
		return Earthquake{}, err
	}

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Earthquake{}, err
	}
	defer resp.Body.Close()

	// Decode the JSON response into the Earthquake struct
	var earthquakeData Earthquake
	if err := json.NewDecoder(resp.Body).Decode(&earthquakeData); err != nil {
		return Earthquake{}, err
	}

	return earthquakeData, nil
}
