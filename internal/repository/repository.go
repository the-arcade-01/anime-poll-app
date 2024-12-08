package repository

import (
	"github.com/the-arcade-01/anime-poll-app/internal/config"
	"github.com/the-arcade-01/anime-poll-app/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	dbClient *gorm.DB
}

func NewRepository() *Repository {
	appConfig := config.NewAppConfig()
	return &Repository{
		dbClient: appConfig.DbClient,
	}
}

func (repo *Repository) InsertAnimeBatch(batch []*models.DBAnimeDetails) error {
	err := repo.dbClient.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&batch).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) FlushAnimeDB() error {
	if err := repo.dbClient.Exec("TRUNCATE TABLE anime_details").Error; err != nil {
		return err
	}
	return nil
}

func (repo *Repository) DeleteAnimeById(id string) error {
	if err := repo.dbClient.Exec("DELETE FROM anime_details WHERE mal_id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (repo *Repository) FetchAllAnime() ([]*models.DBAnimeDetails, error) {
	var animes []*models.DBAnimeDetails
	err := repo.dbClient.Find(&animes).Error
	if err != nil {
		return nil, err
	}
	return animes, nil
}
