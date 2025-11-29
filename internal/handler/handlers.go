package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/FedorSidorow/shortener/internal/service"
)

func ShortThisURL(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Only POST requests are allowed!", http.StatusMethodNotAllowed)
		return
	}
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	url := req.FormValue("url")
	data := service.ShortURL(url)
	fmt.Fprint(res, data)
}

func GetFullURL(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Only GET requests are allowed!", http.StatusMethodNotAllowed)
		return
	}
	key := strings.TrimPrefix(req.URL.Path, "/")
	url, err := service.ReturnFullURL(key)

	if err != nil {
		http.NotFound(res, req)
	}

	http.Redirect(res, req, url, http.StatusMovedPermanently)
}
