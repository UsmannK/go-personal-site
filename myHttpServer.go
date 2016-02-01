package main

import (
	"html/template"
	"net/http"
	"time"
	"strings"
	"fmt"
)

type Page struct {
	Answer string
	Body []byte
}

func servePersonalSite(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,"static/" + r.URL.Path[1:])	
}

func serveKanyeSite(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.URL.Path, "static") {
		answer := "NO"
		if isWaves() {
			answer = "YES"
		}
		p := &Page{Answer: answer}
		t, _ := template.ParseFiles("kanye/index.html")
		t.Execute(w, p)
	} else {
			fmt.Print("static/" + r.URL.Path[8:])
			http.ServeFile(w,r,"static/" + r.URL.Path[8:])	
	}
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
	if strings.Contains(strings.ToLower(r.Host), "iswavesoutyet.com") {
		serveKanyeSite(w,r)
	} else {
		servePersonalSite(w,r)
	}
}

func isWaves() (bool) {
	waves,_ := time.Parse(time.RFC822, "11 Feb 16 00:00")
	now := time.Now()
	return waves.After(now)
}
