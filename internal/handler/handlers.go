package handler

import (
	"net/http"

	"github.com/FedorSidorow/shortener/internal/service"
)

func ShortThisURL(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	url := req.FormValue("url")
	data := service.ShortURL(url)
	res.Header().Set("content-type", "text/plain")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(data))
}

func GetFullURL(res http.ResponseWriter, req *http.Request) {
	key := req.PathValue("key")
	url, err := service.ReturnFullURL(key)

	if err != nil {
		http.NotFound(res, req)
	}

	http.Redirect(res, req, url, http.StatusMovedPermanently)
}
