package main

import (
	"fmt"
	"log"
)

type response struct {
	Origin string `json:"origin"`
	// TODO: add more attributes
}

func main() {
	hostname := "http://httpbin.org"
	resp, err := getIt(hostname, "/get")
	if err != nil {
		log.Printf("Error fetching from %s: %+v", hostname, err)
		return
	}
	fmt.Printf("%+v", resp)

}

func getIt(hostname string, url string) (*response, error) {
	// TODO: create a http client that fetches json response into the 'response'-struct

	// TODO: Use the http package to perform a http-get

	// ODO: use the encoding package to unserialize json into a struct

	return nil, fmt.Errorf("%s not implemented", "getIt")
}
