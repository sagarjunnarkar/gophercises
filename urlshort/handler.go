package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

//MapHandler blabla
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

//YamlHandler blabla
// func YamlHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
// 	pathUrls, err := parseYaml(yamlBytes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	pathToUrls := buildMap(pathUrls)
// 	return MapHandler(pathToUrls, fallback), nil
// }

func parseYaml(data []byte) ([]pathURL, error) {
	var pathUrls []pathURL
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

func buildMap(pathUrls []pathURL) map[string]string {
	pathToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathToUrls[pu.Path] = pu.URL
	}
	return pathToUrls
}

// YamlFileHandler bla
func YamlFileHandler(yamlFileName string, fallback http.Handler) (http.HandlerFunc, error) {
	yamlFile, err := ioutil.ReadFile(yamlFileName)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return nil, err
	}
	pathUrls, err := parseYaml(yamlFile)
	if err != nil {
		return nil, err
	}

	pathToUrls := buildMap(pathUrls)
	return MapHandler(pathToUrls, fallback), nil
}

//JsonFileHandler bla
func JsonFileHandler(jsonFileName string, fallback http.Handler) (http.HandlerFunc, error) {
	jsonFile, err := ioutil.ReadFile(jsonFileName)
	if err != nil {
		log.Printf("jsonFile.Get err   #%v ", err)
		return nil, err
	}
	pathUrls, err := parseJSON(jsonFile)
	if err != nil {
		return nil, err
	}

	pathToUrls := buildMap(pathUrls)
	return MapHandler(pathToUrls, fallback), nil
}

func parseJSON(data []byte) ([]pathURL, error) {
	var pathUrls []pathURL
	err := json.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

type pathURL struct {
	Path string
	URL  string
}
