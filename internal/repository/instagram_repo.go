package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iameggi/instagram-scrapper/internal/entity"
)

type InstagramRepository struct{}

func NewInstagramRepository() *InstagramRepository {
	return &InstagramRepository{}
}

func (r *InstagramRepository) GetInstagramData(username string) (*entity.InstagramData, error) {
	url := fmt.Sprintf("https://www.instagram.com/%s/?__a=1", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data entity.InstagramData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
