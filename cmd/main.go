package main

import (
	"hexagonal_video_streaming/internal/adapters/handlers"
	"hexagonal_video_streaming/internal/adapters/repository"
	"hexagonal_video_streaming/internal/core/service"

	"github.com/gin-gonic/gin"
)

func main() {
	videoRepo := repository.New("../videos")
	videoService := service.New(videoRepo)
	videoHandler := handlers.New(videoService)

	r := gin.Default()
	r.GET("/stream/:filename", videoHandler.StreamVideo)

	r.Run(":8080")
}
