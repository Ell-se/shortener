package main

import (
	"net/http"

	"github.com/Ell-se/shortener/internal/config"
	"github.com/Ell-se/shortener/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// обрабатываем аргументы командной строки
	config.ParseFlags()

	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("text/plain"))
	r.Post("/", handlers.AliasHandler)
	r.Get("/{id}", handlers.URLHandler)
	r.Post("/{content}", handlers.BadRequest)
	r.Get("/", handlers.BadRequest)

	err := http.ListenAndServe(config.FlagRunAddr, r)
	if err != nil {
		panic(err)
	}
}
