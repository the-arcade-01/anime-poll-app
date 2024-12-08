package service

import (
	"log"

	"github.com/the-arcade-01/anime-poll-app/internal/config"
	"github.com/the-arcade-01/anime-poll-app/internal/models"
	"github.com/the-arcade-01/anime-poll-app/internal/repository"
	"github.com/the-arcade-01/anime-poll-app/internal/utils"
)

type IngestService struct {
	url        string
	restClient *utils.RestUtil
	repo       *repository.Repository
}

func NewIngestService() *IngestService {
	appConfig := config.NewAppConfig()
	return &IngestService{
		restClient: utils.NewRestUtil(),
		url:        appConfig.ApiUrl,
		repo:       repository.NewRepository(),
	}
}

func (ingest *IngestService) Start() {
	res, err := ingest.restClient.Get(ingest.url)
	if err != nil {
		log.Println(err)
		return
	}
	var batch []*models.DBAnimeDetails
	for _, item := range res.Data {
		dbItem, err := models.NewDBAnimeDetails(item)
		if err != nil {
			log.Println(err)
			continue
		}
		batch = append(batch, dbItem)
	}
	err = ingest.repo.InsertAnimeBatch(batch)
	if err != nil {
		log.Printf("[IngestService] error while inserting batch in db, %v\n", err)
		return
	}
	log.Println("[IngestService] anime details db ingestion completed")
}
