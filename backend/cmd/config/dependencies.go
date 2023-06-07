package config

import (
	"github.com/aboglioli/video-streaming-poc/backend/domain"
	"github.com/aboglioli/video-streaming-poc/backend/infrastructure"
)

type Dependencies struct {
	VideoRepository domain.VideoRepository
}

func BuildDependencies(config *Config) (*Dependencies, error) {
	return &Dependencies{
		VideoRepository: infrastructure.NewInMemVideoRepository(),
	}, nil
}
