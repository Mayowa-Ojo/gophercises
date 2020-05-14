package handler

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler - returns an http.HandlerFunc that will
// attempt to map any paths to their corresponding urls.
// If the path is not provided in the map, the fallback
// http.Handler will be called instead.
func MapHandler(pathToURLs map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		// if we can match a path...redirect to it
		if dest, ok := pathToURLs[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
		}
		// else call fallback
		fallback.ServeHTTP(w, r)
	}
}

type pathToURL struct {
	URL  string `yaml:"url"`
	Path string `yaml:"path"`
}

// YamlHandler - parses the provided YAML and returns an
// http.HandlerFunc that will match any path to their
// corresponding URL. If the path is not provided in
// the YAML, the fallback is called.
func YamlHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// parse the yaml
	pathURLs, err := parseYaml(yml)
	if err != nil {
		return nil, err
	}
	// convert yaml array into map
	pathToURLs := buildMap(pathURLs)
	// return a map handler
	return MapHandler(pathToURLs, fallback), nil
}

func parseYaml(data []byte) ([]pathToURL, error) {
	var pathURLs []pathToURL

	err := yaml.Unmarshal(data, &pathURLs)
	if err != nil {
		return nil, err
	}

	return pathURLs, nil
}

func buildMap(pathURLs []pathToURL) map[string]string {
	m := make(map[string]string)

	for _, p := range pathURLs {
		m[p.Path] = p.URL
	}

	return m
}
