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
	var yamlBytes []byte
	if *yamlFile != "" {
		b, err := os.ReadFile(*yamlFile)
		if err != nil {
			log.Fatalf("Failed to read provided yaml file: %v", err)
		}
		yamlBytes = b
		
	} else {
		yamlBytes = []byte(`
		- path: /urlshort
		url: https://github.com/gophercises/urlshort
		- path: /urlshort-final
		url: https://github.com/gophercises/urlshort/tree/solution
		`)
	}
	yamlHandler, err := urlshort.YAMLHandler(yamlBytes, mapHandler)
		if err != nil {
			panic(err)
		}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
