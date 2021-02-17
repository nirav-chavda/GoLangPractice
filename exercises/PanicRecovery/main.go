package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

type newResponseWriter struct {
	http.ResponseWriter
	content [][]byte
}

func (nw *newResponseWriter) Write(b []byte) (int, error) {
	nw.content = append(nw.content, b)
	return len(b), nil
}

func (nw *newResponseWriter) flush() error {
	for _, data := range nw.content {
		if _, err := nw.ResponseWriter.Write(data); err != nil {
			return err
		}
	}
	return nil
}

func (nw *newResponseWriter) Flush() {
	flusher, ok := nw.ResponseWriter.(http.Flusher)
	if !ok {
		return
	}
	flusher.Flush()
}

// PanicRecoveryMiddleware s
func PanicRecoveryMiddleware(app http.Handler, showTrace bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				trace := debug.Stack()
				log.Printf("Panic %s\n%s", err, string(trace))
				if showTrace {
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprintf(w, "<h1>Panic : %s</h1><pre>%s</pre>", err, string(trace))
				} else {
					http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
					return
				}
			}
		}()
		writer := &newResponseWriter{ResponseWriter: w}
		app.ServeHTTP(writer, r)
		writer.flush()
	}
}

func arisePanic() {
	panic("Ohh God !")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Print This Hello Message")
}

func panicSituation(w http.ResponseWriter, r *http.Request) {
	arisePanic()
}

func panicAfterSituation(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Message before panic")
	arisePanic()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/panic/", panicSituation)
	// Situation where you already write something to writer and then panic happens
	// It will result in 200
	// While displaying all the content including data and error messages
	mux.HandleFunc("/panic-after/", panicAfterSituation)
	log.Println("Server Started")
	http.ListenAndServe("localhost:9090", PanicRecoveryMiddleware(mux, true))
}
