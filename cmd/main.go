package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Ed-cred/urlshort"
)

func main() {
	yamlFile := flag.String("yaml", "", "YAML file to read URLs from formatted as -path:... url:...")
	jsonFile := flag.String("json", "", "JSON file to read URLs from")
	flag.Parse()
	mux := defaultMux()
	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	if *yamlFile != "" {
		yamlBytes, err := os.ReadFile(*yamlFile)
		if err != nil {
			log.Fatalf("Failed to read provided yaml file: %v", err)
		}
		yamlHandler, err := urlshort.YAMLHandler(yamlBytes, mapHandler)
		if err != nil {
			panic(err)
		}
		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", yamlHandler)

	} else if *jsonFile != "" {

		jsonBytes, err := os.ReadFile(*jsonFile)
		if err != nil {
			log.Fatalf("Failed to read provided yaml file: %v", err)
		}

		jsonHandler, err := urlshort.JSONHandler(jsonBytes, mapHandler)
		if err != nil {
			panic(err)
		}
		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", jsonHandler)
	} else {

		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", mapHandler)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
