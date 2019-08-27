package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test405(t *testing.T) {
	// Get a Router instance
	r := getRouter()

	// Create a test http server
	mockserver := httptest.NewServer(r)

	// Make sure Posting to /hello gives the correct "method not allowed" response
	resp, err := http.Post(mockserver.URL+"/hello", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status code should be 405, got: %d", resp.StatusCode)
	}

}
