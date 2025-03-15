package domain

type Video struct {
	ID       string
	FileName string
	FilePath string
	Size     int64
}

type VideoChunk struct {
	Content     []byte
	StartOffset int64
	EndOffset   int64
	TotalSize   int64
}
