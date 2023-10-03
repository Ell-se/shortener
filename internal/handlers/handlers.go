package handlers

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Ell-se/shortener/internal/config"
	"github.com/Ell-se/shortener/internal/storage"
)

func AliasHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	if string(body) == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uf := storage.GetAlias(string(body))
	w.Write([]byte(uf))
	w.WriteHeader(http.StatusCreated)

}

func UrlHandler(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Path
	id := strings.Split(url, "/")[1]

	uf := storage.GetUrl(id)
	if uf != "" {
		w.Header().Set("Location", config.Host.protocol+config.Host.url+config.Host.port+`/`+uf)
		w.WriteHeader(http.StatusTemporaryRedirect)
		//fmt.Println(uf)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}
