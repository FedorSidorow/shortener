package models

import (
	"github.com/FedorSidorow/shortener/internal/shortenererrors"
)

type JSONShortenRequest struct {
	URL string `json:"url"`
}

type JSONShortenResponse struct {
	Result string `json:"result"`
}

func (req *JSONShortenRequest) IsValid() error {
	if req.URL == "" {
		return &shortenererrors.ValidationError{Field: "url", Msg: "пустая строка"}
	}
	return nil
}

type ListJSONShortenRequest struct {
	CorrelationID string `json:"correlation_id,omitempty"`
	OriginalURL   string `json:"original_url,omitempty"`
}

type ListJSONShortenResponse struct {
	CorrelationID string `json:"correlation_id,omitempty"`
	ShortURL      string `json:"short_url,omitempty"`
}

func (req *ListJSONShortenRequest) IsValid() error {
	if req.CorrelationID == "" || req.OriginalURL == "" {
		return &shortenererrors.ValidationError{Field: "url", Msg: "пустая строка"}
	}
	return nil
}

type UserListJSONShortenResponse struct {
	OriginalURL string `json:"original_url,omitempty"`
	ShortURL    string `json:"short_url,omitempty"`
}
