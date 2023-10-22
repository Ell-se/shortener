package handlers

import (
	"io"
	"net/http"
	"strings"

	"github.com/Ell-se/shortener/internal/config"
	"github.com/Ell-se/shortener/internal/storage"
)

// структура хендлеров
type Handlers struct {
}

func (h *Handlers) AliasHandler(w http.ResponseWriter, r *http.Request) {

	body, _ := io.ReadAll(r.Body)
	if string(body) == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uf := storage.GetAlias(string(body))
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(config.FlagShortAddr + `/` + uf))
}
func (h *Handlers) URLHandler(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Path
	id := strings.Split(url, "/")[1]

	uf := storage.GetURL(id)
	if uf == "" {
		h.BadRequest(w, r)
		return
	}
	w.Header().Set("Location", uf)
	w.WriteHeader(http.StatusTemporaryRedirect)
	//fmt.Println(uf)

}

/*
	используется не в одном месте, еще в main. Не придумала как реализовать все остальное поведение, короме описанного

Вдруг дальше будет градация ошибок, переделаю ее в разные варианты ответов
*/
func (h *Handlers) BadRequest(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "400 bad request", http.StatusBadRequest)
}
