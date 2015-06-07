package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
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

	var (
		baseURL    = flag.String("baseURL", "http://localhost:3000", "Remote base url") // HL
		delay      = flag.Int("delay", 1000, "Server delay in msec")
		deadline   = flag.Int("deadline", 1000, "Deadline in msec")
		concurrent = flag.Int("concurrent", 10, "Num concurrent clients to be started")
	)
	flag.Parse()

	log.Printf("Using baseURL: %s, delay: %d, deadline: %d, concurrent: %d",
		*baseURL, *delay, *deadline, *concurrent)

	// run in parallel with deadline
	numClientsWithinDeadline, averageResponseTime :=
		runClients(*baseURL, *delay, *deadline, *concurrent)

	// report stats
	fmt.Printf("%d of %d clients returned within deadline %d\n",
		numClientsWithinDeadline, *concurrent, *deadline)
	fmt.Printf("average response-time: %d\n", averageResponseTime)
}

func runClients(baseURL string, delay int, deadline int, concurrent int) (int, int) {
	// Start all go-routines and let them report on a buffered result channel
	resultChannel := make(chan int, concurrent)
	for i := 0; i < concurrent; i++ {
		go getitWrapped(baseURL, delay, deadline, i, resultChannel)
	}

	// create slice to collect results
	results := make([]int, 0, concurrent)

	// Waitfor completion of go-routines or deadline
	timeout := time.After(time.Duration(deadline) * time.Millisecond)

SelectLoop:
	// whatever happens first: all completed or timer expire
	for {
		select {
		case result := <-resultChannel:
			results = append(results, result)
			log.Printf("Got finished client: %d", len(results))
			if len(results) >= concurrent {
				log.Printf("Done")
				break SelectLoop
			}
		case <-timeout:
			log.Printf("Timeout")

			break SelectLoop
		}
	}
	// summarize
	sum := 0
	for _, res := range results {
		sum += res
	}

	average := 0
	if len(results) > 0 {
		average = sum / len(results)
	}

	// Note: writers can still write theit response after the deadline and won't block on it
	time.Sleep(time.Duration(2 * time.Second))

	return len(results), average
}

func getitWrapped(baseURL string, delay int, deadline int, loop int, resultChannel chan int) {
	start := time.Now()

	url := fmt.Sprintf("%s/step6?delay=%d&loop=%d", baseURL, delay, loop)
	_, err := getIt(url, deadline)
	if err != nil {
		log.Printf("Client timout fetching with url %s", url)
		//log.Printf("Timout fetching with url %s: %s", url, err.Error())
		// do not report back
		return
	}

	duration := int(time.Since(start) / 1000000)

	log.Printf("Client completed successfully in %d", duration)

	resultChannel <- duration
}

func getIt(url string, timeoutMsec int) (*Response, error) {
	// Note use timeout so go-routines do not keep hanging around

	// Perform HTTP GET
	timeout := time.Duration(timeoutMsec) * time.Millisecond

	client := http.Client{Timeout: timeout}
	// Perform HTTP GET
	httpResponse, err := client.Get(url)
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
