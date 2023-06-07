package application

import (
	"context"

	"github.com/aboglioli/video-streaming-poc/backend/domain"
)

type GetVideo struct {
	videoRepository domain.VideoRepository
}

type GetVideoCommand struct {
	ID string `json:"id"`
}

type GetVideoResponse struct {
	ID       string   `json:"id"`
	Original string   `json:"original"`
	Parts    []string `json:"parts"`
}

func NewGetVideo(videoRepository domain.VideoRepository) *GetVideo {
	return &GetVideo{
		videoRepository: videoRepository,
	}
}

func (s *GetVideo) Exec(ctx context.Context, cmd *GetVideoCommand) (*domain.Video, error) {
	id, err := domain.NewID(cmd.ID)
	if err != nil {
		return nil, err
	}

	return s.videoRepository.FindByID(ctx, id)
}
