package ports

import (
	"context"
	"hexagonal_video_streaming/internal/core/domain"
)

type VideoRepository interface {
	FindVideo(ctx context.Context, fileName string) (*domain.Video, error)
	ReadVideoChunk(ctx context.Context, video *domain.Video, start, end int64) ([]byte, error)
}
