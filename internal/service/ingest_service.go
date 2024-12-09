package service

import (
	"fmt"
	"log"
	"sync"

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
	wg := sync.WaitGroup{}
	pageChan := make(chan int, 4)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for page := range pageChan {
				url := fmt.Sprintf("%v?page=%v", ingest.url, page)
				res, err := ingest.restClient.Get(url)
				if err != nil {
					log.Println(err)
					continue
				}
				ingest.process(res)
			}
		}()
	}

	for i := 1; i <= 20; i++ {
		pageChan <- i
	}

	close(pageChan)
	wg.Wait()
	log.Println("[IngestService] anime details db ingestion completed")
}

func (ingest *IngestService) process(res *models.RestApiDataResponse) {
	var batch []*models.DBAnimeDetails
	for _, item := range res.Data {
		dbItem, err := models.NewDBAnimeDetails(item)
		if err != nil {
			log.Println(err)
			continue
		}
		batch = append(batch, dbItem)
	}
	err := ingest.repo.InsertAnimeBatch(batch)
	if err != nil {
		log.Printf("[IngestService] error while inserting batch in db, %v\n", err)
		return
	}
}
