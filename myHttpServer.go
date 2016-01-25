package main

import (
	"net/http"
)

func servePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,"static/" + r.URL.Path[1:])
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	server := http.Server{
		Addr:    ":8000",
		Handler: &myHandler{},
	}
	
	server.ListenAndServe()
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	servePage(w,r)
}
