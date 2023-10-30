package main

import (
	"net/http"

	"github.com/Ell-se/shortener/internal/config"
	"github.com/Ell-se/shortener/internal/handlers"
	"github.com/Ell-se/shortener/internal/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// обрабатываем аргументы командной строки
	config.ParseFlags()

	r := chi.NewRouter()
	h := handlers.Handlers{}
	r.Use(middleware.AllowContentType("text/plain"))
	r.Post("/", logger.RequestLogger(h.AliasHandler))
	r.Get("/{id}", logger.RequestLogger(h.URLHandler))
	r.Post("/{content}", logger.RequestLogger(h.BadRequest))
	r.Get("/", logger.RequestLogger(h.BadRequest))

	err := http.ListenAndServe(config.FlagRunAddr, r)
	if err != nil {
		panic(err)
	}
}
