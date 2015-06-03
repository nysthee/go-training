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
	delay := flag.Int("delay", 1000, "Server delay in msec")
	deadline := flag.Int("deadline", 1000, "Deadline in msec")
	concurrent := flag.Int("concurrent", 10, "Num concurrent clients to be started")
	flag.Parse()

	numClientsWithinDeadline := runClients(*baseURL, *delay, *deadline, *concurrent)

	fmt.Printf("%d of %d clients returned within deadline\n", numClientsWithinDeadline, concurrent)
}

func runClients(baseURL string, delay int, deadline int, concurrent int) int {
	numClientsWithinDeadline := 0
	//
	// TODO add your code here: fetch concurrently
	//

	// This is not concurrent but sequential
	for i := 0; i < concurrent; i++ {
		_, err := getIt(fmt.Sprintf("%s?delay=%d&loop=%d", baseURL, delay, i))
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
