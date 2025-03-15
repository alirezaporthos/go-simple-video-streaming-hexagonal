package service

import (
	"context"
	"hexagonal_video_streaming/internal/core/domain"
	"hexagonal_video_streaming/internal/core/ports"
)

type VideoService struct {
	repo ports.VideoRepository
}

func New(repo ports.VideoRepository) *VideoService {
	return &VideoService{repo: repo}
}

func (service *VideoService) GetVideo(ctx context.Context, fileName string) (*domain.Video, error) {
	return service.repo.FindVideo(ctx, fileName)
}

func (service *VideoService) GetVideoChunk(ctx context.Context, fileName string, start, end int64) (*domain.VideoChunk, error) {
	video, err := service.GetVideo(ctx, fileName)
	if err != nil {
		return nil, err
	}

	data, err := service.repo.ReadVideoChunk(ctx, video, start, end)
	if err != nil {
		return nil, err
	}

	return &domain.VideoChunk{
		Content:     data,
		StartOffset: start,
		EndOffset:   end,
		TotalSize:   video.Size,
	}, nil
}
