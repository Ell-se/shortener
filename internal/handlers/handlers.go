package handlers

import (
	"io"
	"net/http"
	"strings"

	"github.com/Ell-se/shortener/internal/storage"
)

func AliasHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := io.ReadAll(r.Body)
	if string(body) == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uf := storage.GetAlias(string(body))
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("http://localhost:8080/" + uf))

}
func URLHandler(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Path
	id := strings.Split(url, "/")[1]

	uf := storage.GetURL(id)
	if uf != "" {
		w.Header().Set("Location", uf)
		w.WriteHeader(http.StatusTemporaryRedirect)
		//fmt.Println(uf)
	} else {
		BadRequest(w, r)
	}

}
func BadRequest(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "400 bad request", http.StatusBadRequest)
}
