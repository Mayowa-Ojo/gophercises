package main

import (
	"fmt"
	"net/http"

	"github.com/gophercises/url_shortener/handler"
)

func main() {
	mux := defaultMux()

	// build the mapHandler with mux as the fallback
	// pathsToURls := map[string]string{
	// 	"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	// 	"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	// }

	yaml := `
- path: /urlshort-godoc
  url: https://github.com/gophercises/urlshort
- path: /yaml-godoc
  url: https://godoc.org/gopkg.in/yaml.v2
`

	// mapHandler := handler.MapHandler(pathsToURls, mux)
	yamlHandler, err := handler.YamlHandler([]byte(yaml), mux)
	if err != nil {
		panic(err)
	}

	fmt.Println("server running on part :8080")
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
