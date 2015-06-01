package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientSuccess(t *testing.T) {
	origin := "127.0.0.1"

	// Force successfull server response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"origin\":\"%s\"}", origin)
	}))
	defer server.Close()

	// Call subject
	_, err := getIt(server.URL, "/get")
	if err != nil {
		t.Errorf("Expected no error, actual '%+v'", err)
	}

	// TODO verify response returned

}

func TestClientFailure(t *testing.T) {

	// This test needs the internet. TODO nock the server an make it testable locally
	_, err := getIt("httpbin.org", "/status/418")
	if err == nil {
		t.Errorf("Expected error, actual no error")
	}
}
