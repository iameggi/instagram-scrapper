package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iameggi/instagram-scrapper/internal/service"
)

type InstagramHandler struct {
	service service.InstagramService
}

func NewInstagramHandler(s service.InstagramService) *InstagramHandler {
	return &InstagramHandler{service: s}
}

func (h *InstagramHandler) GetLatestPosts(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing username parameter"})
		return
	}

	posts, err := h.service.GetLatestPosts(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []map[string]interface{}
	for _, post := range posts {
		response = append(response, map[string]interface{}{
			"id":        post.ID,
			"caption":   post.Caption,
			"image_url": post.ImageURL,
			"timestamp": post.Timestamp,
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}
