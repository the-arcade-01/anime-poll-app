package service

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/the-arcade-01/quotes-poll-app/internal/models"
	"github.com/the-arcade-01/quotes-poll-app/internal/repository"
)

type ApiService struct {
	repo *repository.Repository
}

func NewApiService() *ApiService {
	return &ApiService{
		repo: repository.NewRepository(),
	}
}

func (service *ApiService) Greet(w http.ResponseWriter, r *http.Request) {
	models.ResponseWithJSON(w, http.StatusOK, "Hello, World!!")
}

func (service *ApiService) StartDBAnimeIngestion(w http.ResponseWriter, r *http.Request) {
	go NewIngestService().Start()
	models.ResponseWithJSON(w, http.StatusOK, "DB Anime details ingestion started.")
}

func (service *ApiService) FlushAnimeDB(w http.ResponseWriter, r *http.Request) {
	err := service.repo.FlushAnimeDB()
	if err != nil {
		models.ResponseWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	models.ResponseWithJSON(w, http.StatusOK, "Anime DB flush successfully.")
}

func (service *ApiService) DeleteAnimeById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := service.repo.DeleteAnimeById(id)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	models.ResponseWithJSON(w, http.StatusOK, fmt.Sprintf("%v anime deteled successfully.", id))
}
