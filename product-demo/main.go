package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nirav-chavda/practice/product-demo/handlers"
)

func main() {

	logger := log.New(os.Stdout, "demo-api", log.LstdFlags)
	sm := http.NewServeMux()

	helloHandler := handlers.HelloHandler(logger)

	// Prints to server log
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello There")
	})

	// Prints to response
	sm.Handle("/greet", helloHandler)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start Server
	http.ListenAndServe("localhost:9090", sm)
}
