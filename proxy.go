package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	generalURL, _ := url.Parse("http://localhost:8083")
	faceURL, _ := url.Parse("http://localhost:8082")

	generalProxy := httputil.NewSingleHostReverseProxy(generalURL)
	faceProxy := httputil.NewSingleHostReverseProxy(faceURL)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/face") {
			r.URL.Path = strings.TrimPrefix(r.URL.Path, "/face")
			faceProxy.ServeHTTP(w, r)
			return
		}
		generalProxy.ServeHTTP(w, r)
	})

	log.Println("Reverse proxy listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
