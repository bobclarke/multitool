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

	// Set up handler for GET on /hello
	r.HandleFunc("/hello", handler).Methods("GET")

	// Set up handler for static content
	//staticFileDirectory := http.Dir("." + assetsDir)
	//staticFileHandler := http.StripPrefix(assetsDir, http.FileServer(staticFileDirectory))
	//r.PathPrefix(assetsDir).Handler(staticFileHandler).Methods("GET")

	staticContent := http.Dir("." + assetsDir)
	staticFileHandler := http.StripPrefix(assetsDir, http.FileServer(staticContent))
	r.PathPrefix(assetsDir).Handler(staticFileHandler).Methods("GET")

	return r
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func staticHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Static")
}
