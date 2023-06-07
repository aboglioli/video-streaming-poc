package domain

import (
	"context"
	"encoding/json"
	"errors"
)

var (
	ErrVideoNotFound = errors.New("video not found")
)

type VideoRepository interface {
	FindByID(ctx context.Context, id ID) (*Video, error)
	Save(ctx context.Context, video *Video) error
}

type Video struct {
	ID       ID
	Original Source
	Parts    []Source
}

func (v *Video) MarshalJSON() ([]byte, error) {
	video := &VideoDTO{
		ID:       v.ID.String(),
		Original: v.Original,
		Parts:    v.Parts,
	}

	return json.Marshal(video)
}
