package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type PaginationObject struct {
	LastVisiblePage int64          `json:"last_visible_page"`
	HasNextPage     bool           `json:"has_next_page"`
	CurrentPage     int64          `json:"current_page"`
	Items           PaginationItem `json:"items"`
}

type PaginationItem struct {
	Count   int64 `json:"count"`
	Total   int64 `json:"total"`
	PerPage int64 `json:"per_page"`
}

type RestApiDataResponse struct {
	Pagination PaginationObject `json:"pagination"`
	Data       []*ApiData       `json:"data"`
}

type ApiData struct {
	MalID  int     `json:"mal_id"`
	Images *Images `json:"images"`
	Title  string  `json:"title"`
}

type Images struct {
	JPG  *Image `json:"jpg"`
	WebP *Image `json:"webp"`
}

type Image struct {
	ImageURL      string `json:"image_url"`
	SmallImageURL string `json:"small_image_url"`
	LargeImageURL string `json:"large_image_url"`
}

type DBAnimeDetails struct {
	MalId     int    `json:"mal_id" gorm:"primaryKey"`
	Title     string `json:"title"`
	ImageLink string `json:"image_link"`
}

func (DBAnimeDetails) TableName() string {
	return "anime_details"
}

func NewDBAnimeDetails(data *ApiData) (*DBAnimeDetails, error) {
	if data.Images == nil {
		return nil, fmt.Errorf("[NewDBAnimeDetails] images is empty, %v", data)
	}

	obj := &DBAnimeDetails{
		MalId: data.MalID,
		Title: data.Title,
	}

	imageURLs := []string{
		data.Images.JPG.ImageURL,
		data.Images.JPG.SmallImageURL,
		data.Images.JPG.LargeImageURL,
		data.Images.WebP.ImageURL,
		data.Images.WebP.SmallImageURL,
		data.Images.WebP.LargeImageURL,
	}

	for _, url := range imageURLs {
		if url != "" {
			obj.ImageLink = url
			return obj, nil
		}
	}

	return nil, fmt.Errorf("[NewDBAnimeDetails] images is empty, %v", data)
}

type DBAnimeVotes struct {
	MalId int `json:"mal_id" gorm:"primaryKey"`
	Vote  int `json:"vote"`
}

func (DBAnimeVotes) TableName() string {
	return "anime_votes"
}

func ResponseWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		log.Printf("[ResponseWithJSON] error parsing payload: %v\n", payload)
	}
}
