package internalhttp

import (
	"context"
	"fmt"
	"github.com/and67o/go-comments/internal/app"
	"github.com/and67o/go-comments/internal/interfaces"
	"github.com/and67o/go-comments/internal/router"
	"net"
	"net/http"
	"time"
)

type Server struct {
	app    *app.App
	server *http.Server
}

func (s Server) Start() error {
	err := s.server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("server start: %w", err)
	}
	return err
}

func (s Server) Stop(ctx context.Context) error {
	err := s.server.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("shutdown error: %w", err)
	}

	return nil
}

func New(app *app.App) interfaces.HTTPApp {
	r := router.New(app)

	httpConfig := app.Config.GetHTTP()
	addr := net.JoinHostPort(httpConfig.Host, httpConfig.Port)

	return &Server{
		app: app,
		server: &http.Server{
			Addr:    addr,
			Handler: r.GetRouter(),
			ReadTimeout: httpConfig.Timeout.Read * time.Second,
			WriteTimeout: httpConfig.Timeout.Write * time.Second,
			IdleTimeout: httpConfig.Timeout.Idle * time.Second,
		},
	}
}
