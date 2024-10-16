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

// minimumMagnitude is the threshold for displaying earthquakes
var minimumMagnitude float64

func main() {
	minimumMagnitude = parseMinimumMagnitude()
	totalEarthquakes := listQuakes(minimumMagnitude)
	fmt.Printf("Total number of Earthquakes: %d\n", totalEarthquakes)
}

func parseMinimumMagnitude() float64 {
	if len(os.Args) > 1 {
		if n, err := strconv.ParseFloat(os.Args[1], 64); err == nil {
			return n
		}
		log.Printf("Invalid magnitude provided, using default of 0")
	}
	return 0
}

func listQuakes(minimumMagnitude float64) int {
	earthquakeData, err := fetchEarthquakeData()
	if err != nil {
		log.Fatalf("Failed to fetch earthquake data: %v", err)
	}

	fmt.Println("-------------------------------------------------------------------")
	fmt.Printf("Earthquake(s) above %.1f degrees, in the last 30 days:\n", minimumMagnitude)
	fmt.Println("-------------------------------------------------------------------")

	totalEarthquakes := 0

	for _, feature := range earthquakeData.Features {
		if feature.Properties.Mag > minimumMagnitude {
			printEarthquakeInfo(feature)
			totalEarthquakes++
		}
	}

	return totalEarthquakes
}

func printEarthquakeInfo(feature struct {
	Type       string `json:"type"`
	Properties struct {
		Mag     float64 `json:"mag"`
		Place   string  `json:"place"`
		Time    int64   `json:"time"`
		Updated int64   `json:"updated"`
		Tz      int     `json:"tz"`
	} `json:"properties"`
}) {
	fmt.Println("Epicenter =", feature.Properties.Place)
	fmt.Printf("Magnitude: %.1f\n", feature.Properties.Mag)
	t := time.UnixMilli(feature.Properties.Time)
	fmt.Println("Time:", t.UTC())
	fmt.Println("-------------------------------------------------------------------")
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
