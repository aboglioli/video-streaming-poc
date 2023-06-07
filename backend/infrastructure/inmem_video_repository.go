package infrastructure

import (
	"context"
	"fmt"

	"github.com/aboglioli/video-streaming-poc/backend/domain"
)

var _ domain.VideoRepository = (*InMemVideoRepository)(nil)

type InMemVideoRepository struct {
	Items map[domain.ID]*domain.Video
}

func NewInMemVideoRepository() *InMemVideoRepository {
	return &InMemVideoRepository{
		Items: make(map[domain.ID]*domain.Video),
	}
}

func (r *InMemVideoRepository) FindByID(ctx context.Context, id domain.ID) (*domain.Video, error) {
	if video, ok := r.Items[id]; ok {
		return video, nil
	}

	return nil, fmt.Errorf("%w: %s", domain.ErrVideoNotFound, id)
}

func (r *InMemVideoRepository) Save(ctx context.Context, video *domain.Video) error {
	r.Items[video.ID] = video
	return nil
}
