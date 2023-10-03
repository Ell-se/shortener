package main

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
)

// хранение данных
var DataUrls = map[string]string{
	"EwHXdJfB": "https://practicum.yandex.ru/",
}

// host struct
var Host struct {
	protocol string
	url      string
	port     string
}

// для создания строки
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// возврат короткого адреса
func getAlias(url string) string {
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
	shortUrl := randSeq(8)
	DataUrls[shortUrl] = url
	return shortUrl
}

// создание рандомной строки
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// возврат длинного адреса
func getUrl(alias string) string {
	for k, v := range DataUrls {
		//if strings.ToLower(k) == strings.ToLower(alias) {
		if k == alias {
			return v
		}

	}
	return ""
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
		UrlHandle(res, req)
	} else if id == "" && method == http.MethodPost {
		AliasHandle(res, req)
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}

}

func AliasHandle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	if string(body) == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uf := getAlias(string(body))
	w.Write([]byte(Host.protocol + Host.url + Host.port + `/` + uf))
	w.WriteHeader(http.StatusCreated)

}

func UrlHandle(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Path
	id := strings.Split(url, "/")[1]

	uf := getUrl(id)
	if uf != "" {
		w.Header().Set("Location", uf)
		w.WriteHeader(http.StatusTemporaryRedirect)
		//fmt.Println(uf)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}
