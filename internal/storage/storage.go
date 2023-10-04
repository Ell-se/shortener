package storage

import (
	"github.com/Ell-se/shortener/internal/rand"
)

// хранение данных
var DataUrls = map[string]string{
	"EwHXdJfB": "https://practicum.yandex.ru/",
}

// возврат короткого адреса
func GetAlias(url string) string {
	for k, v := range DataUrls {
		//if strings.ToLower(v) == strings.ToLower(url) {
		if v == url {
			return k
		}
	}
	// если такого длинного адреса нет, то создается короткий и добавляется в мапу
	return setAlias(url)
}

func setAlias(url string) string {
	shortURL := rand.RandSeq(8)
	DataUrls[shortURL] = url
	return shortURL
}

// возврат длинного адреса
func GetURL(alias string) string {
	for k, v := range DataUrls {
		//if strings.ToLower(k) == strings.ToLower(alias) {
		if k == alias {
			return v
		}

	}
	return ""
}
