package repository

import (
	"context"
	"hexagonal_video_streaming/internal/core/domain"
	"hexagonal_video_streaming/internal/core/ports"
	"os"
	"path/filepath"
)

type FileSystemRepository struct {
	basePath string
}

func New(basePath string) ports.VideoRepository {
	return &FileSystemRepository{basePath: basePath}
}

func (r *FileSystemRepository) FindVideo(ctx context.Context, fileName string) (*domain.Video, error) {
	path := filepath.Join(r.basePath, fileName)
	file, error := os.Open(path)

	if error != nil {
		return nil, error
	}
	defer file.Close()

	stat, _ := file.Stat()
	return &domain.Video{
		FileName: fileName,
		FilePath: path,
		Size:     stat.Size(),
	}, nil
}

func (r *FileSystemRepository) ReadVideoChunk(ctx context.Context, video *domain.Video, start, end int64) ([]byte, error) {

	file, err := os.Open(video.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	chunk := make([]byte, end-start)

	_, err = file.ReadAt(chunk, start)
	return chunk, err

	// OR
	//  errChan := make(chan error, 1)
	// go func() {
	//     _, err := file.ReadAt(chunk, start)
	//     errChan <- err
	// }()

	// select {
	// case err := <-errChan:
	//     return chunk, err
	// case <-ctx.Done():
	//     return nil, ctx.Err()
	// }
}
