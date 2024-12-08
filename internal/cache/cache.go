package cache

import (
	"math/rand"
	"sync"
	"time"

	"github.com/the-arcade-01/anime-poll-app/internal/models"
)

var mtx = sync.Mutex{}
var AnimeDetailsCache = map[int]*models.DBAnimeDetails{}

func CacheAnimeDetails(animes []*models.DBAnimeDetails) {
	mtx.Lock()
	defer mtx.Unlock()
	AnimeDetailsCache = map[int]*models.DBAnimeDetails{}
	for _, anime := range animes {
		AnimeDetailsCache[anime.MalId] = anime
	}
}

func GetTwoRandomAnime() []*models.DBAnimeDetails {
	mtx.Lock()
	defer mtx.Unlock()

	if len(AnimeDetailsCache) < 2 {
		return []*models.DBAnimeDetails{}
	}

	rand.Seed(time.Now().UnixNano())
	keys := make([]int, 0, len(AnimeDetailsCache))
	for key := range AnimeDetailsCache {
		keys = append(keys, key)
	}

	rand.Shuffle(len(keys), func(i, j int) { keys[i], keys[j] = keys[j], keys[i] })

	return []*models.DBAnimeDetails{
		AnimeDetailsCache[keys[0]],
		AnimeDetailsCache[keys[1]],
	}
}
