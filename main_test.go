package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchEarthquakeData(t *testing.T) {
	// Store the original API URL
	originalURL := EarthquakeAPIURL

	// Create a test server to mock the API response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"type": "FeatureCollection",
			"metadata": {
				"generated": 1633455637000,
				"url": "https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/significant_month.geojson",
				"title": "USGS Significant Earthquakes, Past 30 Days",
				"status": 200,
				"api": "1.12.0",
				"count": 2
			},
			"features": [
				{
					"type": "Feature",
					"properties": {
						"mag": 6.5,
						"place": "Location 1",
						"time": 1633455600000
					}
				},
				{
					"type": "Feature",
					"properties": {
						"mag": 7.0,
						"place": "Location 2",
						"time": 1633455700000
					}
				}
			]
		}`))
	}))
	defer server.Close()

	// Override the API URL with the test server's URL
	EarthquakeAPIURL = server.URL

	earthquakeData, err := fetchEarthquakeData()
	if err != nil {
		t.Errorf("fetchEarthquakeData() returned an error: %v", err)
	}

	// Write more test cases for other scenarios as needed
	if len(earthquakeData.Features) != 2 {
		t.Errorf("Expected 2 earthquake features, got %d", len(earthquakeData.Features))
	}

	// Reset the EarthquakeAPIURL to the original value after the test
	EarthquakeAPIURL = originalURL
}
