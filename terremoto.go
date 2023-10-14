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

type terremotoData struct {
	Type     string `json:"type"`
	Metadata struct {
		Generated int64  `json:"generated"`
		URL       string `json:"url"`
		Title     string `json:"title"`
		Status    int    `json:"status"`
		API       string `json:"api"`
		Count     int    `json:"count"`
	} `json:"metadata"`
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

	for item := 0; item < terremotos.Metadata.Count; item++ {

		magnitude := terremotos.Features[item].Properties.Mag

		if magnitude > 6 {
			fmt.Println("Epicentro =", terremotos.Features[item].Properties.Place)
			fmt.Println("Magnitude:", terremotos.Features[item].Properties.Mag)
			fmt.Println("-------------------------------------------------------------------")
		}
	}
}
