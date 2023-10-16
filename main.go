package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"time"
)

type NewsRelease struct {
	Title       string `json:"title"`
	RevisedDate string `json:"revisedDate"`
	Abstract    string `json:"abstract"`
}

type Response struct {
	Data []NewsRelease `json:"data"`
}

func main() {
	var apiKey string
	var i int

	flag.StringVar(&apiKey, "apikey", "", "API key for the National Park API")
	flag.IntVar(&i, "i", 5, "Time Interval (5 for 5 seconds)")
	flag.Parse()

	if apiKey == "" {
		fmt.Println("API key is required. Use -apikey flag to specify it.")
		return
	}

	address := "https://developer.nps.gov/api/v1/newsreleases?api_key=" + apiKey

	// Create a timer to fetch data at regular intervals.
	ticker := time.NewTicker(time.Duration(i) * time.Second) // Change the interval as needed.

	for {
		select {
		case <-ticker.C:
			err := fetchData(address)
			if err != nil {
				log.Printf("Error: %v", err)
			}
		}
	}
}

func fetchData(address string) error {
	resp, err := http.Get(address)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned a non-OK status: %d", resp.StatusCode)
	}

	var response Response
	decoder := json.NewDecoder(resp.Body)

	if err := decoder.Decode(&response); err != nil {
		return fmt.Errorf("Error decoding JSON response: %v", err)
	}

	if len(response.Data) > 0 {
		// Log the first data entry (article)
		firstArticle := response.Data[0]
		log.Printf("Title: %s\nRevised Date: %s\nAbstract: %s", firstArticle.Title, firstArticle.RevisedDate, firstArticle.Abstract)
	}

	return nil
}
