package handlers

import (
	"io/ioutil"
	"net/http"
	"strings"

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
	w.Write([]byte("http://localhost:8080/" + uf))
	w.WriteHeader(http.StatusOK)

}

func UrlHandler(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Path
	id := strings.Split(url, "/")[1]

	uf := storage.GetUrl(id)
	if uf != "" {
		w.Header().Set("Location", uf)
		w.WriteHeader(http.StatusTemporaryRedirect)
		//fmt.Println(uf)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}
