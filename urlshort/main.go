package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	yamlFilename := flag.String("yaml-file", "data.yaml", "Yaml file name with redirection URLs")
	jsonFilename := flag.String("json-file", "data.json", "json file name with redirection URLs")
	flag.Parse()

	mux := defaultMux()

	pathsToURL := map[string]string{
		"/a": "https://google.com",
		"/b": "https://facebook.com",
	}
	mapHandler := MapHandler(pathsToURL, mux)

	yamlHandler, err := YamlFileHandler(*yamlFilename, mapHandler)
	if err != nil {
		panic(err)
	}

	jsonHandler, err := JsonFileHandler(*jsonFilename, yamlHandler)
	if err != nil {
		panic(err)
	}
	// yamlHandler, err := YamlHandler([]byte(yaml), mapHandler)
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Printf("http serve on 8080...")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
