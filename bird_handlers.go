package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Bird ...
/*
type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}
*/
type Bird struct {
	Species     string
	Description string
}

// Until we put a proper database in place, the birds variable
// is where our list of Birds is held

var birds []Bird

func getBird(w http.ResponseWriter, r *http.Request) {

	// Convert brids into JSON (represented as []byte)
	birdListBytes, err := json.Marshal(birds)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(birdListBytes)

}

func addBird(w http.ResponseWriter, r *http.Request) {
	// Create an instance of Bird
	bird := Bird{}

	// Our POST request will contain a form. We use the ParseForm
	// method of our Request object
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	// Update our Bird instance
	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	// Append our Bird to the Birds array
	birds = append(birds, bird)

	// Some logging for debug
	//fmt.Printf("bird is: %v\n", bird)
	//fmt.Printf("birds is: %v\n", birds)

	// Redirect to index.html
	http.Redirect(w, r, "/assets/", http.StatusFound)

}
