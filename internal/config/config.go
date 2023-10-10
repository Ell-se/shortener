package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS"`
	BaseUrl       string `env:"BASE_URL"`
}

var FlagRunAddr string
var FlagShortAddr string
var Host string

// parseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных
func ParseFlags() {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	Host = "http://localhost"
	flag.StringVar(&FlagRunAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&FlagShortAddr, "b", "http://localhost:8080", "address and port to short")

	// парсим переданные серверу аргументы в зарегистрированные переменные
	flag.Parse()
	// выбираем значения
	if cfg.BaseUrl != "" {
		FlagShortAddr = cfg.BaseUrl
	}
	if cfg.ServerAddress != "" {
		FlagRunAddr = cfg.ServerAddress
	}

}
