package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatic(t *testing.T) {
	r := getRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/assets/")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	expectedContentType := "text/html; charset=utf-8"
	actualContentType := resp.Header.Get("Content-Type")

	if actualContentType != expectedContentType {
		t.Errorf("Expected %s got %s ", expectedContentType, actualContentType)
	}

}
