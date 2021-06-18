package app

import "github.com/and67o/go-comments/internal/interfaces"

type App struct {
	Config  interfaces.Config
}

func New(config interfaces.Config) *App {
	return &App{
		Config: config,
	}
}
