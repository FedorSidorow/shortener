package serializers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/FedorSidorow/shortener/internal/models"
)

// PostShortURLUnmarshalBody сериализатор для сокращения одного ЮРЛ.
func PostShortURLUnmarshalBody(req *http.Request) (*models.JSONShortenRequest, error) {

	defer req.Body.Close()

	var data models.JSONShortenRequest
	var buf bytes.Buffer
	_, err := buf.ReadFrom(req.Body)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(buf.Bytes(), &data); err != nil {
		return nil, err
	}

	err = data.IsValid()

	if err != nil {
		return nil, err
	}

	return &data, nil
}

// ListPostShortURLUnmarshalBody сериализатор для списка ЮРЛ.
func ListPostShortURLUnmarshalBody(req *http.Request) ([]models.ListJSONShortenRequest, error) {

	defer req.Body.Close()

	var data []models.ListJSONShortenRequest
	var buf bytes.Buffer
	_, err := buf.ReadFrom(req.Body)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(buf.Bytes(), &data); err != nil {
		return nil, err
	}

	var errs []error

	for _, v := range data {
		err := v.IsValid()
		errs = append(errs, err)
	}

	return data, errors.Join(errs...)
}

// DeleteListUserURLUnmarshalBody сериализатор для удаления ЮРЛ пользователя.
func DeleteListUserURLUnmarshalBody(req *http.Request) ([]string, error) {

	defer req.Body.Close()

	var data []string
	var buf bytes.Buffer
	_, err := buf.ReadFrom(req.Body)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(buf.Bytes(), &data); err != nil {
		return nil, err
	}

	return data, nil
}
