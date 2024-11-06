package app

import (
	"lockStock/internal/config"
	"lockStock/internal/network/http"
	"lockStock/internal/server"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		panic(err)
	}

	handlers := http.NewHandler(usecase.User)
	srv = server.NewServer(cfg, handlers.Init())
}
