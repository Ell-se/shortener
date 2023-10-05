package config

import (
	"flag"
)

var FlagRunAddr string
var FlagShortAddr string
var Host string

// parseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных
func ParseFlags() {

	Host = "http://localhost"
	flag.StringVar(&FlagRunAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&FlagShortAddr, "b", ":8080", "address and port to short")

	// парсим переданные серверу аргументы в зарегистрированные переменные
	flag.Parse()
}
