package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGetBird(t *testing.T) {

	// Set up some test data - two ways of doing this
	/*
		testBird := Bird{}
		testBird.Species = "Sparrow"
		testBird.Description = "A small harmless bird"
		birds = append(birds, testBird)
	*/

	testBird := Bird{
		Species:     "Sparrow",
		Description: "A small harmless bird",
	}
	birds = append(birds, testBird)

	// Get a router instance
	r := getRouter()

	// Set up a mock http server and pass it our router
	mockServer := httptest.NewServer(r)

	// Call /birds URL from mockServer
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

	// Check the response body make make sure it contains our test data
	expectedRespBody, _ := json.Marshal(testBird) // Convert from type Bird to JSON
	actualRespBody, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	x := actualRespBody[:len(actualRespBody)-1]
	trimmedRespBody := x[1:]

	if string(expectedRespBody) != string(trimmedRespBody) {
		t.Errorf("Expected response of: %s, got %s", expectedRespBody, trimmedRespBody)
	}
}

func TestAddBird(t *testing.T) {

	// Get a router instance
	r := getRouter()

	// Set up a mock server
	mockServer := httptest.NewServer(r)

	// Create some test data to post
	form := newAddBirdForm()
	requestBody := form.Encode()

	// Carry out a POST
	resp, err := http.Post(mockServer.URL+"/bird", "application/x-www-form-urlencoded", bytes.NewBufferString(requestBody))

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Received status code %d, expected %d", resp.StatusCode, http.StatusOK)
	}

	a, _ := http.Get(mockServer.URL + "/bird")
	actualResp, _ := ioutil.ReadAll(a.Body)
	expectedResp := `[{"Species":"Sparrow","Description":"A small harmless bird"},{"Species":"Eagle","Description":"A large bird pf prey"}]`

	//fmt.Printf("actualResp is: %s\n", string(actualResp))
	//fmt.Printf("expectedResp is: %s\n", expectedResp)

	if string(actualResp) != expectedResp {
		t.Errorf("Expected response of: %s, got %s", expectedResp, actualResp)
	}
}

func newAddBirdForm() *url.Values {
	form := url.Values{}
	form.Set("species", "Eagle")
	form.Set("description", "A large bird pf prey")
	return &form
}
