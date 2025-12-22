package serializers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/FedorSidorow/shortener/internal/shortenererrors"
)

type JsonShortenRequest struct {
	URL string `json:"url"`
}

type JsonShortenResponse struct {
	Result string `json:"result"`
}

func (req *JsonShortenRequest) isValid() error {
	if req.URL == "" {
		return &shortenererrors.ValidationError{Field: "url", Msg: "пустая строка"}
	}
	return nil
}

func PostShortURLUnmarshalBody(req *http.Request) (*JsonShortenRequest, error) {

	defer req.Body.Close()

	var data JsonShortenRequest
	var buf bytes.Buffer
	_, err := buf.ReadFrom(req.Body)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(buf.Bytes(), &data); err != nil {
		return nil, err
	}

	err = data.isValid()

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func PostShortURLMarshalBody(data *JsonShortenResponse) (*[]byte, error) {

	bytesToReturn, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return &bytesToReturn, nil
}
