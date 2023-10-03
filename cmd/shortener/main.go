package main

import (
	"net/http"
	"strings"

	"github.com/Ell-se/shortener/internal/handlers"
)

// host struct
var Host struct {
	protocol string
	url      string
	port     string
}

func main() {
	Host.url = "localhost"
	Host.port = ":8080"
	Host.protocol = `http://`
	http.HandleFunc(`/`, UrlRouter)
	err := http.ListenAndServe(Host.url+Host.port, nil)
	if err != nil {
		panic(err)
	}
}

func UrlRouter(res http.ResponseWriter, req *http.Request) {

	url := req.URL.Path
	method := req.Method
	id := strings.Split(url, "/")[1]

	res.Header().Set("content-type", "text/plain")
	if id != "" && method == http.MethodGet {
		handlers.UrlHandler(res, req)
	} else if id == "" && method == http.MethodPost {
		handlers.AliasHandler(res, req)
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}

}
