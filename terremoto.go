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
)

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
	Meta     Metadata
	Features []struct {
		Type       string
		Properties struct {
			Mag     float64
			Place   string
			Time    int64
			Updated int64
			Tz      int
			URL     string
			Detail  string
			Felt    int
			Cdi     float64
			Mmi     float64
			Alert   string
			Status  string
			Tsunami int
			Sig     int
			Net     string
			Code    string
			Ids     string
			Sources string
			Types   string
			Nst     interface{}
			Dmin    float64
			Rms     float64
			Gap     int
			MagType string
			Type    string
			Title   string
		}
		Geometry struct {
			Type        string
			Coordinates []float64
		}
		ID string
	}
	Bbox []float64
}

func main() {
	// API endpoint
	url := fmt.Sprintf("https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/significant_month.geojson")

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
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
	fmt.Println("Terremotos acima de 6 graus na escala Richter, nos últimos 30 dias:")
	fmt.Println("-------------------------------------------------------------------")

	for item := 0; item < terremotos.Meta.Count; item++ {

		magnitude := terremotos.Features[item].Properties.Mag

		if magnitude > 6 {
			fmt.Println("Epicentro =", terremotos.Features[item].Properties.Place)
			fmt.Println("Magnitude:", terremotos.Features[item].Properties.Mag)
			fmt.Println("-------------------------------------------------------------------")
		}
	}
}
