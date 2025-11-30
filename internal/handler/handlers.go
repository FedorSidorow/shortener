package handler

import (
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

type APIHandler struct {
	shortService ShortenerServicer
}

func NewHandler(service ShortenerServicer) (h *APIHandler, err error) {
	hendler := &APIHandler{
		shortService: service,
	}

	return hendler, err
}

type ShortenerServicer interface {
	GetURLByKey(key string) (string, error)
	GenerateShortURL(URL string) (string, error)
}

func (h *APIHandler) GenerateShortKeyHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	url, _ := io.ReadAll(req.Body)
	req.Body.Close()
	if string(url) == "" {
		http.Error(res, "Bad Request", http.StatusBadRequest)
		return
	}

	host := req.Host
	data, err := h.shortService.GenerateShortURL(string(url))
	if err != nil {
		http.Error(res, "", http.StatusInternalServerError)
		return
	}

	res.Header().Set("content-type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte("http://" + host + "/" + data))
}

func (h *APIHandler) GetURLByKeyHandler(res http.ResponseWriter, req *http.Request) {
	key := chi.URLParam(req, "key")
	url, err := h.shortService.GetURLByKey(key)
	if err != nil {
		http.NotFound(res, req)
		return
	}
	res.Header().Set("Location", url)
	res.WriteHeader(http.StatusTemporaryRedirect)
}
