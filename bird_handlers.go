package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Bird ...
type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

var birds []Bird

func getBird(w http.ResponseWriter, r *http.Request) {
	birdListBytes, err := json.Marshal(birds)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(birdListBytes)
}

func addBird(w http.ResponseWriter, r *http.Request) {
	// To be implemented

}
