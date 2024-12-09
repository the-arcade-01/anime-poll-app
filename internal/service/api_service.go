package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/the-arcade-01/anime-poll-app/internal/cache"
	"github.com/the-arcade-01/anime-poll-app/internal/models"
	"github.com/the-arcade-01/anime-poll-app/internal/repository"
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

func (service *ApiService) FetchAllAnimes(w http.ResponseWriter, r *http.Request) {
	animes, err := service.repo.FetchAllAnime()
	if err != nil {
		models.ResponseWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	models.ResponseWithJSON(w, http.StatusOK, animes)
}

func (service *ApiService) GetAnimesForFaceOff(w http.ResponseWriter, r *http.Request) {
	animes := cache.GetTwoRandomAnime()
	models.ResponseWithJSON(w, http.StatusOK, animes)
}

func (service *ApiService) VoteAnime(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	malId, err := strconv.Atoi(id)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusBadRequest, "please provide correct id")
		return
	}
	err = service.repo.UpsertAnimeVote(malId)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	models.ResponseWithJSON(w, http.StatusOK, nil)
}

func (service *ApiService) GetLeaderBoard(w http.ResponseWriter, r *http.Request) {
	num := chi.URLParam(r, "count")
	count, err := strconv.Atoi(num)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusBadRequest, "please provide correct number input")
		return
	}
	animes, err := service.repo.GetTopAnimesLeaderBoard(count)
	if err != nil {
		models.ResponseWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	models.ResponseWithJSON(w, http.StatusOK, animes)
}
