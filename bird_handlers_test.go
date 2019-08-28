package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBird(t *testing.T) {

	// Get a router instance
	r := getRouter()

	// Get a Mock Server
	mockServer := httptest.NewServer(r)

	// Get the /birds URL from mockServer
	resp, err := http.Get(mockServer.URL + "/bird")
	if err != nil {
		t.Fatalf("Error is: %s", err)
	}

	// Check the status code of the response
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Check the Content-Type header
	expected := "text/plain; charset=utf-8"
	actual := resp.Header.Get("Content-Type")
	if expected != actual {
		t.Errorf("Problem with Content-Type header, expected %s, got %s", expected, actual)
	}
}

func TestAddBird(t *testing.T) {

	// Get a router instance
	r := getRouter()

	// Set up a mock server
	mockServer := httptest.NewServer(r)

	// Create some test data to post
	requestBody := []byte("This is some test data")

	// Carry out a POST
	resp, err := http.Post(mockServer.URL+"/bird", "application/text", bytes.NewBuffer(requestBody))

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Received status code %d, expected %d", resp.StatusCode, http.StatusOK)
	}

}
