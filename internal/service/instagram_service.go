package service

import (
	"github.com/iameggi/instagram-scrapper/internal/model"
	"github.com/iameggi/instagram-scrapper/internal/repository"
)

type InstagramService interface {
	GetLatestPosts(username string) ([]model.Post, error)
}

type instagramService struct {
	repo *repository.InstagramRepository
}

func NewInstagramService(r *repository.InstagramRepository) InstagramService {
	return &instagramService{repo: r}
}

func (s *instagramService) GetLatestPosts(username string) ([]model.Post, error) {
	data, err := s.repo.GetInstagramData(username)
	if err != nil {
		return nil, err
	}

	var posts []model.Post
	for i, edge := range data.GraphQL.User.EdgeOwnerToTimelineMedia.Edges {
		if i == 6 {
			break
		}
		post := model.Post{
			ID:        edge.Node.ID,
			Caption:   edge.Node.EdgeMediaToCaption.Edges[0].Node.Text,
			ImageURL:  edge.Node.ThumbnailSrc,
			Timestamp: edge.Node.TakenAtTimestamp,
		}
		posts = append(posts, post)
	}
	return posts, nil
}
