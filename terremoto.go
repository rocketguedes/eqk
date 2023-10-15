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

type Metadata struct {
	Generated int64
	URL       string
	Title     string
	Status    int
	API       string
	Count     int
}
type terremotoData struct {
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
	// Build the request
	req, err := http.NewRequest("GET", EarthquakeAPIURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Fatal: ", err)
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var terremotos terremotoData

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&terremotos); err != nil {
		log.Println(err)
	}

	fmt.Println("-------------------------------------------------------------------")
	fmt.Println(" Terremotos acima de 6 graus na escala Richter, nos últimos 30 dias:")
	fmt.Println("-------------------------------------------------------------------")

	totTerremotos := terremotos.Meta.Count

	for item := 0; item < totTerremotos; item++ {

		magnitude := terremotos.Features[item].Properties.Mag

		if magnitude > 6 {
			fmt.Println("Epicentro =", terremotos.Features[item].Properties.Place)
			fmt.Println("Magnitude:", terremotos.Features[item].Properties.Mag)

			t := time.UnixMilli(terremotos.Features[item].Properties.Time)
			fmt.Println("Time:", t.UTC())

			fmt.Println("-------------------------------------------------------------------")
		}
	}
}
