# Video Streaming Service (Hexagonal Architecture)

This is a simple video streaming service built using Golang and Gin, following Hexagonal Architecture principles. It allows you to stream video files chunk by chunk, supporting range requests for efficient video playback, I'll be updating this.

## Features

* Chunked Video Streaming: Streams video files in configurable chunks

* Range Requests: Supports HTTP range headers for seeking and partial content

* Hexagonal Architecture: Clean separation of concerns with ports and adapters

* File System Storage: Stores and serves videos from a local directory

## Installation

1. Clone the repository:

       git clone https://github.com/alirezaporthos/go-simple-video-streaming-hexagonal
       cd go-simple-video-streaming-hexagonal
2. Install dependencies:

       go mod download

3. Create a videos directory:
  
       mkdir videos

4. Add a video file to the videos directory:

5. Run the app:

       go run cmd/main.go

6. Access the video stream:

    * Open your browser or video player and navigate to:

          http://localhost:8080/stream/your-video.mp4
