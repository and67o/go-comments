package app

import "github.com/and67o/go-comments/internal/interfaces"

type App struct {
	Config  interfaces.Config
	Storage interfaces.Storage
	Redis   interfaces.Redis
}

func New(
	config interfaces.Config,
	storage interfaces.Storage,
	redis interfaces.Redis,
) *App {
	return &App{
		Config:  config,
		Storage: storage,
		Redis:   redis,
	}
}
