package handler

import (
	"io"
	"net/http"

	"github.com/FedorSidorow/shortener/internal/service"
)

func ShortThisURL(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	url, _ := io.ReadAll(req.Body)
	host := req.Host
	data := service.ShortURL(string(url))
	res.Header().Set("content-type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(host + "/" + data))
}

func GetFullURL(res http.ResponseWriter, req *http.Request) {
	key := req.PathValue("key")
	url, err := service.ReturnFullURL(key)

	if err != nil {
		http.NotFound(res, req)
	}
	res.Header().Set("Location", url)
	res.WriteHeader(http.StatusTemporaryRedirect)
}
