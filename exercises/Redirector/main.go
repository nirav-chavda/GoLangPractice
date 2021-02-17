package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	mux := defaultMux()

	pathsToURL := map[string]string{
		"/gle": "https://google.com/search?q=hello+there",
		"/ytb": "https://youtube.com/results?search_query=go+lang",
	}

	mapHandler := MapHandler(pathsToURL, mux)

	yaml := yamlReader("data.yaml")

	yamlHandler, err := YamlHandler([]byte(yaml), mapHandler)

	if err != nil {
		panic(fmt.Sprintf("Error Happened While Parsing YAML\n%s", err))
	}

	fmt.Println("Server Started")
	http.ListenAndServe("localhost:9090", yamlHandler)
}

func defaultMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", sayHello)
	return mux
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "Hello There")
}

func yamlReader(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Error Happened Reading YAML\n%s", err))
	}
	return content
}
