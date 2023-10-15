// Programa em Go para listar os terremotos acima de 6 graus nos últimos 30 dias.
// Fonte dos dados: https://earthquake.usgs.gov/ - USGS
// Autor: Marcelo Pinheiro - [Twitter](http://twitter.com/mpinheir)
//---------------------------------------------------------------------------------------

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

type Earthquake struct {
	Type     string
	Meta     Metadata `json:"metadata"`
	Features []struct {
		Type       string `json:"type"`
		Properties struct {
			Mag     float64     `json:"mag"`
			Place   string      `json:"place"`
			Time    int64       `json:"time"`
			Updated int64       `json:"updated"`
			Tz      int         `json:"tz"`
			URL     string      `json:"url"`
			Detail  string      `json:"detail"`
			Felt    int         `json:"felt"`
			Cdi     float64     `json:"cdi"`
			Mmi     float64     `json:"mmi"`
			Alert   string      `json:"alert"`
			Status  string      `json:"status"`
			Tsunami int         `json:"tsunami"`
			Sig     int         `json:"sig"`
			Net     string      `json:"net"`
			Code    string      `json:"code"`
			Ids     string      `json:"ids"`
			Sources string      `json:"sources"`
			Types   string      `json:"types"`
			Nst     interface{} `json:"nst"`
			Dmin    float64     `json:"dmin"`
			Rms     float64     `json:"rms"`
			Gap     int         `json:"gap"`
			MagType string      `json:"magType"`
			Type    string      `json:"type"`
			Title   string      `json:"title"`
		} `json:"properties"`
		Geometry struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
		ID string `json:"id"`
	} `json:"features"`
	Bbox []float64 `json:"bbox"`
}

func main() {
	// Fetch earthquake data from the API
	earthquakeData, err := fetchEarthquakeData()
	if err != nil {
		log.Fatal("Failed to fetch earthquake data:", err)
	}

	fmt.Println("-------------------------------------------------------------------")
	fmt.Println(" Terremotos acima de 6 graus na escala Richter, nos últimos 30 dias:")
	fmt.Println("-------------------------------------------------------------------")

	totTerremotos := earthquakeData.Meta.Count

	for item := 0; item < totTerremotos; item++ {

		magnitude := earthquakeData.Features[item].Properties.Mag

		if magnitude > 6 {
			fmt.Println("Epicentro =", earthquakeData.Features[item].Properties.Place)
			fmt.Println("Magnitude:", earthquakeData.Features[item].Properties.Mag)

			t := time.UnixMilli(earthquakeData.Features[item].Properties.Time)
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
