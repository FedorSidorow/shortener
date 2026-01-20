package handler

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/FedorSidorow/shortener/internal/interfaces"
	"github.com/FedorSidorow/shortener/internal/models"
	"github.com/FedorSidorow/shortener/internal/serializers"
	"github.com/FedorSidorow/shortener/internal/shortenererrors"
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

	shortURL, err := h.shortService.GenerateShortURL(string(urlToShort), req.Host)
	if err != nil {
		log.Printf("error while generate short URL: %s\n", err)
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

func (h *APIHandler) JSONGenerateShortkeyHandler(res http.ResponseWriter, req *http.Request) {

	var (
		data          *models.JSONShortenRequest
		responseData  models.JSONShortenResponse
		err           error
		validationErr *shortenererrors.ValidationError
	)

	data, err = serializers.PostShortURLUnmarshalBody(req)
	if err != nil {
		switch {
		case errors.As(err, &validationErr):
			log.Printf("validation error: %s\n", err)
			http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		default:
			log.Printf("error in PostShortURLUnmarshalBody: %s\n", err)
			http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	shortURL, err := h.shortService.GenerateShortURL(data.URL, req.Host)
	if err != nil {
		log.Printf("error while generate short URL: %s\n", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseData.Result = shortURL

	response, err := json.Marshal(responseData)
	if err != nil {
		log.Printf("error while serializing: %s\n", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	res.Write(response)
}

func (h *APIHandler) PingDB(res http.ResponseWriter, req *http.Request) {
	log.Print("Проверка состояния подключения к БД")
	if ok := h.shortService.PingStorage(); !ok {
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
}

func (h *APIHandler) ListJSONGenerateShortkeyHandler(res http.ResponseWriter, req *http.Request) {
	var (
		validationErr *shortenererrors.ValidationError
		ctx           = req.Context()
	)

	data, err := serializers.ListPostShortURLUnmarshalBody(req)
	if err != nil {
		switch {
		case errors.As(err, &validationErr):
			log.Printf("validation error: %s\n", err)
			http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		default:
			log.Printf("error in PostShortURLUnmarshalBody: %s\n", err)
			http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	responseData, err := h.shortService.ListGenerateShortURL(ctx, data, req.Host)
	if err != nil {
		log.Printf("error while creating rows: %s\n", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(responseData)
	if err != nil {
		log.Printf("error while serializing: %s\n", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	res.Write(response)
}
