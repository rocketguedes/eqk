package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//"net/url"
)

type terremoto struct {
	Type string `json:"type"`
	Metadata struct {
		Generated int64 `json:"generated"`
		URL string `json:"url"`
		Title string `json:"title"`
		Status int `json:"status"`
		API string `json:"api"`
		Count int `json:"count"`
	} `json:"metadata"`
	Features []struct {
		Type string `json:"type"`
		Properties struct {
			Mag float64 `json:"mag"`
			Place string `json:"place"`
			Time int64 `json:"time"`
			Updated int64 `json:"updated"`
			Tz int `json:"tz"`
			URL string `json:"url"`
			Detail string `json:"detail"`
			Felt int `json:"felt"`
			Cdi float64 `json:"cdi"`
			Mmi float64 `json:"mmi"`
			Alert string `json:"alert"`
			Status string `json:"status"`
			Tsunami int `json:"tsunami"`
			Sig int `json:"sig"`
			Net string `json:"net"`
			Code string `json:"code"`
			Ids string `json:"ids"`
			Sources string `json:"sources"`
			Types string `json:"types"`
			Nst interface{} `json:"nst"`
			Dmin float64 `json:"dmin"`
			Rms float64 `json:"rms"`
			Gap int `json:"gap"`
			MagType string `json:"magType"`
			Type string `json:"type"`
			Title string `json:"title"`
		} `json:"properties"`
		Geometry struct {
			Type string `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
		ID string `json:"id"`
	} `json:"features"`
	Bbox []float64 `json:"bbox"`
}


func main() {
	// JSON URL
	url := fmt.Sprintf("https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/significant_month.geojson")

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
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
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record terremoto

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	fmt.Println("Terremotos acima de 6 graus na escala Richter, nos últimos 30 dias:")
	fmt.Println("------------------------------------------------")
	
	var i int

	for i = 0; i < record.Metadata.Count; i++ {

			j := record.Features[i].Properties.Mag

			if ( j > 6){
				fmt.Println("Epicentro = ", record.Features[i].Properties.Place)
				fmt.Println("Magnitude: ", record.Features[i].Properties.Mag)
				fmt.Println("------------------------------------------------")
				//fmt.Println("Fuso horário: ", record.Features[i].Properties.Tz)
				//fmt.Println("Horário: ", record.Features[i].Properties.Time)
			}
	}
}
