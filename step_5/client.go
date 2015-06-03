package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Origin  string              `json:"origin"`
	Method  string              `json:"method"`
	Url     string              `json:"url"`
	Args    map[string][]string `json:"args"`
	Headers map[string][]string `json:"headers"`
	Body    string              `json:"body"`
}

func main() {
	baseURL := flag.String("baseURL", "http://localhost:3000", "Remote base url") // HL
	concurrent := flag.Int("concurrent", 10, "Num concurrent clients to be started")
	flag.Parse()

	numClientsWithinDeadline := runClients(*baseURL, *concurrent)

	fmt.Printf("%d of %d clients returned a result\n", numClientsWithinDeadline, *concurrent)
}

func runClients(baseURL string, concurrent int) int {
	numClientsWithinDeadline := 0

	//
	// This is not concurrent but sequential
	// TODO add your code here: fetch concurrently
	// and report the result back to the caller
	//
	for i := 0; i < concurrent; i++ {
		_, err := getIt(fmt.Sprintf("%s/step_5?delay=500&loop=%d", baseURL, i))
		if err != nil {
			log.Printf("Error fetching in loop %d:%s", i, err.Error())
		}
	}

	return numClientsWithinDeadline
}

func getIt(url string) (*Response, error) {
	// Perform HTTP GET
	httpResponse, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// Decode json
	dec := json.NewDecoder(httpResponse.Body)
	var resp Response
	err = dec.Decode(&resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
