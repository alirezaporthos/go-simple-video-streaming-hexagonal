package handlers

import (
	"fmt"
	"hexagonal_video_streaming/internal/core/domain"
	"hexagonal_video_streaming/internal/core/ports"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VideoHandler struct {
	service ports.VideoService
}

func New(service ports.VideoService) *VideoHandler {
	return &VideoHandler{service: service}
}

func (handler *VideoHandler) StreamVideo(c *gin.Context) {
	fileName := c.Param("filename")
	ctx := c.Request.Context()

	video, error := handler.service.GetVideo(ctx, fileName)
	if error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found."})
		return
	}

	rangeHeader := c.GetHeader("Range")
	start, end := parseRange(rangeHeader, video.Size)

	chunk, err := handler.service.GetVideoChunk(ctx, fileName, start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read chunk."})
		return
	}

	c.Header("Content-Type", "video/mp4")
	c.Header("Content-Length", strconv.FormatInt(end-start, 10))
	c.Header("Content-Range", contentRangeHeader(chunk))
	c.Status(http.StatusPartialContent)

	c.Writer.Write(chunk.Content)
}

func parseRange(rangeHeader string, size int64) (int64, int64) {
	var start, end int64
	fmt.Sscanf(rangeHeader, "bytes=%d-%d", &start, &end)

	if end == 0 || end > size {
		end = size
	}
	return start, end
}

func contentRangeHeader(chunk *domain.VideoChunk) string {
	return "bytes" +
		strconv.FormatInt(chunk.StartOffset, 10) + "-" +
		strconv.FormatInt(chunk.EndOffset, 10) + "/" +
		strconv.FormatInt(chunk.TotalSize, 10)
}
