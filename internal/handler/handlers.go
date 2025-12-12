package handler

import (
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/FedorSidorow/shortener/internal/interfaces"
	"github.com/go-chi/chi/v5"
)

type APIHandler struct {
	shortService interfaces.ShortenerServicer
}

func NewHandler(service interfaces.ShortenerServicer) (h *APIHandler, err error) {
	log.Printf("Инициализация обработчиков событий")
	hendler := &APIHandler{
		shortService: service,
	}
	return hendler, err
}

func (h *APIHandler) GenerateShortKeyHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	urlToShort, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("error while read request body: %s\n", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	req.Body.Close()
	if string(urlToShort) == "" {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	host := req.Host
	data, err := h.shortService.GenerateShortURL(string(urlToShort))
	if err != nil {
		log.Printf("error while generate short URL: %s\n", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	shortURL, err := url.JoinPath("http://", host, data)
	if err != nil {
		log.Printf("error in JoinPath: %s\n", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	res.Header().Set("content-type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(shortURL))
}

func (h *APIHandler) GetURLByKeyHandler(res http.ResponseWriter, req *http.Request) {
	key := chi.URLParam(req, "*")
	log.Printf("Ключ полученный из chi.URLParam: %s \n", key)
	url, err := h.shortService.GetURLByKey(key)
	if err != nil {
		http.NotFound(res, req)
		log.Printf("Not found")
		return
	}
	res.Header().Set("Location", url)
	res.WriteHeader(http.StatusTemporaryRedirect)
}
