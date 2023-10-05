package main

import (
	"net/http"

	"github.com/Ell-se/shortener/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("text/plain"))
	r.Post("/", handlers.AliasHandler)
	r.Get("/{id}", handlers.URLHandler)
	r.Post("/{content}", handlers.BadRequest)
	r.Get("/", handlers.BadRequest)

	http.HandleFunc(`/`, r)
	err := http.ListenAndServe(Host.url+Host.port, nil)
	if err != nil {
		panic(err)
	}
}
