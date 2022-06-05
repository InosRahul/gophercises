package main

import (
	"fmt"
	"net/http"
	"urlShort/urlshort"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string {
		"/urlshort-godoc" : "https://godoc.org/githu.com/gophercises/urlshort",

	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yaml := `
 - path: /urlshort
   url: https://github.com/gophercises/urlshort
 - path: /urlshort-final
   url: https://github.com/gophercises/urlshort/tree/solution	
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting server on 8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello World")
}