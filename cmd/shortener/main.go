package main

import (
	"net/http"
	"strings"

	"github.com/Ell-se/shortener/internal/config"
	"github.com/Ell-se/shortener/internal/handlers"
)

func main() {

	http.HandleFunc(`/`, UrlRouter)
	err := http.ListenAndServe(config.Host.url+config.Host.port, nil)
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
