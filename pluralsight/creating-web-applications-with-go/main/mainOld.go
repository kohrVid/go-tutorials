package main

import (
	"bufio"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.Handle("/", new(MyHandler))
	http.ListenAndServe(":9000", nil)
}

type MyHandler struct {
	http.Handler
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := "public/" + req.URL.Path
	f, err := os.Open(path)
	if err == nil {
		var bufferedReader = bufio.NewReader(f)
		var contentType string

		switch {
		case strings.HasSuffix(path, ".css"):
			contentType = "text/css"
		case strings.HasSuffix(path, ".html"):
			contentType = "text/html"
		case strings.HasSuffix(path, ".js"):
			contentType = "application/html"
		case strings.HasSuffix(path, ".png"):
			contentType = "image/png"
		case strings.HasSuffix(path, ".mp4"):
			contentType = "video/mp4"
		default:
			contentType = "text/plain"
		}
		w.Header().Add("Content-Type", contentType)
		bufferedReader.WriteTo(w)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}
