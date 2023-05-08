package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iameggi/instagram-scrapper/internal/handler"
	"github.com/iameggi/instagram-scrapper/internal/repository"
	"github.com/iameggi/instagram-scrapper/internal/service"
)

func main() {
	r := gin.Default()

	// Enable CORS for all origins
	r.Use(cors.Default())

	instagramRepo := repository.NewInstagramRepository()
	instagramService := service.NewInstagramService(instagramRepo)
	instagramHandler := handler.NewInstagramHandler(instagramService)

	r.GET("/ig/v1/", instagramHandler.GetLatestPosts)

	r.Run(":8080")
}
