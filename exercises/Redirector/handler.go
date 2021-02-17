package main

import (
	"net/http"

	yaml "gopkg.in/yaml.v3"
)

// PathURL is just a structure to represent key value pair
type PathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// MapHandler will handle requests that will be matched from Map
// Map pathToURLs
func MapHandler(pathToURLs map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathToURLs[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YamlHandler will try to convert yaml to map and then serve the requests
func YamlHandler(yamlContent []byte, fallback http.Handler) (http.HandlerFunc, error) {
	data, err := yamlParser(yamlContent)
	if err != nil {
		return nil, err
	}
	return MapHandler(arrayToMapConvertor(data), fallback), nil
}

func yamlParser(data []byte) ([]PathURL, error) {
	var pathArray []PathURL
	err := yaml.Unmarshal(data, &pathArray)
	if err != nil {
		return nil, err
	}
	return pathArray, nil
}

func arrayToMapConvertor(pathURLs []PathURL) map[string]string {
	urlMap := make(map[string]string)
	for _, ele := range pathURLs {
		urlMap[ele.Path] = ele.URL
	}
	return urlMap
}
