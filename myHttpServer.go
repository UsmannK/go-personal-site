package main

import (
	"net/http"
	"fmt"
)

func servePersonalSite(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,"static/" + r.URL.Path[1:])
}

func serveKanyeSite(w httpResponseWriter, r *http.Request) {
	http.ServeFile(w,r,"kanye/")
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	server := http.Server{
		Addr:    ":80",
		Handler: &myHandler{},
	}
	
	server.ListenAndServe()
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if string.Contains(string,ToLower(r.Host), "iswavesoutyet.com") {
		serveKanye(w,r)
	}
	else {
		servePersonalSite(w,r)
	}
}
