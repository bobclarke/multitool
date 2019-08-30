package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var assetsDir = "/assets/"

func main() {
	r := getRouter()
	http.ListenAndServe("127.0.0.1:8000", r)
}

func getRouter() *mux.Router {
	r := mux.NewRouter()
	// Hello

	// Set up handler for GET on /hello
	r.HandleFunc("/hello", helloHandler).Methods("GET")

	// Set up GET and POST handlers for /bird
	r.HandleFunc("/bird", getBird).Methods("GET")
	r.HandleFunc("/bird", addBird).Methods("POST")

	// Static content handler (i.e for stuff in /assets/)
	staticContent := http.Dir("." + assetsDir)
	staticFileHandler := http.StripPrefix(assetsDir, http.FileServer(staticContent))
	r.PathPrefix(assetsDir).Handler(staticFileHandler).Methods("GET")

	return r
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func staticHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Static")
}
