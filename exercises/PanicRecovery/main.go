package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime/debug"
	"strings"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
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
				if showTrace {
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprintf(w, "<h1>Panic : %s</h1><pre>%s</pre>", err, addHyperLinks(string(trace)))
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

func codeHandler(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path") //"F:/work/go/src/github.com/nirav-chavda/practice/exercises/PanicRecovery/main.go"
	file, err := os.Open(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	buffer := bytes.NewBuffer(nil)
	_, err = io.Copy(buffer, file)
	if err != nil {
		panic(err.Error())
	}
	
	lexer := lexers.Get("go")
	if lexer == nil {
		lexer = lexers.Fallback
	}
	lexer = chroma.Coalesce(lexer)
	style := styles.Get("github")
	formatter := html.New(html.WithClasses(false))
	err = formatter.WriteCSS(w, style)
	iterator, err := lexer.Tokenise(nil, buffer.String())
	err = formatter.Format(w, style, iterator)
	//if err = quick.Highlight(w, buffer.String(), "go", "html", "github"); err != nil {
	if err != nil {
		log.Println(err.Error())
	}
	//}
}

func addHyperLinks(stack string) string {
	lines := strings.Split(stack, "\n")
	for i, line := range lines {
		filename := ""
		if len(line) == 0 || line[0] != '\t' {
			continue
		}
		for index, char := range line {
			if char == ':' && line[index+1] != '/' {
				filename = line[1:index] // 1 - because tab is considered as one character, which is at place 0
				break
			}
		}
		values := url.Values{}
		values.Set("path", filename)
		lines[i] = "\t<a href='/debug?" + values.Encode() + "' target='_blank'>" + filename + "</a>" + line[len(filename)+1:]
	}
	return strings.Join(lines, "\n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/panic/", panicSituation)
	// Situation where you already write something to writer and then panic happens
	// It will result in 200
	// While displaying all the content including data and error messages
	mux.HandleFunc("/panic-after/", panicAfterSituation)
	mux.HandleFunc("/debug/", codeHandler)
	log.Println("Server Started")
	http.ListenAndServe("localhost:9090", PanicRecoveryMiddleware(mux, true))
}
