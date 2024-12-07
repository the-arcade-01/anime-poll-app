package utils

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/the-arcade-01/quotes-poll-app/internal/models"
)

type RestUtil struct {
	client *http.Client
}

func NewRestUtil() *RestUtil {
	return &RestUtil{
		client: http.DefaultClient,
	}
}

func (rest *RestUtil) Get(url string) (*models.RestApiDataResponse, error) {
	res, err := rest.client.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode < 200 || res.StatusCode >= 400 {
		return nil, errors.New("bad response code")
	}
	var apiRes models.RestApiDataResponse
	if err := json.NewDecoder(res.Body).Decode(&apiRes); err != nil {
		return nil, err
	}
	return &apiRes, nil
}
