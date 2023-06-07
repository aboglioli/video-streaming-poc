package http

import (
	"fmt"
	"net/http"

	"github.com/aboglioli/video-streaming-poc/backend/cmd/config"
	"github.com/go-chi/chi/v5"
)

func StartServer(config *config.Config, deps *config.Dependencies) error {
	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/videos/{id}", getVideo(config, deps))
	})

	return http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r)
}
