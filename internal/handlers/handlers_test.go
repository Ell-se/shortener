package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLHandler(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}
	tests := []struct {
		name    string
		request string
		want    want
	}{
		{
			name:    "positive test #1",
			request: "/EwHXdJfB",
			want: want{
				code:        307,
				response:    `https://practicum.yandex.ru/`,
				contentType: "text/plain",
			},
		},
		{
			name:    "positive test #2",
			request: "/111",
			want: want{
				code:        400,
				response:    ``,
				contentType: "text/plain",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, test.request, nil)
			// создаём новый Recorder
			w := httptest.NewRecorder()
			URLHandler(w, request)

			res := w.Result()
			// проверяем код ответа
			assert.Equal(t, test.want.code, res.StatusCode)
			defer res.Body.Close()
			assert.Equal(t, test.want.response, res.Header.Get("Location"))
			fmt.Println(res.Header.Get("Location"))
		})
	}
}

func TestAliasHandler(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}
	tests := []struct {
		name    string
		request string
		want    want
	}{
		{
			name:    "positive test #1",
			request: "https://practicum.yandex.ru/",
			want: want{
				code:        201,
				response:    `http://localhost:8080/EwHXdJfB`,
				contentType: "text/plain",
			},
		},
		{
			name:    "positive test #2",
			request: "",
			want: want{
				code:        400,
				response:    ``,
				contentType: "text/plain",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reqBody := strings.NewReader(test.request)
			request := httptest.NewRequest(http.MethodPost, "/", reqBody)
			// создаём новый Recorder
			w := httptest.NewRecorder()
			AliasHandler(w, request)

			res := w.Result()
			// проверяем код ответа
			assert.Equal(t, test.want.code, res.StatusCode)
			// получаем и проверяем тело запроса
			defer res.Body.Close()
			resBody, _ := io.ReadAll(res.Body)
			// require.NoError(t, err)
			assert.Equal(t, test.want.response, string(resBody))

		})
	}
}
