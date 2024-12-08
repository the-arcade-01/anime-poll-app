package repository

import (
	"github.com/the-arcade-01/anime-poll-app/internal/cache"
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
	cache.CacheAnimeDetails(animes)
	return animes, nil
}

func (repo *Repository) UpsertAnimeVote(malId int) error {
	return repo.dbClient.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&models.DBAnimeVotes{}).
			Where(models.DBAnimeVotes{MalId: malId}).
			Attrs(models.DBAnimeVotes{Vote: 1}).
			UpdateColumn("vote", gorm.Expr("CASE WHEN vote = 0 THEN ? ELSE vote + 1 END", 1))

		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return tx.Create(&models.DBAnimeVotes{MalId: malId, Vote: 1}).Error
		}

		return nil
	})
}
