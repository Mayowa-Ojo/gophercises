package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gophercises/url_shortener/handler"
)

func main() {
	mux := defaultMux()
	var file string

	// parse flags
	flag.StringVar(&file, "yaml", "paths.yaml", "path to yaml file")
	flag.Parse()

	// read file
	yaml, err := readFile(file)
	if err != nil {
		panic(err)
	}
	// mapHandler := handler.MapHandler(pathsToURls, mux)
	yamlHandler, err := handler.YamlHandler(yaml, mux)
	if err != nil {
		panic(err)
	}

	fmt.Printf("[>] reading paths from file: %s", file)
	fmt.Println("[>] server running on part :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", root)

	return mux
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Gophercise 2 - URL shortener")
}

func readFile(path string) ([]byte, error) {
	// validate file path
	ok := strings.HasSuffix(path, ".yaml")
	if !ok {
		return nil, errors.New("Invalid file. File must have a .yaml extension")
	}
	// read yaml file
	yamlFile, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return yamlFile, nil
}
