package serializers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/FedorSidorow/shortener/internal/models"
)

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
