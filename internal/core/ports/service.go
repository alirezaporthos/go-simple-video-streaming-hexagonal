package ports

import (
	"context"
	"hexagonal_video_streaming/internal/core/domain"
)

type videoService interface {
	GetVideo(ctx context.Context, fileName string) (*domain.Video, error)
	GetVideoChunk(ctx context.Context, fileName string, start, end int64) (*domain.VideoChunk, error)
}
