package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/aboglioli/video-streaming-poc/backend/cmd/config"
	"github.com/aboglioli/video-streaming-poc/backend/cmd/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	println("[Config] Loaded")

	deps, err := config.BuildDependencies(cfg)
	if err != nil {
		panic(err)
	}

	println("[Dependencies] Built")

	go func() {
		println("[Server] Listening on", cfg.Port)

		if err := http.StartServer(cfg, deps); err != nil {
			panic(err)
		}
	}()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	<-exit
}
