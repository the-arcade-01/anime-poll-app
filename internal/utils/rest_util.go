package utils

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/the-arcade-01/anime-poll-app/internal/models"
	"golang.org/x/time/rate"
)

type RestUtil struct {
	client    *http.Client
	rateLimit *rate.Limiter
}

func NewRestUtil() *RestUtil {
	return &RestUtil{
		client:    http.DefaultClient,
		rateLimit: rate.NewLimiter(rate.Every(1*time.Second), 3),
	}
}

func (rest *RestUtil) Get(url string) (*models.RestApiDataResponse, error) {
	if err := rest.rateLimit.Wait(context.Background()); err != nil {
		return nil, errors.New("rate limit reached")
	}
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
