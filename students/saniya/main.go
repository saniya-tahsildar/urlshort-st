package main

import (
	"fmt"
	"flag"
	"os"
	"net/http"
    "saniya/urlFile"
)

func main() {
	// yaml file flag
	yamlFilename := flag.String("yaml-file", "", "Yaml file name")
    jsonFilename := flag.String("json-file", "", "Json file name")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	defaultYaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlBytes := []byte(defaultYaml)

	// Read Yaml file if provided as flag
	if *yamlFilename != "" {
		fileBytes, err := os.ReadFile(*yamlFilename)
		if err != nil {
			fmt.Println("Failed to read YAML file: %v\n", err)
			os.Exit(1)
		}
		yamlBytes = fileBytes
	} else if *jsonFilename != "" {
	fileBytes, err := os.ReadFile(*jsonFilename)
	if err != nil {
    			fmt.Println("Failed to read YAML file: %v\n", err)
    			os.Exit(1)
    		}
        yamlBytes = fileBytes
	}

	yamlHandler, err := urlshort.YAMLHandler(yamlBytes, mapHandler)
	if err != nil {
		fmt.Println("Failed to parse YAML: %v\n", err)
		os.Exit(1)
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
