package models

import "github.com/FedorSidorow/shortener/internal/shortenererrors"

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
