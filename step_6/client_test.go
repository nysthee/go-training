package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetItSuccess(t *testing.T) {
	// Create a fake server and let it return a 200 OK with json response body
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// configurable sleep
		sleepIfNeeded(r.RequestURI)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `
{
   "origin": "145.53.70.8",
   "method":"GET",
   "url":"/getIt",
   "args": {
      "delay": ["1000"],
	  "arg1": ["1"],
      "arg2": ["two"]
   },
   "headers": {
      "Accept": ["application/json"],
      "Accept-Encoding": ["gzip"]
    }
}`)

	}))
	defer server.Close()

	// Perform the action against the fake server
	actualResponse, err := getIt(fmt.Sprintf("%s/getIt?delay=1000&arg1=1&arg2=two", server.URL))
	log.Printf("%+v\n", actualResponse)

	// Verify the response
	assert.NoError(t, err)
	assert.Equal(t, "GET", actualResponse.Method)
	assert.Equal(t, "/getIt", actualResponse.Url)
	assert.Equal(t, "1000", actualResponse.Args["delay"][0])
	assert.Equal(t, "1", actualResponse.Args["arg1"][0])
	assert.Equal(t, "two", actualResponse.Args["arg2"][0])
	assert.Equal(t, "application/json", actualResponse.Headers["Accept"][0])
	assert.Equal(t, "gzip", actualResponse.Headers["Accept-Encoding"][0])
}

func sleepIfNeeded(requestUri string) {
	u, _ := url.Parse(requestUri)
	params, _ := url.ParseQuery(u.RawQuery)
	delayString := params.Get("delay")
	if delayString != "" {
		delay, err := strconv.Atoi(delayString)
		if err == nil {
			log.Printf("Sleeping for %d msec", delay)
			time.Sleep(time.Duration(delay) * time.Millisecond)
		}
	}
}
