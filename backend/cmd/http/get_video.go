package http

import (
	"encoding/json"
	"net/http"

	"github.com/aboglioli/video-streaming-poc/backend/application"
	"github.com/aboglioli/video-streaming-poc/backend/cmd/config"
	"github.com/go-chi/chi/v5"
)

func getVideo(config *config.Config, deps *config.Dependencies) http.HandlerFunc {
	serv := application.NewGetVideo(deps.VideoRepository)

	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		cmd := &application.GetVideoCommand{
			ID: id,
		}

		resp, err := serv.Exec(r.Context(), cmd)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(resp)
	}
}
