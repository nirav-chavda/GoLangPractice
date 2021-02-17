package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func HelloHandler(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello Request Served")

	d, err := ioutil.ReadAll(r.Body)

	if err != nil || string(d) == "" {
		http.Error(rw, "Body not provided", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello There, %s", d)
}
